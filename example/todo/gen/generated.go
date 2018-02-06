// This file was generated by github.com/vektah/graphql-go, DO NOT EDIT

package gen

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/vektah/graphql-go/errors"
	"github.com/vektah/graphql-go/example/todo"
	"github.com/vektah/graphql-go/introspection"
	"github.com/vektah/graphql-go/jsonw"
	"github.com/vektah/graphql-go/query"
	"github.com/vektah/graphql-go/relay"
	"github.com/vektah/graphql-go/schema"
	"github.com/vektah/graphql-go/validation"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Resolvers interface {
	MyMutation_createTodo(ctx context.Context, text string) (todo.Todo, error)
	MyMutation_updateTodo(ctx context.Context, id int, changes map[string]interface{}) (*todo.Todo, error)
	MyQuery_todo(ctx context.Context, id int) (*todo.Todo, error)
	MyQuery_lastTodo(ctx context.Context) (*todo.Todo, error)
	MyQuery_todos(ctx context.Context) ([]todo.Todo, error)
}

var (
	myMutationSatisfies   = []string{"MyMutation"}
	myQuerySatisfies      = []string{"MyQuery"}
	todoSatisfies         = []string{"Todo"}
	__DirectiveSatisfies  = []string{"__Directive"}
	__EnumValueSatisfies  = []string{"__EnumValue"}
	__FieldSatisfies      = []string{"__Field"}
	__InputValueSatisfies = []string{"__InputValue"}
	__SchemaSatisfies     = []string{"__Schema"}
	__TypeSatisfies       = []string{"__Type"}
)

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _myMutation(sel []query.Selection, it *interface{}) {
	groupedFieldSet := ec.collectFields(sel, myMutationSatisfies, map[string]bool{})
	ec.json.BeginObject()
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "createTodo":
			ec.json.ObjectKey(field.Alias)
			var arg0 string
			if tmp, ok := field.Args["text"]; ok {
				tmp2, err := coerceString(tmp)
				if err != nil {
					ec.Error(err)
					continue
				}
				arg0 = tmp2
			}
			res, err := ec.resolvers.MyMutation_createTodo(ec.ctx, arg0)
			if err != nil {
				ec.Error(err)
				continue
			}
			ec._todo(field.Selections, &res)
			continue

		case "updateTodo":
			ec.json.ObjectKey(field.Alias)
			var arg0 int
			if tmp, ok := field.Args["id"]; ok {
				tmp2, err := coerceInt(tmp)
				if err != nil {
					ec.Error(err)
					continue
				}
				arg0 = tmp2
			}
			var arg1 map[string]interface{}
			if tmp, ok := field.Args["changes"]; ok {
				arg1 = tmp.(map[string]interface{})
			}
			res, err := ec.resolvers.MyMutation_updateTodo(ec.ctx, arg0, arg1)
			if err != nil {
				ec.Error(err)
				continue
			}
			if res == nil {
				ec.json.Null()
			} else {
				ec._todo(field.Selections, res)
			}
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	ec.json.EndObject()
}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _myQuery(sel []query.Selection, it *interface{}) {
	groupedFieldSet := ec.collectFields(sel, myQuerySatisfies, map[string]bool{})
	ec.json.BeginObject()
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "todo":
			ec.json.ObjectKey(field.Alias)
			var arg0 int
			if tmp, ok := field.Args["id"]; ok {
				tmp2, err := coerceInt(tmp)
				if err != nil {
					ec.Error(err)
					continue
				}
				arg0 = tmp2
			}
			res, err := ec.resolvers.MyQuery_todo(ec.ctx, arg0)
			if err != nil {
				ec.Error(err)
				continue
			}
			if res == nil {
				ec.json.Null()
			} else {
				ec._todo(field.Selections, res)
			}
			continue

		case "lastTodo":
			ec.json.ObjectKey(field.Alias)
			res, err := ec.resolvers.MyQuery_lastTodo(ec.ctx)
			if err != nil {
				ec.Error(err)
				continue
			}
			if res == nil {
				ec.json.Null()
			} else {
				ec._todo(field.Selections, res)
			}
			continue

		case "todos":
			ec.json.ObjectKey(field.Alias)
			res, err := ec.resolvers.MyQuery_todos(ec.ctx)
			if err != nil {
				ec.Error(err)
				continue
			}
			ec.json.BeginArray()
			for _, val := range res {
				ec._todo(field.Selections, &val)
			}
			ec.json.EndArray()
			continue

		case "__schema":
			ec.json.ObjectKey(field.Alias)
			res := ec.introspectSchema()
			if res == nil {
				ec.json.Null()
			} else {
				ec.___Schema(field.Selections, res)
			}
			continue

		case "__type":
			ec.json.ObjectKey(field.Alias)
			var arg0 string
			if tmp, ok := field.Args["name"]; ok {
				tmp2, err := coerceString(tmp)
				if err != nil {
					ec.Error(err)
					continue
				}
				arg0 = tmp2
			}
			res := ec.introspectType(arg0)
			if res == nil {
				ec.json.Null()
			} else {
				ec.___Type(field.Selections, res)
			}
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	ec.json.EndObject()
}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _todo(sel []query.Selection, it *todo.Todo) {
	groupedFieldSet := ec.collectFields(sel, todoSatisfies, map[string]bool{})
	ec.json.BeginObject()
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "id":
			ec.json.ObjectKey(field.Alias)
			res := it.ID
			ec.json.Int(res)
			continue

		case "text":
			ec.json.ObjectKey(field.Alias)
			res := it.Text
			ec.json.String(res)
			continue

		case "done":
			ec.json.ObjectKey(field.Alias)
			res := it.Done
			ec.json.Bool(res)
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	ec.json.EndObject()
}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(sel []query.Selection, it *introspection.Directive) {
	groupedFieldSet := ec.collectFields(sel, __DirectiveSatisfies, map[string]bool{})
	ec.json.BeginObject()
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "name":
			ec.json.ObjectKey(field.Alias)
			res := it.Name()
			ec.json.String(res)
			continue

		case "description":
			ec.json.ObjectKey(field.Alias)
			res := it.Description()
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.String(*res)
			}
			continue

		case "locations":
			ec.json.ObjectKey(field.Alias)
			res := it.Locations()
			ec.json.BeginArray()
			for _, val := range res {
				ec.json.String(val)
			}
			ec.json.EndArray()
			continue

		case "args":
			ec.json.ObjectKey(field.Alias)
			res := it.Args()
			ec.json.BeginArray()
			for _, val := range res {
				if val == nil {
					ec.json.Null()
				} else {
					ec.___InputValue(field.Selections, val)
				}
			}
			ec.json.EndArray()
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	ec.json.EndObject()
}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(sel []query.Selection, it *introspection.EnumValue) {
	groupedFieldSet := ec.collectFields(sel, __EnumValueSatisfies, map[string]bool{})
	ec.json.BeginObject()
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "name":
			ec.json.ObjectKey(field.Alias)
			res := it.Name()
			ec.json.String(res)
			continue

		case "description":
			ec.json.ObjectKey(field.Alias)
			res := it.Description()
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.String(*res)
			}
			continue

		case "isDeprecated":
			ec.json.ObjectKey(field.Alias)
			res := it.IsDeprecated()
			ec.json.Bool(res)
			continue

		case "deprecationReason":
			ec.json.ObjectKey(field.Alias)
			res := it.DeprecationReason()
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.String(*res)
			}
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	ec.json.EndObject()
}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(sel []query.Selection, it *introspection.Field) {
	groupedFieldSet := ec.collectFields(sel, __FieldSatisfies, map[string]bool{})
	ec.json.BeginObject()
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "name":
			ec.json.ObjectKey(field.Alias)
			res := it.Name()
			ec.json.String(res)
			continue

		case "description":
			ec.json.ObjectKey(field.Alias)
			res := it.Description()
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.String(*res)
			}
			continue

		case "args":
			ec.json.ObjectKey(field.Alias)
			res := it.Args()
			ec.json.BeginArray()
			for _, val := range res {
				if val == nil {
					ec.json.Null()
				} else {
					ec.___InputValue(field.Selections, val)
				}
			}
			ec.json.EndArray()
			continue

		case "type":
			ec.json.ObjectKey(field.Alias)
			res := it.Type()
			if res == nil {
				ec.json.Null()
			} else {
				ec.___Type(field.Selections, res)
			}
			continue

		case "isDeprecated":
			ec.json.ObjectKey(field.Alias)
			res := it.IsDeprecated()
			ec.json.Bool(res)
			continue

		case "deprecationReason":
			ec.json.ObjectKey(field.Alias)
			res := it.DeprecationReason()
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.String(*res)
			}
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	ec.json.EndObject()
}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(sel []query.Selection, it *introspection.InputValue) {
	groupedFieldSet := ec.collectFields(sel, __InputValueSatisfies, map[string]bool{})
	ec.json.BeginObject()
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "name":
			ec.json.ObjectKey(field.Alias)
			res := it.Name()
			ec.json.String(res)
			continue

		case "description":
			ec.json.ObjectKey(field.Alias)
			res := it.Description()
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.String(*res)
			}
			continue

		case "type":
			ec.json.ObjectKey(field.Alias)
			res := it.Type()
			if res == nil {
				ec.json.Null()
			} else {
				ec.___Type(field.Selections, res)
			}
			continue

		case "defaultValue":
			ec.json.ObjectKey(field.Alias)
			res := it.DefaultValue()
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.String(*res)
			}
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	ec.json.EndObject()
}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(sel []query.Selection, it *introspection.Schema) {
	groupedFieldSet := ec.collectFields(sel, __SchemaSatisfies, map[string]bool{})
	ec.json.BeginObject()
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "types":
			ec.json.ObjectKey(field.Alias)
			res := it.Types()
			ec.json.BeginArray()
			for _, val := range res {
				if val == nil {
					ec.json.Null()
				} else {
					ec.___Type(field.Selections, val)
				}
			}
			ec.json.EndArray()
			continue

		case "queryType":
			ec.json.ObjectKey(field.Alias)
			res := it.QueryType()
			if res == nil {
				ec.json.Null()
			} else {
				ec.___Type(field.Selections, res)
			}
			continue

		case "mutationType":
			ec.json.ObjectKey(field.Alias)
			res := it.MutationType()
			if res == nil {
				ec.json.Null()
			} else {
				ec.___Type(field.Selections, res)
			}
			continue

		case "subscriptionType":
			ec.json.ObjectKey(field.Alias)
			res := it.SubscriptionType()
			if res == nil {
				ec.json.Null()
			} else {
				ec.___Type(field.Selections, res)
			}
			continue

		case "directives":
			ec.json.ObjectKey(field.Alias)
			res := it.Directives()
			ec.json.BeginArray()
			for _, val := range res {
				if val == nil {
					ec.json.Null()
				} else {
					ec.___Directive(field.Selections, val)
				}
			}
			ec.json.EndArray()
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	ec.json.EndObject()
}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(sel []query.Selection, it *introspection.Type) {
	groupedFieldSet := ec.collectFields(sel, __TypeSatisfies, map[string]bool{})
	ec.json.BeginObject()
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "kind":
			ec.json.ObjectKey(field.Alias)
			res := it.Kind()
			ec.json.String(res)
			continue

		case "name":
			ec.json.ObjectKey(field.Alias)
			res := it.Name()
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.String(*res)
			}
			continue

		case "description":
			ec.json.ObjectKey(field.Alias)
			res := it.Description()
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.String(*res)
			}
			continue

		case "fields":
			ec.json.ObjectKey(field.Alias)
			var arg0 bool
			if tmp, ok := field.Args["includeDeprecated"]; ok {
				tmp2, err := coerceBool(tmp)
				if err != nil {
					ec.Error(err)
					continue
				}
				arg0 = tmp2
			}
			res := it.Fields(arg0)
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.BeginArray()
				for _, val := range *res {
					if val == nil {
						ec.json.Null()
					} else {
						ec.___Field(field.Selections, val)
					}
				}
				ec.json.EndArray()
			}
			continue

		case "interfaces":
			ec.json.ObjectKey(field.Alias)
			res := it.Interfaces()
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.BeginArray()
				for _, val := range *res {
					if val == nil {
						ec.json.Null()
					} else {
						ec.___Type(field.Selections, val)
					}
				}
				ec.json.EndArray()
			}
			continue

		case "possibleTypes":
			ec.json.ObjectKey(field.Alias)
			res := it.PossibleTypes()
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.BeginArray()
				for _, val := range *res {
					if val == nil {
						ec.json.Null()
					} else {
						ec.___Type(field.Selections, val)
					}
				}
				ec.json.EndArray()
			}
			continue

		case "enumValues":
			ec.json.ObjectKey(field.Alias)
			var arg0 bool
			if tmp, ok := field.Args["includeDeprecated"]; ok {
				tmp2, err := coerceBool(tmp)
				if err != nil {
					ec.Error(err)
					continue
				}
				arg0 = tmp2
			}
			res := it.EnumValues(arg0)
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.BeginArray()
				for _, val := range *res {
					if val == nil {
						ec.json.Null()
					} else {
						ec.___EnumValue(field.Selections, val)
					}
				}
				ec.json.EndArray()
			}
			continue

		case "inputFields":
			ec.json.ObjectKey(field.Alias)
			res := it.InputFields()
			if res == nil {
				ec.json.Null()
			} else {
				ec.json.BeginArray()
				for _, val := range *res {
					if val == nil {
						ec.json.Null()
					} else {
						ec.___InputValue(field.Selections, val)
					}
				}
				ec.json.EndArray()
			}
			continue

		case "ofType":
			ec.json.ObjectKey(field.Alias)
			res := it.OfType()
			if res == nil {
				ec.json.Null()
			} else {
				ec.___Type(field.Selections, res)
			}
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	ec.json.EndObject()
}

var parsedSchema = schema.MustParse("schema {\n\tquery: MyQuery\n\tmutation: MyMutation\n}\n\ntype MyQuery {\n\ttodo(id: Int!): Todo\n\tlastTodo: Todo\n\ttodos: [Todo!]!\n}\n\ntype MyMutation {\n\tcreateTodo(text: String!): Todo!\n\tupdateTodo(id: Int!, changes: TodoInput!): Todo\n}\n\ntype Todo {\n\tid: Int!\n\ttext: String!\n\tdone: Boolean!\n}\n\ninput TodoInput {\n\ttext: String\n\tdone: Boolean\n}\n")
var _ = fmt.Print

func NewResolver(resolvers Resolvers) relay.Resolver {
	return func(ctx context.Context, document string, operationName string, variables map[string]interface{}, w io.Writer) []*errors.QueryError {
		doc, qErr := query.Parse(document)
		if qErr != nil {
			return []*errors.QueryError{qErr}
		}

		errs := validation.Validate(parsedSchema, doc)
		if len(errs) != 0 {
			return errs
		}

		op, err := doc.GetOperation(operationName)
		if err != nil {
			return []*errors.QueryError{errors.Errorf("%s", err)}
		}

		if op.Type != query.Query && op.Type != query.Mutation {
			return []*errors.QueryError{errors.Errorf("unsupported operation type")}
		}

		c := executionContext{
			resolvers: resolvers,
			variables: variables,
			doc:       doc,
			ctx:       ctx,
			json:      jsonw.New(w),
		}

		// TODO: parallelize if query.

		c.json.BeginObject()

		c.json.ObjectKey("data")

		if op.Type == query.Query {
			c._myQuery(op.Selections, nil)
		} else if op.Type == query.Mutation {
			c._myMutation(op.Selections, nil)
		} else {
			c.Errorf("unsupported operation %s", op.Type)
			c.json.Null()
		}

		if len(c.Errors) > 0 {
			c.json.ObjectKey("errors")
			errors.WriteErrors(w, c.Errors)
		}

		c.json.EndObject()
		return nil
	}
}

type executionContext struct {
	errors.Builder
	json      *jsonw.Writer
	resolvers Resolvers
	variables map[string]interface{}
	doc       *query.Document
	ctx       context.Context
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

func instanceOf(val string, satisfies []string) bool {
	for _, s := range satisfies {
		if val == s {
			return true
		}
	}
	return false
}

func (ec *executionContext) collectFields(selSet []query.Selection, satisfies []string, visited map[string]bool) []collectedField {
	var groupedFields []collectedField

	for _, sel := range selSet {
		switch sel := sel.(type) {
		case *query.Field:
			f := getOrCreateField(&groupedFields, sel.Name.Name, func() collectedField {
				f := collectedField{
					Alias: sel.Alias.Name,
					Name:  sel.Name.Name,
				}
				if len(sel.Arguments) > 0 {
					f.Args = map[string]interface{}{}
					for _, arg := range sel.Arguments {
						f.Args[arg.Name.Name] = arg.Value.Value(ec.variables)
					}
				}
				return f
			})

			f.Selections = append(f.Selections, sel.Selections...)
		case *query.InlineFragment:
			if !instanceOf(sel.On.Ident.Name, satisfies) {
				continue
			}

			for _, childField := range ec.collectFields(sel.Selections, satisfies, visited) {
				f := getOrCreateField(&groupedFields, childField.Name, func() collectedField { return childField })
				f.Selections = append(f.Selections, childField.Selections...)
			}

		case *query.FragmentSpread:
			fragmentName := sel.Name.Name
			if _, seen := visited[fragmentName]; seen {
				continue
			}
			visited[fragmentName] = true

			fragment := ec.doc.Fragments.Get(fragmentName)
			if fragment == nil {
				ec.Errorf("missing fragment %s", fragmentName)
				continue
			}

			if !instanceOf(fragment.On.Ident.Name, satisfies) {
				continue
			}

			for _, childField := range ec.collectFields(fragment.Selections, satisfies, visited) {
				f := getOrCreateField(&groupedFields, childField.Name, func() collectedField { return childField })
				f.Selections = append(f.Selections, childField.Selections...)
			}

		default:
			panic(fmt.Errorf("unsupported %T", sel))
		}
	}

	return groupedFields
}

type collectedField struct {
	Alias      string
	Name       string
	Args       map[string]interface{}
	Selections []query.Selection
}

func decodeHook(sourceType reflect.Type, destType reflect.Type, value interface{}) (interface{}, error) {
	if destType.PkgPath() == "time" && destType.Name() == "Time" {
		if dateStr, ok := value.(string); ok {
			return time.Parse(time.RFC3339, dateStr)
		}
		return nil, errors.Errorf("time should be an RFC3339 formatted string")
	}
	return value, nil
}

// nolint: deadcode, megacheck
func unpackComplexArg(result interface{}, data interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:     "graphql",
		ErrorUnused: true,
		Result:      result,
		DecodeHook:  decodeHook,
	})
	if err != nil {
		panic(err)
	}

	return decoder.Decode(data)
}

func getOrCreateField(c *[]collectedField, name string, creator func() collectedField) *collectedField {
	for i, cf := range *c {
		if cf.Alias == name {
			return &(*c)[i]
		}
	}

	f := creator()

	*c = append(*c, f)
	return &(*c)[len(*c)-1]
}

// nolint: deadcode, megacheck
func coerceString(v interface{}) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case int:
		return strconv.Itoa(v), nil
	case float64:
		return fmt.Sprintf("%f", v), nil
	case bool:
		if v {
			return "true", nil
		} else {
			return "false", nil
		}
	case nil:
		return "null", nil
	default:
		return "", fmt.Errorf("%T is not a string", v)
	}
}

// nolint: deadcode, megacheck
func coerceBool(v interface{}) (bool, error) {
	switch v := v.(type) {
	case string:
		return "true" == strings.ToLower(v), nil
	case int:
		return v != 0, nil
	default:
		return false, fmt.Errorf("%T is not a bool", v)
	}
}

// nolint: deadcode, megacheck
func coerceInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case string:
		return strconv.Atoi(v)
	case int:
		return v, nil
	case float64:
		return int(v), nil
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}

// nolint: deadcode, megacheck
func coercefloat64(v interface{}) (float64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("%T is not an float", v)
	}
}
