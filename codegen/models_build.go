package codegen

import (
	"sort"

	"github.com/vektah/gqlparser/ast"
	"golang.org/x/tools/go/loader"
)

func (g *Generator) buildModels(types NamedTypes, prog *loader.Program) ([]Model, error) {
	var models []Model

	for _, typ := range g.schema.Types {
		var model Model
		switch typ.Kind {
		case ast.Object:
			obj, err := g.buildObject(types, typ)
			if err != nil {
				return nil, err
			}
			if obj.Root || obj.IsUserDefined {
				continue
			}
			model = g.obj2Model(obj)
		case ast.InputObject:
			obj, err := g.buildInput(types, typ)
			if err != nil {
				return nil, err
			}
			if obj.IsUserDefined {
				continue
			}
			model = g.obj2Model(obj)
		case ast.Interface, ast.Union:
			intf := g.buildInterface(types, typ, prog)
			if intf.IsUserDefined {
				continue
			}
			model = int2Model(intf)
		default:
			continue
		}
		model.Description = typ.Description // It's this or change both obj2Model and buildObject

		models = append(models, model)
	}

	sort.Slice(models, func(i, j int) bool {
		return models[i].GQLType < models[j].GQLType
	})

	return models, nil
}

func (g *Generator) obj2Model(obj *Object) Model {
	model := Model{
		TypeDefinition: obj.TypeDefinition,
		Implements:     obj.Implements,
		Fields:         []ModelField{},
	}

	model.GoType = ucFirst(obj.GQLType)
	model.Marshaler = &TypeImplementation{GoType: obj.GoType}

	for i := range obj.Fields {
		field := &obj.Fields[i]
		mf := ModelField{TypeReference: field.TypeReference, GQLName: field.GQLName}

		if field.GoFieldName != "" {
			mf.GoFieldName = field.GoFieldName
		} else {
			mf.GoFieldName = field.GoNameExported()
		}

		model.Fields = append(model.Fields, mf)
	}

	return model
}

func int2Model(obj *Interface) Model {
	model := Model{
		TypeDefinition: obj.TypeDefinition,
		Fields:         []ModelField{},
	}

	model.GoType = ucFirst(obj.GQLType)
	model.Marshaler = &TypeImplementation{GoType: obj.GoType}

	return model
}
