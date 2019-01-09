package unified

import (
	"go/types"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/ast"
)

func (g *Schema) buildObject(typ *ast.Definition) (*Object, error) {
	dirs, err := g.getDirectives(typ.Directives)
	if err != nil {
		return nil, errors.Wrap(err, typ.Name)
	}

	isRoot := typ == g.Schema.Query || typ == g.Schema.Mutation || typ == g.Schema.Subscription

	obj := &Object{
		Definition:         g.NamedTypes[typ.Name],
		InTypemap:          g.Config.Models.UserDefined(typ.Name) || isRoot,
		Root:               isRoot,
		DisableConcurrency: typ == g.Schema.Mutation,
		Stream:             typ == g.Schema.Subscription,
		Directives:         dirs,
		ResolverInterface: types.NewNamed(
			types.NewTypeName(0, g.Config.Exec.Pkg(), typ.Name+"Resolver", nil),
			nil,
			nil,
		),
	}

	for _, intf := range g.Schema.GetImplements(typ) {
		obj.Implements = append(obj.Implements, g.NamedTypes[intf.Name])
	}

	for _, field := range typ.Fields {
		if strings.HasPrefix(field.Name, "__") {
			continue
		}

		f, err := g.buildField(obj, field)
		if err != nil {
			return nil, errors.Wrap(err, typ.Name+"."+field.Name)
		}

		if typ.Kind == ast.InputObject && !f.TypeReference.Definition.GQLDefinition.IsInputType() {
			return nil, errors.Errorf(
				"%s.%s: cannot use %s because %s is not a valid input type",
				typ.Name,
				field.Name,
				f.Definition.GQLDefinition.Name,
				f.TypeReference.Definition.GQLDefinition.Kind,
			)
		}

		obj.Fields = append(obj.Fields, f)
	}

	if _, isMap := obj.Definition.GoType.(*types.Map); !isMap && obj.InTypemap {
		for _, bindErr := range bindObject(obj, g.Config.StructTag) {
			log.Println(bindErr.Error())
			log.Println("  Adding resolver method")
		}
	}

	return obj, nil
}

func (g *Schema) buildField(obj *Object, field *ast.FieldDefinition) (*Field, error) {
	dirs, err := g.getDirectives(field.Directives)
	if err != nil {
		return nil, err
	}

	f := Field{
		GQLName:        field.Name,
		TypeReference:  g.NamedTypes.getType(field.Type),
		Object:         obj,
		Directives:     dirs,
		GoFieldName:    lintName(ucFirst(field.Name)),
		GoFieldType:    GoFieldVariable,
		GoReceiverName: "obj",
	}

	if field.DefaultValue != nil {
		var err error
		f.Default, err = field.DefaultValue.Value(nil)
		if err != nil {
			return nil, errors.Errorf("default value %s is not valid: %s", field.Name, err.Error())
		}
	}

	typeEntry, entryExists := g.Config.Models[obj.Definition.GQLDefinition.Name]
	if entryExists {
		if typeField, ok := typeEntry.Fields[field.Name]; ok {
			if typeField.Resolver {
				f.IsResolver = true
			}
			if typeField.FieldName != "" {
				f.GoFieldName = lintName(ucFirst(typeField.FieldName))
			}
		}
	}

	for _, arg := range field.Arguments {
		newArg, err := g.buildArg(obj, arg)
		if err != nil {
			return nil, err
		}
		f.Args = append(f.Args, newArg)
	}
	return &f, nil
}

func (g *Schema) buildArg(obj *Object, arg *ast.ArgumentDefinition) (*FieldArgument, error) {
	argDirs, err := g.getDirectives(arg.Directives)
	if err != nil {
		return nil, err
	}
	newArg := FieldArgument{
		GQLName:       arg.Name,
		TypeReference: g.NamedTypes.getType(arg.Type),
		Object:        obj,
		GoVarName:     sanitizeArgName(arg.Name),
		Directives:    argDirs,
	}

	if !newArg.TypeReference.Definition.GQLDefinition.IsInputType() {
		return nil, errors.Errorf(
			"cannot use %s as argument %s because %s is not a valid input type",
			newArg.Definition.GQLDefinition.Name,
			arg.Name,
			newArg.TypeReference.Definition.GQLDefinition.Kind,
		)
	}

	if arg.DefaultValue != nil {
		var err error
		newArg.Default, err = arg.DefaultValue.Value(nil)
		if err != nil {
			return nil, errors.Errorf("default value is not valid: %s", err.Error())
		}
	}

	return &newArg, nil
}
