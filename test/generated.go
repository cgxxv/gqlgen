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
	introspection1 "github.com/vektah/gqlgen/test/introspection"
	invalid_identifier "github.com/vektah/gqlgen/test/invalid-identifier"
	models "github.com/vektah/gqlgen/test/models"
)

func MakeExecutableSchema(resolvers Resolvers) graphql.ExecutableSchema {
	return &executableSchema{resolvers: resolvers}
}

type Resolvers interface {
	OuterObject_inner(ctx context.Context, obj *models.OuterObject) (models.InnerObject, error)
	Query_nestedInputs(ctx context.Context, input [][]models.OuterInput) (*bool, error)
	Query_nestedOutputs(ctx context.Context) ([][]models.OuterObject, error)
	Query_shapes(ctx context.Context) ([]Shape, error)
	Query_recursive(ctx context.Context, input *RecursiveInputSlice) (*bool, error)
	Query_mapInput(ctx context.Context, input *map[string]interface{}) (*bool, error)
	Query_collision(ctx context.Context) (*introspection1.It, error)
	Query_invalidIdentifier(ctx context.Context) (*invalid_identifier.InvalidIdentifier, error)
}

type executableSchema struct {
	resolvers Resolvers
}

func (e *executableSchema) Schema() *schema.Schema {
	return parsedSchema
}

func (e *executableSchema) Query(ctx context.Context, op *query.Operation) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e.resolvers}

	data := ec._Query(ctx, op.Selections)
	var buf bytes.Buffer
	data.MarshalGQL(&buf)

	return &graphql.Response{
		Data:   buf.Bytes(),
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, op *query.Operation) *graphql.Response {
	return &graphql.Response{Errors: []*errors.QueryError{{Message: "mutations are not supported"}}}
}

func (e *executableSchema) Subscription(ctx context.Context, op *query.Operation) func() *graphql.Response {
	return graphql.OneShot(&graphql.Response{Errors: []*errors.QueryError{{Message: "subscriptions are not supported"}}})
}

type executionContext struct {
	*graphql.RequestContext

	resolvers Resolvers
}

var circleImplementors = []string{"Circle", "Shape"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Circle(ctx context.Context, sel []query.Selection, obj *Circle) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, circleImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Circle")
		case "radius":
			out.Values[i] = ec._Circle_radius(ctx, field, obj)
		case "area":
			out.Values[i] = ec._Circle_area(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Circle_radius(ctx context.Context, field graphql.CollectedField, obj *Circle) graphql.Marshaler {
	res := obj.Radius
	return graphql.MarshalFloat(res)
}

func (ec *executionContext) _Circle_area(ctx context.Context, field graphql.CollectedField, obj *Circle) graphql.Marshaler {
	res := obj.Area()
	return graphql.MarshalFloat(res)
}

var innerObjectImplementors = []string{"InnerObject"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _InnerObject(ctx context.Context, sel []query.Selection, obj *models.InnerObject) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, innerObjectImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("InnerObject")
		case "id":
			out.Values[i] = ec._InnerObject_id(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _InnerObject_id(ctx context.Context, field graphql.CollectedField, obj *models.InnerObject) graphql.Marshaler {
	res := obj.ID
	return graphql.MarshalInt(res)
}

var invalidIdentifierImplementors = []string{"InvalidIdentifier"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _InvalidIdentifier(ctx context.Context, sel []query.Selection, obj *invalid_identifier.InvalidIdentifier) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, invalidIdentifierImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("InvalidIdentifier")
		case "id":
			out.Values[i] = ec._InvalidIdentifier_id(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _InvalidIdentifier_id(ctx context.Context, field graphql.CollectedField, obj *invalid_identifier.InvalidIdentifier) graphql.Marshaler {
	res := obj.ID
	return graphql.MarshalInt(res)
}

var itImplementors = []string{"It"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _It(ctx context.Context, sel []query.Selection, obj *introspection1.It) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, itImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("It")
		case "id":
			out.Values[i] = ec._It_id(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _It_id(ctx context.Context, field graphql.CollectedField, obj *introspection1.It) graphql.Marshaler {
	res := obj.ID
	return graphql.MarshalID(res)
}

var outerObjectImplementors = []string{"OuterObject"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _OuterObject(ctx context.Context, sel []query.Selection, obj *models.OuterObject) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, outerObjectImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("OuterObject")
		case "inner":
			out.Values[i] = ec._OuterObject_inner(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _OuterObject_inner(ctx context.Context, field graphql.CollectedField, obj *models.OuterObject) graphql.Marshaler {
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		rctx := graphql.WithResolverContext(ctx, &graphql.ResolverContext{Field: field})
		res, err := ec.resolvers.OuterObject_inner(rctx, obj)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
		return ec._InnerObject(ctx, field.Selections, &res)
	})
}

var queryImplementors = []string{"Query"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Query(ctx context.Context, sel []query.Selection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, queryImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "nestedInputs":
			out.Values[i] = ec._Query_nestedInputs(ctx, field)
		case "nestedOutputs":
			out.Values[i] = ec._Query_nestedOutputs(ctx, field)
		case "shapes":
			out.Values[i] = ec._Query_shapes(ctx, field)
		case "recursive":
			out.Values[i] = ec._Query_recursive(ctx, field)
		case "mapInput":
			out.Values[i] = ec._Query_mapInput(ctx, field)
		case "collision":
			out.Values[i] = ec._Query_collision(ctx, field)
		case "invalidIdentifier":
			out.Values[i] = ec._Query_invalidIdentifier(ctx, field)
		case "__schema":
			out.Values[i] = ec._Query___schema(ctx, field)
		case "__type":
			out.Values[i] = ec._Query___type(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Query_nestedInputs(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	var arg0 [][]models.OuterInput
	if tmp, ok := field.Args["input"]; ok {
		var err error
		rawIf1 := tmp.([]interface{})
		arg0 = make([][]models.OuterInput, len(rawIf1))
		for idx1 := range rawIf1 {
			rawIf2 := rawIf1[idx1].([]interface{})
			arg0[idx1] = make([]models.OuterInput, len(rawIf2))
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
		arg0 = make([][]models.OuterInput, len(rawIf1))
		for idx1 := range rawIf1 {
			rawIf2 := rawIf1[idx1].([]interface{})
			arg0[idx1] = make([]models.OuterInput, len(rawIf2))
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
				userErr := ec.Recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		rctx := graphql.WithResolverContext(ctx, &graphql.ResolverContext{Field: field})
		res, err := ec.resolvers.Query_nestedInputs(rctx, arg0)
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

func (ec *executionContext) _Query_nestedOutputs(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		rctx := graphql.WithResolverContext(ctx, &graphql.ResolverContext{Field: field})
		res, err := ec.resolvers.Query_nestedOutputs(rctx)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
		arr1 := graphql.Array{}
		for idx1 := range res {
			arr1 = append(arr1, func() graphql.Marshaler {
				arr2 := graphql.Array{}
				for idx2 := range res[idx1] {
					arr2 = append(arr2, func() graphql.Marshaler { return ec._OuterObject(ctx, field.Selections, &res[idx1][idx2]) }())
				}
				return arr2
			}())
		}
		return arr1
	})
}

func (ec *executionContext) _Query_shapes(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		rctx := graphql.WithResolverContext(ctx, &graphql.ResolverContext{Field: field})
		res, err := ec.resolvers.Query_shapes(rctx)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
		arr1 := graphql.Array{}
		for idx1 := range res {
			arr1 = append(arr1, func() graphql.Marshaler { return ec._Shape(ctx, field.Selections, &res[idx1]) }())
		}
		return arr1
	})
}

func (ec *executionContext) _Query_recursive(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	var arg0 *RecursiveInputSlice
	if tmp, ok := field.Args["input"]; ok {
		var err error
		var ptr1 RecursiveInputSlice
		if tmp != nil {
			ptr1, err = UnmarshalRecursiveInputSlice(tmp)
			arg0 = &ptr1
		}

		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		rctx := graphql.WithResolverContext(ctx, &graphql.ResolverContext{Field: field})
		res, err := ec.resolvers.Query_recursive(rctx, arg0)
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

func (ec *executionContext) _Query_mapInput(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	var arg0 *map[string]interface{}
	if tmp, ok := field.Args["input"]; ok {
		var err error
		var ptr1 map[string]interface{}
		if tmp != nil {
			ptr1 = tmp.(map[string]interface{})
			arg0 = &ptr1
		}

		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		rctx := graphql.WithResolverContext(ctx, &graphql.ResolverContext{Field: field})
		res, err := ec.resolvers.Query_mapInput(rctx, arg0)
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

func (ec *executionContext) _Query_collision(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		rctx := graphql.WithResolverContext(ctx, &graphql.ResolverContext{Field: field})
		res, err := ec.resolvers.Query_collision(rctx)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
		if res == nil {
			return graphql.Null
		}
		return ec._It(ctx, field.Selections, res)
	})
}

func (ec *executionContext) _Query_invalidIdentifier(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		rctx := graphql.WithResolverContext(ctx, &graphql.ResolverContext{Field: field})
		res, err := ec.resolvers.Query_invalidIdentifier(rctx)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
		if res == nil {
			return graphql.Null
		}
		return ec._InvalidIdentifier(ctx, field.Selections, res)
	})
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	res := ec.introspectSchema()
	if res == nil {
		return graphql.Null
	}
	return ec.___Schema(ctx, field.Selections, res)
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
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
	return ec.___Type(ctx, field.Selections, res)
}

var rectangleImplementors = []string{"Rectangle", "Shape"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Rectangle(ctx context.Context, sel []query.Selection, obj *Rectangle) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, rectangleImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Rectangle")
		case "length":
			out.Values[i] = ec._Rectangle_length(ctx, field, obj)
		case "width":
			out.Values[i] = ec._Rectangle_width(ctx, field, obj)
		case "area":
			out.Values[i] = ec._Rectangle_area(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Rectangle_length(ctx context.Context, field graphql.CollectedField, obj *Rectangle) graphql.Marshaler {
	res := obj.Length
	return graphql.MarshalFloat(res)
}

func (ec *executionContext) _Rectangle_width(ctx context.Context, field graphql.CollectedField, obj *Rectangle) graphql.Marshaler {
	res := obj.Width
	return graphql.MarshalFloat(res)
}

func (ec *executionContext) _Rectangle_area(ctx context.Context, field graphql.CollectedField, obj *Rectangle) graphql.Marshaler {
	res := obj.Area()
	return graphql.MarshalFloat(res)
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(ctx context.Context, sel []query.Selection, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, __DirectiveImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	res := obj.Locations()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler { return graphql.MarshalString(res[idx1]) }())
	}
	return arr1
}

func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	res := obj.Args()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___InputValue(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(ctx context.Context, sel []query.Selection, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, __EnumValueImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	res := obj.IsDeprecated()
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	res := obj.DeprecationReason()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(ctx context.Context, sel []query.Selection, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, __FieldImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.Args()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___InputValue(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.Type()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.IsDeprecated()
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.DeprecationReason()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(ctx context.Context, sel []query.Selection, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, __InputValueImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	res := obj.Type()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	res := obj.DefaultValue()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(ctx context.Context, sel []query.Selection, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, __SchemaImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.Types()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Type(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.QueryType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.MutationType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.SubscriptionType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.Directives()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Directive(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(ctx context.Context, sel []query.Selection, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, __TypeImplementors, ec.Variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.Kind()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.Name()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
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
			return ec.___Field(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.Interfaces()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Type(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.PossibleTypes()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Type(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
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
			return ec.___EnumValue(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.InputFields()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___InputValue(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.OfType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) _Shape(ctx context.Context, sel []query.Selection, obj *Shape) graphql.Marshaler {
	switch obj := (*obj).(type) {
	case nil:
		return graphql.Null
	case *Circle:
		return ec._Circle(ctx, sel, obj)
	case *Rectangle:
		return ec._Rectangle(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func (ec *executionContext) _ShapeUnion(ctx context.Context, sel []query.Selection, obj *ShapeUnion) graphql.Marshaler {
	switch obj := (*obj).(type) {
	case nil:
		return graphql.Null
	case *Circle:
		return ec._Circle(ctx, sel, obj)
	case *Rectangle:
		return ec._Rectangle(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func UnmarshalInnerInput(v interface{}) (models.InnerInput, error) {
	var it models.InnerInput

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

func UnmarshalOuterInput(v interface{}) (models.OuterInput, error) {
	var it models.OuterInput

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

func UnmarshalRecursiveInputSlice(v interface{}) (RecursiveInputSlice, error) {
	var it RecursiveInputSlice

	for k, v := range v.(map[string]interface{}) {
		switch k {
		case "self":
			var err error
			var ptr1 []*RecursiveInputSlice
			if v != nil {
				rawIf2 := v.([]interface{})
				ptr1 = make([]*RecursiveInputSlice, len(rawIf2))
				for idx2 := range rawIf2 {
					var ptr3 RecursiveInputSlice
					if rawIf2[idx2] != nil {
						ptr3, err = UnmarshalRecursiveInputSlice(rawIf2[idx2])
						ptr1[idx2] = &ptr3
					}
				}
				it.Self = &ptr1
			}

			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

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

var parsedSchema = schema.MustParse(`input InnerInput {
    id:Int!
}

input OuterInput {
    inner: InnerInput!
}

type OuterObject {
    inner: InnerObject!
}

type InnerObject {
    id: Int!
}

interface Shape {
    area: Float
}

type Circle implements Shape {
    radius: Float
    area: Float
}

type Rectangle implements Shape {
    length: Float
    width: Float
    area: Float
}

input RecursiveInputSlice {
    self: [RecursiveInputSlice!]
}

union ShapeUnion = Circle | Rectangle

input Changes {
    a: Int
    b: Int
}

type It {
    id: ID!
}

type InvalidIdentifier {
    id: Int!
}

type Query {
    nestedInputs(input: [[OuterInput]] = [[{inner: {id: 1}}]]): Boolean
    nestedOutputs: [[OuterObject]]
    shapes: [Shape]
    recursive(input: RecursiveInputSlice): Boolean
    mapInput(input: Changes): Boolean
    collision: It
    invalidIdentifier: InvalidIdentifier
}
`)
