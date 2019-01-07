package codegen

import (
	"go/types"
	"strings"

	"github.com/vektah/gqlparser/ast"
	"golang.org/x/tools/go/loader"
)

// namedTypeFromSchema objects for every graphql type, including scalars. There should only be one instance of TypeReference for each thing
func (g *Generator) buildNamedTypes() NamedTypes {
	types := map[string]*TypeDefinition{}
	for _, schemaType := range g.schema.Types {
		t := namedTypeFromSchema(schemaType)

		if userEntry, ok := g.Models[t.GQLType]; ok && userEntry.Model != "" {
			t.IsUserDefined = true
			t.Package, t.GoType = pkgAndType(userEntry.Model)
		} else if t.IsScalar {
			t.Package = "github.com/99designs/gqlgen/graphql"
			t.GoType = "String"
		}

		types[t.GQLType] = t
	}
	return types
}

func (g *Generator) bindTypes(namedTypes NamedTypes, destDir string, prog *loader.Program) {
	for _, t := range namedTypes {
		if t.Package == "" {
			continue
		}

		def, _ := findGoType(prog, t.Package, "Marshal"+t.GoType)
		switch def := def.(type) {
		case *types.Func:
			sig := def.Type().(*types.Signature)
			cpy := t.TypeImplementation
			t.Marshaler = &cpy

			t.Package, t.GoType = pkgAndType(sig.Params().At(0).Type().String())
		}
	}
}

// namedTypeFromSchema objects for every graphql type, including primitives.
// don't recurse into object fields or interfaces yet, lets make sure we have collected everything first.
func namedTypeFromSchema(schemaType *ast.Definition) *TypeDefinition {
	switch schemaType.Kind {
	case ast.Scalar, ast.Enum:
		return &TypeDefinition{GQLType: schemaType.Name, IsScalar: true}
	case ast.Interface, ast.Union:
		return &TypeDefinition{GQLType: schemaType.Name, IsInterface: true}
	case ast.InputObject:
		return &TypeDefinition{GQLType: schemaType.Name, IsInput: true}
	default:
		return &TypeDefinition{GQLType: schemaType.Name}
	}
}

// take a string in the form github.com/package/blah.TypeReference and split it into package and type
func pkgAndType(name string) (string, string) {
	parts := strings.Split(name, ".")
	if len(parts) == 1 {
		return "", name
	}

	return normalizeVendor(strings.Join(parts[:len(parts)-1], ".")), parts[len(parts)-1]
}
