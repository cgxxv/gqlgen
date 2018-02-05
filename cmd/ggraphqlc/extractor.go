package main

import (
	"fmt"
	"go/types"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/vektah/graphql-go/common"
	"github.com/vektah/graphql-go/schema"
	"golang.org/x/tools/go/loader"
)

type extractor struct {
	Errors      []string
	PackageName string
	Objects     []object
	goTypeMap   map[string]string
	Imports     map[string]string // local -> full path
	schemaRaw   string
}

func (e *extractor) errorf(format string, args ...interface{}) {
	e.Errors = append(e.Errors, fmt.Sprintf(format, args...))
}

// getType to put in a file for a given fully resolved type, and add any Imports required
// eg name = github.com/my/pkg.myType will return `pkg.myType` and add an import for `github.com/my/pkg`
func (e *extractor) getType(name string) Type {
	if fieldType, ok := e.goTypeMap[name]; ok {
		parts := strings.Split(fieldType, ".")
		if len(parts) == 1 {
			return Type{
				GraphQLName: name,
				Name:        parts[0],
			}
		}

		packageName := strings.Join(parts[:len(parts)-1], ".")
		typeName := parts[len(parts)-1]

		localName := filepath.Base(packageName)
		i := 0
		for pkg, found := e.Imports[localName]; found && pkg != packageName; localName = filepath.Base(packageName) + strconv.Itoa(i) {
			i++
			if i > 10 {
				panic("too many collisions")
			}
		}

		e.Imports[localName] = packageName
		return Type{
			GraphQLName: name,
			ImportedAs:  localName,
			Name:        typeName,
			Package:     packageName,
		}
	}
	fmt.Fprintf(os.Stderr, "unknown go type for %s, using interface{}. you should add it to types.json\n", name)
	e.goTypeMap[name] = "interface{}"
	return Type{
		GraphQLName: name,
		Name:        "interface{}",
	}
}

func (e *extractor) buildType(t common.Type) Type {
	var modifiers []string
	usePtr := true
	for {
		if _, nonNull := t.(*common.NonNull); nonNull {
			usePtr = false
		} else if _, nonNull := t.(*common.List); nonNull {
			usePtr = false
		} else {
			if usePtr {
				modifiers = append(modifiers, modPtr)
			}
			usePtr = true
		}

		switch val := t.(type) {
		case *common.NonNull:
			t = val.OfType
		case *common.List:
			modifiers = append(modifiers, modList)
			t = val.OfType
		case *schema.Scalar:
			var goType string

			switch val.Name {
			case "String":
				goType = "string"
			case "ID":
				goType = "string"
			case "Boolean":
				goType = "bool"
			case "Int":
				goType = "int"
			case "Float":
				goType = "float64"
			default:
				panic(fmt.Errorf("unknown scalar %s", val.Name))
			}
			return Type{
				Basic:       true,
				Modifiers:   modifiers,
				GraphQLName: val.Name,
				Name:        goType,
			}
		case *schema.Object:
			t := e.getType(val.Name)
			t.Modifiers = modifiers
			return t
		case *common.TypeName:
			t := e.getType(val.Name)
			t.Modifiers = modifiers
			return t
		case *schema.Interface:
			t := e.getType(val.Name)
			t.Modifiers = modifiers
			if t.Modifiers[len(t.Modifiers)-1] == modPtr {
				t.Modifiers = t.Modifiers[0 : len(t.Modifiers)-1]
			}

			for _, implementor := range val.PossibleTypes {
				t.Implementors = append(t.Implementors, e.getType(implementor.Name))
			}

			return t
		case *schema.Union:
			t := e.getType(val.Name)
			t.Modifiers = modifiers

			for _, implementor := range val.PossibleTypes {
				t.Implementors = append(t.Implementors, e.getType(implementor.Name))
			}

			return t
		case *schema.InputObject:
			t := e.getType(val.Name)
			t.Modifiers = modifiers
			return t
		case *schema.Enum:
			return Type{
				Basic:       true,
				Modifiers:   modifiers,
				GraphQLName: val.Name,
				Name:        "string",
			}
		default:
			panic(fmt.Errorf("unknown type %T", t))
		}
	}
}

func (e *extractor) extract(s *schema.Schema) {
	for _, schemaType := range s.Types {
		schemaObject, ok := schemaType.(*schema.Object)
		if !ok {
			continue
		}
		object := object{
			Name: schemaObject.Name,
			Type: e.getType(schemaObject.Name),
		}

		for _, i := range schemaObject.Interfaces {
			object.satisfies = append(object.satisfies, i.Name)
		}

		for _, field := range schemaObject.Fields {
			var args []Arg
			for _, arg := range field.Args {
				args = append(args, Arg{
					Name: arg.Name.Name,
					Type: e.buildType(arg.Type),
				})
			}

			object.Fields = append(object.Fields, Field{
				GraphQLName: field.Name,
				Type:        e.buildType(field.Type),
				Args:        args,
			})
		}
		e.Objects = append(e.Objects, object)
	}

	sort.Slice(e.Objects, func(i, j int) bool {
		return strings.Compare(e.Objects[i].Name, e.Objects[j].Name) == -1
	})
}

func (e *extractor) introspect() error {
	var conf loader.Config
	for _, name := range e.Imports {
		conf.Import(name)
	}

	prog, err := conf.Load()
	if err != nil {
		return err
	}

	for _, o := range e.Objects {
		if o.Type.Package == "" {
			continue
		}
		pkg := prog.Package(o.Type.Package)

		for astNode, object := range pkg.Defs {
			if astNode.Name != o.Type.Name {
				continue
			}

			e.findBindTargets(object.Type(), o)
			// todo: break!
		}
	}

	return nil
}

func (e *extractor) modifiersFromGoType(t types.Type) []string {
	var modifiers []string
	for {
		switch val := t.(type) {
		case *types.Pointer:
			modifiers = append(modifiers, modPtr)
			t = val.Elem()
		case *types.Array:
			modifiers = append(modifiers, modList)
			t = val.Elem()
		case *types.Slice:
			modifiers = append(modifiers, modList)
			t = val.Elem()
		default:
			return modifiers
		}
	}
}

func (e *extractor) findBindTargets(t types.Type, object object) {
	switch t := t.(type) {
	case *types.Named:
		for i := 0; i < t.NumMethods(); i++ {
			method := t.Method(i)
			if methodField := object.GetField(method.Name()); methodField != nil {
				methodField.MethodName = "it." + method.Name()
				sig := method.Type().(*types.Signature)

				methodField.Type.Modifiers = e.modifiersFromGoType(sig.Results().At(0).Type())

				// check arg order matches code, not gql

				var newArgs []Arg
			l2:
				for j := 0; j < sig.Params().Len(); j++ {
					param := sig.Params().At(j)
					for _, oldArg := range methodField.Args {
						if strings.EqualFold(oldArg.Name, param.Name()) {
							oldArg.Type.Modifiers = e.modifiersFromGoType(param.Type())
							newArgs = append(newArgs, oldArg)
							continue l2
						}
					}
					e.errorf("cannot match argument " + param.Name() + " to any argument in " + t.String())
				}
				methodField.Args = newArgs

				if sig.Results().Len() == 1 {
					methodField.NoErr = true
				} else if sig.Results().Len() != 2 {
					e.errorf("weird number of results on %s. expected either (result), or (result, error)", method.Name())
				}
			}
		}

		e.findBindTargets(t.Underlying(), object)

	case *types.Struct:
		for i := 0; i < t.NumFields(); i++ {
			field := t.Field(i)
			// Todo: struct tags, name and - at least

			// Todo: check for type matches before binding too?
			if objectField := object.GetField(field.Name()); objectField != nil {
				objectField.VarName = "it." + field.Name()
				objectField.Type.Modifiers = e.modifiersFromGoType(field.Type())
			}
		}
		t.Underlying()

	case *types.Signature:
		// ignored

	default:
		panic(fmt.Errorf("unknown type %T looking at %s", t, object.Name))
	}

}

const (
	modList = "[]"
	modPtr  = "*"
)

type Type struct {
	GraphQLName  string
	Name         string
	Package      string
	ImportedAs   string
	Modifiers    []string
	Basic        bool
	Implementors []Type
}

func (t Type) Local() string {
	if t.ImportedAs == "" {
		return strings.Join(t.Modifiers, "") + t.Name
	}
	return strings.Join(t.Modifiers, "") + t.ImportedAs + "." + t.Name
}

func (t Type) Ptr() Type {
	t.Modifiers = append(t.Modifiers, modPtr)
	return t
}

type object struct {
	Name      string
	Fields    []Field
	Type      Type
	satisfies []string
}

type Field struct {
	GraphQLName string
	MethodName  string
	VarName     string
	Type        Type
	Args        []Arg
	NoErr       bool
}

func (o *object) GetField(name string) *Field {
	for i, field := range o.Fields {
		if strings.EqualFold(field.GraphQLName, name) {
			return &o.Fields[i]
		}
	}
	return nil
}

func (e *extractor) GetObject(name string) *object {
	for i, o := range e.Objects {
		if strings.EqualFold(o.Name, name) {
			return &e.Objects[i]
		}
	}
	return nil
}

type Arg struct {
	Name string
	Type Type
}
