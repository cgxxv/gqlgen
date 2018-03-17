// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT

package test

import (
	"bytes"
	context "context"
	fmt "fmt"
	strconv "strconv"

	graphql "github.com/vektah/gqlgen/graphql"
	errors "github.com/vektah/gqlgen/neelance/errors"
	introspection "github.com/vektah/gqlgen/neelance/introspection"
	query "github.com/vektah/gqlgen/neelance/query"
	schema "github.com/vektah/gqlgen/neelance/schema"
)

func MakeExecutableSchema(resolvers Resolvers) graphql.ExecutableSchema {
	return &executableSchema{resolvers}
}

type Resolvers interface {
	OuterObject_inner(ctx context.Context, obj *OuterObject) (InnerObject, error)
	Query_nestedInputs(ctx context.Context, input [][]OuterInput) (*bool, error)
	Query_nestedOutputs(ctx context.Context) ([][]OuterObject, error)
	Query_shapes(ctx context.Context) ([]Shape, error)
}

type executableSchema struct {
	resolvers Resolvers
}

func (e *executableSchema) Schema() *schema.Schema {
	return parsedSchema
}

func (e *executableSchema) Query(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation, recover graphql.RecoverFunc) *graphql.Response {
	ec := executionContext{resolvers: e.resolvers, variables: variables, doc: doc, ctx: ctx, recover: recover}

	data := ec._Query(op.Selections)
	var buf bytes.Buffer
	data.MarshalGQL(&buf)

	return &graphql.Response{
		Data:   buf.Bytes(),
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation, recover graphql.RecoverFunc) *graphql.Response {
	return &graphql.Response{Errors: []*errors.QueryError{{Message: "mutations are not supported"}}}
}

func (e *executableSchema) Subscription(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation, recover graphql.RecoverFunc) func() *graphql.Response {
	return graphql.OneShot(&graphql.Response{Errors: []*errors.QueryError{{Message: "subscriptions are not supported"}}})
}

type executionContext struct {
	errors.Builder
	resolvers Resolvers
	variables map[string]interface{}
	doc       *query.Document
	ctx       context.Context
	recover   graphql.RecoverFunc
}

var circleImplementors = []string{"Circle", "Shape"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Circle(sel []query.Selection, obj *Circle) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, circleImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Circle")
		case "radius":
			out.Values[i] = ec._Circle_radius(field, obj)
		case "area":
			out.Values[i] = ec._Circle_area(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Circle_radius(field graphql.CollectedField, obj *Circle) graphql.Marshaler {
	res := obj.Radius
	return graphql.MarshalFloat(res)
}

func (ec *executionContext) _Circle_area(field graphql.CollectedField, obj *Circle) graphql.Marshaler {
	res := obj.Area()
	return graphql.MarshalFloat(res)
}

var innerObjectImplementors = []string{"InnerObject"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _InnerObject(sel []query.Selection, obj *InnerObject) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, innerObjectImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("InnerObject")
		case "id":
			out.Values[i] = ec._InnerObject_id(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _InnerObject_id(field graphql.CollectedField, obj *InnerObject) graphql.Marshaler {
	res := obj.ID
	return graphql.MarshalInt(res)
}

var outerObjectImplementors = []string{"OuterObject"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _OuterObject(sel []query.Selection, obj *OuterObject) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, outerObjectImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("OuterObject")
		case "inner":
			out.Values[i] = ec._OuterObject_inner(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _OuterObject_inner(field graphql.CollectedField, obj *OuterObject) graphql.Marshaler {
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		res, err := ec.resolvers.OuterObject_inner(ec.ctx, obj)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
		return ec._InnerObject(field.Selections, &res)
	})
}

var queryImplementors = []string{"Query"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Query(sel []query.Selection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, queryImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "nestedInputs":
			out.Values[i] = ec._Query_nestedInputs(field)
		case "nestedOutputs":
			out.Values[i] = ec._Query_nestedOutputs(field)
		case "shapes":
			out.Values[i] = ec._Query_shapes(field)
		case "__schema":
			out.Values[i] = ec._Query___schema(field)
		case "__type":
			out.Values[i] = ec._Query___type(field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Query_nestedInputs(field graphql.CollectedField) graphql.Marshaler {
	var arg0 [][]OuterInput
	if tmp, ok := field.Args["input"]; ok {
		var err error
		rawIf1 := tmp.([]interface{})
		arg0 = make([][]OuterInput, len(rawIf1))
		for idx1 := range rawIf1 {
			rawIf2 := rawIf1[idx1].([]interface{})
			arg0[idx1] = make([]OuterInput, len(rawIf2))
			for idx2 := range rawIf2 {

				arg0[idx1][idx2], err = UnmarshalOuterInput(rawIf2[idx2])
			}
		}
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	} else {
		var tmp interface{} = []interface{}{[]interface{}{map[string]interface{}{"inner": map[string]interface{}{"id": 1}}}}
		var err error
		rawIf1 := tmp.([]interface{})
		arg0 = make([][]OuterInput, len(rawIf1))
		for idx1 := range rawIf1 {
			rawIf2 := rawIf1[idx1].([]interface{})
			arg0[idx1] = make([]OuterInput, len(rawIf2))
			for idx2 := range rawIf2 {

				arg0[idx1][idx2], err = UnmarshalOuterInput(rawIf2[idx2])
			}
		}
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}

	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		res, err := ec.resolvers.Query_nestedInputs(ec.ctx, arg0)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
		if res == nil {
			return graphql.Null
		}
		return graphql.MarshalBoolean(*res)
	})
}

func (ec *executionContext) _Query_nestedOutputs(field graphql.CollectedField) graphql.Marshaler {
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		res, err := ec.resolvers.Query_nestedOutputs(ec.ctx)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
		arr1 := graphql.Array{}
		for idx1 := range res {
			arr1 = append(arr1, func() graphql.Marshaler {
				arr2 := graphql.Array{}
				for idx2 := range res[idx1] {
					arr2 = append(arr2, func() graphql.Marshaler { return ec._OuterObject(field.Selections, &res[idx1][idx2]) }())
				}
				return arr2
			}())
		}
		return arr1
	})
}

func (ec *executionContext) _Query_shapes(field graphql.CollectedField) graphql.Marshaler {
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		res, err := ec.resolvers.Query_shapes(ec.ctx)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
		arr1 := graphql.Array{}
		for idx1 := range res {
			arr1 = append(arr1, func() graphql.Marshaler { return ec._Shape(field.Selections, &res[idx1]) }())
		}
		return arr1
	})
}

func (ec *executionContext) _Query___schema(field graphql.CollectedField) graphql.Marshaler {
	res := ec.introspectSchema()
	if res == nil {
		return graphql.Null
	}
	return ec.___Schema(field.Selections, res)
}

func (ec *executionContext) _Query___type(field graphql.CollectedField) graphql.Marshaler {
	var arg0 string
	if tmp, ok := field.Args["name"]; ok {
		var err error

		arg0, err = graphql.UnmarshalString(tmp)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	res := ec.introspectType(arg0)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

var rectangleImplementors = []string{"Rectangle", "Shape"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Rectangle(sel []query.Selection, obj *Rectangle) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, rectangleImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Rectangle")
		case "length":
			out.Values[i] = ec._Rectangle_length(field, obj)
		case "width":
			out.Values[i] = ec._Rectangle_width(field, obj)
		case "area":
			out.Values[i] = ec._Rectangle_area(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Rectangle_length(field graphql.CollectedField, obj *Rectangle) graphql.Marshaler {
	res := obj.Length
	return graphql.MarshalFloat(res)
}

func (ec *executionContext) _Rectangle_width(field graphql.CollectedField, obj *Rectangle) graphql.Marshaler {
	res := obj.Width
	return graphql.MarshalFloat(res)
}

func (ec *executionContext) _Rectangle_area(field graphql.CollectedField, obj *Rectangle) graphql.Marshaler {
	res := obj.Area()
	return graphql.MarshalFloat(res)
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(sel []query.Selection, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __DirectiveImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(field, obj)
		case "description":
			out.Values[i] = ec.___Directive_description(field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(field, obj)
		case "args":
			out.Values[i] = ec.___Directive_args(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Directive_name(field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Directive_description(field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Directive_locations(field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	res := obj.Locations()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler { return graphql.MarshalString(res[idx1]) }())
	}
	return arr1
}

func (ec *executionContext) ___Directive_args(field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	res := obj.Args()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___InputValue(field.Selections, res[idx1])
		}())
	}
	return arr1
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(sel []query.Selection, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __EnumValueImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(field, obj)
		case "description":
			out.Values[i] = ec.___EnumValue_description(field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(field, obj)
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___EnumValue_name(field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___EnumValue_description(field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___EnumValue_isDeprecated(field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	res := obj.IsDeprecated()
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) ___EnumValue_deprecationReason(field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	res := obj.DeprecationReason()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(sel []query.Selection, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __FieldImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(field, obj)
		case "description":
			out.Values[i] = ec.___Field_description(field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(field, obj)
		case "type":
			out.Values[i] = ec.___Field_type(field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(field, obj)
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Field_name(field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Field_description(field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Field_args(field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.Args()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___InputValue(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Field_type(field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.Type()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

func (ec *executionContext) ___Field_isDeprecated(field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.IsDeprecated()
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) ___Field_deprecationReason(field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.DeprecationReason()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(sel []query.Selection, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __InputValueImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(field, obj)
		case "description":
			out.Values[i] = ec.___InputValue_description(field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(field, obj)
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___InputValue_name(field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___InputValue_description(field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___InputValue_type(field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	res := obj.Type()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

func (ec *executionContext) ___InputValue_defaultValue(field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	res := obj.DefaultValue()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(sel []query.Selection, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __SchemaImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(field, obj)
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(field, obj)
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Schema_types(field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.Types()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Type(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Schema_queryType(field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.QueryType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

func (ec *executionContext) ___Schema_mutationType(field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.MutationType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

func (ec *executionContext) ___Schema_subscriptionType(field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.SubscriptionType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

func (ec *executionContext) ___Schema_directives(field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.Directives()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Directive(field.Selections, res[idx1])
		}())
	}
	return arr1
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(sel []query.Selection, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __TypeImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(field, obj)
		case "name":
			out.Values[i] = ec.___Type_name(field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Type_kind(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.Kind()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Type_name(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.Name()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Type_description(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Type_fields(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	var arg0 bool
	if tmp, ok := field.Args["includeDeprecated"]; ok {
		var err error

		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	res := obj.Fields(arg0)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Field(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_interfaces(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.Interfaces()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Type(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_possibleTypes(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.PossibleTypes()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Type(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_enumValues(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	var arg0 bool
	if tmp, ok := field.Args["includeDeprecated"]; ok {
		var err error

		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	res := obj.EnumValues(arg0)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___EnumValue(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_inputFields(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.InputFields()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___InputValue(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_ofType(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.OfType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

func (ec *executionContext) _Shape(sel []query.Selection, obj *Shape) graphql.Marshaler {
	switch obj := (*obj).(type) {
	case nil:
		return graphql.Null
	case Circle:
		return ec._Circle(sel, &obj)

	case *Circle:
		return ec._Circle(sel, obj)
	case Rectangle:
		return ec._Rectangle(sel, &obj)

	case *Rectangle:
		return ec._Rectangle(sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func UnmarshalInnerInput(v interface{}) (InnerInput, error) {
	var it InnerInput

	for k, v := range v.(map[string]interface{}) {
		switch k {
		case "id":
			var err error

			it.ID, err = graphql.UnmarshalInt(v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func UnmarshalOuterInput(v interface{}) (OuterInput, error) {
	var it OuterInput

	for k, v := range v.(map[string]interface{}) {
		switch k {
		case "inner":
			var err error

			it.Inner, err = UnmarshalInnerInput(v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

var parsedSchema = schema.MustParse("input InnerInput {\n    id:Int!\n}\n\ninput OuterInput {\n    inner: InnerInput!\n}\n\ntype OuterObject {\n    inner: InnerObject!\n}\n\ntype InnerObject {\n    id: Int!\n}\n\ninterface Shape {\n    area: Float\n}\n\ntype Circle implements Shape {\n    radius: Float\n    area: Float\n}\n\ntype Rectangle implements Shape {\n    length: Float\n    width: Float\n    area: Float\n}\n\ntype Query {\n    nestedInputs(input: [[OuterInput]] = [[{inner: {id: 1}}]]): Boolean\n    nestedOutputs: [[OuterObject]]\n    shapes: [Shape]\n}\n")

func (ec *executionContext) introspectSchema() *introspection.Schema {
	return introspection.WrapSchema(parsedSchema)
}

func (ec *executionContext) introspectType(name string) *introspection.Type {
	t := parsedSchema.Resolve(name)
	if t == nil {
		return nil
	}
	return introspection.WrapType(t)
}
