package todo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

// region    ************************** generated!.gotpl **************************

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	MyMutation() MyMutationResolver
	MyQuery() MyQueryResolver
}

type DirectiveRoot struct {
	HasRole func(ctx context.Context, obj interface{}, next graphql.Resolver, role Role) (res interface{}, err error)

	User func(ctx context.Context, obj interface{}, next graphql.Resolver, id int) (res interface{}, err error)
}

type ComplexityRoot struct {
	MyMutation struct {
		CreateTodo func(childComplexity int, todo TodoInput) int
		UpdateTodo func(childComplexity int, id int, changes map[string]interface{}) int
	}

	MyQuery struct {
		LastTodo func(childComplexity int) int
		Todo     func(childComplexity int, id int) int
		Todos    func(childComplexity int) int
	}

	Todo struct {
		Done func(childComplexity int) int
		ID   func(childComplexity int) int
		Text func(childComplexity int) int
	}
}

type MyMutationResolver interface {
	CreateTodo(ctx context.Context, todo TodoInput) (*Todo, error)
	UpdateTodo(ctx context.Context, id int, changes map[string]interface{}) (*Todo, error)
}
type MyQueryResolver interface {
	Todo(ctx context.Context, id int) (*Todo, error)
	LastTodo(ctx context.Context) (*Todo, error)
	Todos(ctx context.Context) ([]*Todo, error)
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "MyMutation.createTodo":
		if e.complexity.MyMutation.CreateTodo == nil {
			break
		}

		args, err := ec.field_MyMutation_createTodo_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.MyMutation.CreateTodo(childComplexity, args["todo"].(TodoInput)), true

	case "MyMutation.updateTodo":
		if e.complexity.MyMutation.UpdateTodo == nil {
			break
		}

		args, err := ec.field_MyMutation_updateTodo_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.MyMutation.UpdateTodo(childComplexity, args["id"].(int), args["changes"].(map[string]interface{})), true

	case "MyQuery.lastTodo":
		if e.complexity.MyQuery.LastTodo == nil {
			break
		}

		return e.complexity.MyQuery.LastTodo(childComplexity), true

	case "MyQuery.todo":
		if e.complexity.MyQuery.Todo == nil {
			break
		}

		args, err := ec.field_MyQuery_todo_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.MyQuery.Todo(childComplexity, args["id"].(int)), true

	case "MyQuery.todos":
		if e.complexity.MyQuery.Todos == nil {
			break
		}

		return e.complexity.MyQuery.Todos(childComplexity), true

	case "Todo.done":
		if e.complexity.Todo.Done == nil {
			break
		}

		return e.complexity.Todo.Done(childComplexity), true

	case "Todo.id":
		if e.complexity.Todo.ID == nil {
			break
		}

		return e.complexity.Todo.ID(childComplexity), true

	case "Todo.text":
		if e.complexity.Todo.Text == nil {
			break
		}

		return e.complexity.Todo.Text(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			data := ec._queryMiddleware(ctx, rc.Operation, func(ctx context.Context) (interface{}, error) {
				return ec._MyQuery(ctx, rc.Operation.SelectionSet), nil
			})
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			data := ec._mutationMiddleware(ctx, rc.Operation, func(ctx context.Context) (interface{}, error) {
				return ec._MyMutation(ctx, rc.Operation.SelectionSet), nil
			})
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var parsedSchema = gqlparser.MustLoadSchema(
	&ast.Source{Name: "schema.graphql", Input: `schema {
    query: MyQuery
    mutation: MyMutation
}

type MyQuery {
    todo(id: ID!): Todo
    lastTodo: Todo
    todos: [Todo!]!
}

type MyMutation {
    createTodo(todo: TodoInput!): Todo!
    updateTodo(id: ID!, changes: Map!): Todo
}

type Todo {
    id: ID!
    text: String!
    done: Boolean! @hasRole(role: OWNER) # only the owner can see if a todo is done
}

"Passed to createTodo to create a new todo"
input TodoInput {
    "The body text"
    text: String!
    "Is it done already?"
    done: Boolean
}

scalar Map

"Prevents access to a field if the user doesnt have the matching role"
directive @hasRole(role: Role!) on FIELD_DEFINITION
directive @user(id: ID!) on MUTATION | QUERY | FIELD

enum Role {
    ADMIN
    OWNER
}
`},
)

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) dir_hasRole_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 Role
	if tmp, ok := rawArgs["role"]; ok {
		arg0, err = ec.unmarshalNRole2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐRole(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["role"] = arg0
	return args, nil
}

func (ec *executionContext) dir_user_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalNID2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_MyMutation_createTodo_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 TodoInput
	if tmp, ok := rawArgs["todo"]; ok {
		arg0, err = ec.unmarshalNTodoInput2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodoInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["todo"] = arg0
	return args, nil
}

func (ec *executionContext) field_MyMutation_updateTodo_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalNID2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	var arg1 map[string]interface{}
	if tmp, ok := rawArgs["changes"]; ok {
		arg1, err = ec.unmarshalNMap2map(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["changes"] = arg1
	return args, nil
}

func (ec *executionContext) field_MyQuery___type_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	return args, nil
}

func (ec *executionContext) field_MyQuery_todo_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["id"]; ok {
		arg0, err = ec.unmarshalNID2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_enumValues_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		arg0, err = ec.unmarshalOBoolean2bool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_fields_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		arg0, err = ec.unmarshalOBoolean2bool(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

func (ec *executionContext) _queryMiddleware(ctx context.Context, obj *ast.OperationDefinition, next func(ctx context.Context) (interface{}, error)) graphql.Marshaler {

	for _, d := range obj.Directives {
		switch d.Name {
		case "user":
			rawArgs := d.ArgumentMap(ec.Variables)
			args, err := ec.dir_user_args(ctx, rawArgs)
			if err != nil {
				ec.Error(ctx, err)
				return graphql.Null
			}
			n := next
			next = func(ctx context.Context) (interface{}, error) {
				if ec.directives.User == nil {
					return nil, errors.New("directive user is not implemented")
				}
				return ec.directives.User(ctx, obj, n, args["id"].(int))
			}
		}
	}
	tmp, err := next(ctx)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if data, ok := tmp.(graphql.Marshaler); ok {
		return data
	}
	ec.Errorf(ctx, `unexpected type %T from directive, should be graphql.Marshaler`, tmp)
	return graphql.Null

}

func (ec *executionContext) _mutationMiddleware(ctx context.Context, obj *ast.OperationDefinition, next func(ctx context.Context) (interface{}, error)) graphql.Marshaler {

	for _, d := range obj.Directives {
		switch d.Name {
		case "user":
			rawArgs := d.ArgumentMap(ec.Variables)
			args, err := ec.dir_user_args(ctx, rawArgs)
			if err != nil {
				ec.Error(ctx, err)
				return graphql.Null
			}
			n := next
			next = func(ctx context.Context) (interface{}, error) {
				if ec.directives.User == nil {
					return nil, errors.New("directive user is not implemented")
				}
				return ec.directives.User(ctx, obj, n, args["id"].(int))
			}
		}
	}
	tmp, err := next(ctx)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if data, ok := tmp.(graphql.Marshaler); ok {
		return data
	}
	ec.Errorf(ctx, `unexpected type %T from directive, should be graphql.Marshaler`, tmp)
	return graphql.Null

}

func (ec *executionContext) _fieldMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver) interface{} {
	fc := graphql.GetFieldContext(ctx)
	for _, d := range fc.Field.Directives {
		switch d.Name {
		case "user":
			rawArgs := d.ArgumentMap(ec.Variables)
			args, err := ec.dir_user_args(ctx, rawArgs)
			if err != nil {
				ec.Error(ctx, err)
				return nil
			}
			n := next
			next = func(ctx context.Context) (interface{}, error) {
				if ec.directives.User == nil {
					return nil, errors.New("directive user is not implemented")
				}
				return ec.directives.User(ctx, obj, n, args["id"].(int))
			}
		}
	}
	res, err := ec.ResolverMiddleware(ctx, next)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	return res
}

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _MyMutation_createTodo(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "MyMutation",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_MyMutation_createTodo_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.MyMutation().CreateTodo(rctx, args["todo"].(TodoInput))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*Todo)
	fc.Result = res
	return ec.marshalNTodo2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodo(ctx, field.Selections, res)
}

func (ec *executionContext) _MyMutation_updateTodo(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "MyMutation",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_MyMutation_updateTodo_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.MyMutation().UpdateTodo(rctx, args["id"].(int), args["changes"].(map[string]interface{}))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Todo)
	fc.Result = res
	return ec.marshalOTodo2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodo(ctx, field.Selections, res)
}

func (ec *executionContext) _MyQuery_todo(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "MyQuery",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_MyQuery_todo_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.MyQuery().Todo(rctx, args["id"].(int))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Todo)
	fc.Result = res
	return ec.marshalOTodo2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodo(ctx, field.Selections, res)
}

func (ec *executionContext) _MyQuery_lastTodo(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "MyQuery",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.MyQuery().LastTodo(rctx)
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Todo)
	fc.Result = res
	return ec.marshalOTodo2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodo(ctx, field.Selections, res)
}

func (ec *executionContext) _MyQuery_todos(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "MyQuery",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.MyQuery().Todos(rctx)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*Todo)
	fc.Result = res
	return ec.marshalNTodo2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodo(ctx, field.Selections, res)
}

func (ec *executionContext) _MyQuery___type(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "MyQuery",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_MyQuery___type_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(args["name"].(string))
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) _MyQuery___schema(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "MyQuery",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	fc.Result = res
	return ec.marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx, field.Selections, res)
}

func (ec *executionContext) _Todo_id(ctx context.Context, field graphql.CollectedField, obj *Todo) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Todo",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNID2int(ctx, field.Selections, res)
}

func (ec *executionContext) _Todo_text(ctx context.Context, field graphql.CollectedField, obj *Todo) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Todo",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Text, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Todo_done(ctx context.Context, field graphql.CollectedField, obj *Todo) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "Todo",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		directive0 := func(rctx context.Context) (interface{}, error) {
			ctx = rctx // use context from middleware stack in children
			return obj.Done, nil
		}
		directive1 := func(ctx context.Context) (interface{}, error) {
			role, err := ec.unmarshalNRole2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐRole(ctx, "OWNER")
			if err != nil {
				return nil, err
			}
			if ec.directives.HasRole == nil {
				return nil, errors.New("directive hasRole is not implemented")
			}
			return ec.directives.HasRole(ctx, obj, directive0, role)
		}

		tmp, err := directive1(rctx)
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return nil, nil
		}
		if data, ok := tmp.(bool); ok {
			return data, nil
		}
		return nil, fmt.Errorf(`unexpected type %T from directive, should be bool`, tmp)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Locations, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	fc.Result = res
	return ec.marshalN__DirectiveLocation2ᚕstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Directive",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	fc.Result = res
	return ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__EnumValue",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	fc.Result = res
	return ec.marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Field",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__InputValue",
		Field:    field,
		Args:     nil,
		IsMethod: false,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DefaultValue, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Types(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.QueryType(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.MutationType(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SubscriptionType(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Schema",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Directives(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Directive)
	fc.Result = res
	return ec.marshalN__Directive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Kind(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalN__TypeKind2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_fields_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Fields(args["includeDeprecated"].(bool)), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Field)
	fc.Result = res
	return ec.marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Interfaces(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PossibleTypes(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_enumValues_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EnumValues(args["includeDeprecated"].(bool)), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.EnumValue)
	fc.Result = res
	return ec.marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.InputFields(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	fc.Result = res
	return ec.marshalO__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) (ret graphql.Marshaler) {
	fc := &graphql.FieldContext{
		Object:   "__Type",
		Field:    field,
		Args:     nil,
		IsMethod: true,
		Stats:    graphql.FieldStats{Started: graphql.Now()},
	}
	defer func() {
		fc.Stats.Completed = graphql.Now()
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	fc.Stats.ArgumentsCompleted = graphql.Now()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OfType(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputTodoInput(ctx context.Context, obj interface{}) (TodoInput, error) {
	var it TodoInput
	var asMap = obj.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "text":
			var err error
			it.Text, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "done":
			var err error
			it.Done, err = ec.unmarshalOBoolean2ᚖbool(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var myMutationImplementors = []string{"MyMutation"}

func (ec *executionContext) _MyMutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, myMutationImplementors)

	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "MyMutation",
		Stats:  graphql.FieldStats{Started: graphql.Now()},
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("MyMutation")
		case "createTodo":
			out.Values[i] = ec._MyMutation_createTodo(ctx, field)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "updateTodo":
			out.Values[i] = ec._MyMutation_updateTodo(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var myQueryImplementors = []string{"MyQuery"}

func (ec *executionContext) _MyQuery(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, myQueryImplementors)

	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "MyQuery",
		Stats:  graphql.FieldStats{Started: graphql.Now()},
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("MyQuery")
		case "todo":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._MyQuery_todo(ctx, field)
				return res
			})
		case "lastTodo":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._MyQuery_lastTodo(ctx, field)
				return res
			})
		case "todos":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._MyQuery_todos(ctx, field)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		case "__type":
			out.Values[i] = ec._MyQuery___type(ctx, field)
		case "__schema":
			out.Values[i] = ec._MyQuery___schema(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var todoImplementors = []string{"Todo"}

func (ec *executionContext) _Todo(ctx context.Context, sel ast.SelectionSet, obj *Todo) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, todoImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Todo")
		case "id":
			out.Values[i] = ec._Todo_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "text":
			out.Values[i] = ec._Todo_text(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "done":
			out.Values[i] = ec._Todo_done(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __DirectiveImplementors = []string{"__Directive"}

func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __DirectiveImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __EnumValueImplementors = []string{"__EnumValue"}

func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __EnumValueImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __FieldImplementors = []string{"__Field"}

func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __FieldImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __InputValueImplementors = []string{"__InputValue"}

func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __InputValueImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __SchemaImplementors = []string{"__Schema"}

func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __SchemaImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var __TypeImplementors = []string{"__Type"}

func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __TypeImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
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
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	return graphql.UnmarshalBoolean(v)
}

func (ec *executionContext) marshalNBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	res := graphql.MarshalBoolean(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNID2int(ctx context.Context, v interface{}) (int, error) {
	return graphql.UnmarshalIntID(v)
}

func (ec *executionContext) marshalNID2int(ctx context.Context, sel ast.SelectionSet, v int) graphql.Marshaler {
	res := graphql.MarshalIntID(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNMap2map(ctx context.Context, v interface{}) (map[string]interface{}, error) {
	if v == nil {
		return nil, nil
	}
	return graphql.UnmarshalMap(v)
}

func (ec *executionContext) marshalNMap2map(ctx context.Context, sel ast.SelectionSet, v map[string]interface{}) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := graphql.MarshalMap(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNRole2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐRole(ctx context.Context, v interface{}) (Role, error) {
	var res Role
	return res, res.UnmarshalGQL(v)
}

func (ec *executionContext) marshalNRole2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐRole(ctx context.Context, sel ast.SelectionSet, v Role) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalNString2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalNString2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) marshalNTodo2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodo(ctx context.Context, sel ast.SelectionSet, v Todo) graphql.Marshaler {
	return ec._Todo(ctx, sel, &v)
}

func (ec *executionContext) marshalNTodo2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodo(ctx context.Context, sel ast.SelectionSet, v []*Todo) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNTodo2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodo(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalNTodo2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodo(ctx context.Context, sel ast.SelectionSet, v *Todo) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._Todo(ctx, sel, v)
}

func (ec *executionContext) unmarshalNTodoInput2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodoInput(ctx context.Context, v interface{}) (TodoInput, error) {
	return ec.unmarshalInputTodoInput(ctx, v)
}

func (ec *executionContext) marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx context.Context, sel ast.SelectionSet, v introspection.Directive) graphql.Marshaler {
	return ec.___Directive(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Directive2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx context.Context, sel ast.SelectionSet, v []introspection.Directive) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Directive2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐDirective(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) unmarshalN__DirectiveLocation2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalN__DirectiveLocation2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalN__DirectiveLocation2ᚕstring(ctx context.Context, v interface{}) ([]string, error) {
	var vSlice []interface{}
	if v != nil {
		if tmp1, ok := v.([]interface{}); ok {
			vSlice = tmp1
		} else {
			vSlice = []interface{}{v}
		}
	}
	var err error
	res := make([]string, len(vSlice))
	for i := range vSlice {
		res[i], err = ec.unmarshalN__DirectiveLocation2string(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalN__DirectiveLocation2ᚕstring(ctx context.Context, sel ast.SelectionSet, v []string) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__DirectiveLocation2string(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx context.Context, sel ast.SelectionSet, v introspection.EnumValue) graphql.Marshaler {
	return ec.___EnumValue(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Field2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx context.Context, sel ast.SelectionSet, v introspection.Field) graphql.Marshaler {
	return ec.___Field(ctx, sel, &v)
}

func (ec *executionContext) marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v introspection.InputValue) graphql.Marshaler {
	return ec.___InputValue(ctx, sel, &v)
}

func (ec *executionContext) marshalN__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v introspection.Type) graphql.Marshaler {
	return ec.___Type(ctx, sel, &v)
}

func (ec *executionContext) marshalN__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v []introspection.Type) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalN__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

func (ec *executionContext) unmarshalN__TypeKind2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalN__TypeKind2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) unmarshalOBoolean2bool(ctx context.Context, v interface{}) (bool, error) {
	return graphql.UnmarshalBoolean(v)
}

func (ec *executionContext) marshalOBoolean2bool(ctx context.Context, sel ast.SelectionSet, v bool) graphql.Marshaler {
	return graphql.MarshalBoolean(v)
}

func (ec *executionContext) unmarshalOBoolean2ᚖbool(ctx context.Context, v interface{}) (*bool, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOBoolean2bool(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOBoolean2ᚖbool(ctx context.Context, sel ast.SelectionSet, v *bool) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOBoolean2bool(ctx, sel, *v)
}

func (ec *executionContext) unmarshalOString2string(ctx context.Context, v interface{}) (string, error) {
	return graphql.UnmarshalString(v)
}

func (ec *executionContext) marshalOString2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	return graphql.MarshalString(v)
}

func (ec *executionContext) unmarshalOString2ᚖstring(ctx context.Context, v interface{}) (*string, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalOString2string(ctx, v)
	return &res, err
}

func (ec *executionContext) marshalOString2ᚖstring(ctx context.Context, sel ast.SelectionSet, v *string) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.marshalOString2string(ctx, sel, *v)
}

func (ec *executionContext) marshalOTodo2githubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodo(ctx context.Context, sel ast.SelectionSet, v Todo) graphql.Marshaler {
	return ec._Todo(ctx, sel, &v)
}

func (ec *executionContext) marshalOTodo2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋexampleᚋtodoᚐTodo(ctx context.Context, sel ast.SelectionSet, v *Todo) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Todo(ctx, sel, v)
}

func (ec *executionContext) marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx context.Context, sel ast.SelectionSet, v []introspection.EnumValue) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__EnumValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Field2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx context.Context, sel ast.SelectionSet, v []introspection.Field) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Field2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐField(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__InputValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx context.Context, sel ast.SelectionSet, v []introspection.InputValue) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__InputValue2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐInputValue(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Schema2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx context.Context, sel ast.SelectionSet, v introspection.Schema) graphql.Marshaler {
	return ec.___Schema(ctx, sel, &v)
}

func (ec *executionContext) marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx context.Context, sel ast.SelectionSet, v *introspection.Schema) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Schema(ctx, sel, v)
}

func (ec *executionContext) marshalO__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v introspection.Type) graphql.Marshaler {
	return ec.___Type(ctx, sel, &v)
}

func (ec *executionContext) marshalO__Type2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v []introspection.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
			Stats:  graphql.FieldStats{Started: graphql.Now()},
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalN__Type2githubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx context.Context, sel ast.SelectionSet, v *introspection.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
