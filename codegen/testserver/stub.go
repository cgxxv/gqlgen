// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package testserver

import (
	"context"

	introspection1 "github.com/99designs/gqlgen/codegen/testserver/introspection"
	"github.com/99designs/gqlgen/codegen/testserver/invalid-packagename"
)

type Stub struct {
	ForcedResolverResolver struct {
		Field func(ctx context.Context, obj *ForcedResolver) (*Circle, error)
	}
	ModelMethodsResolver struct {
		ResolverField func(ctx context.Context, obj *ModelMethods) (bool, error)
	}
	PanicsResolver struct {
		FieldScalarMarshal func(ctx context.Context, obj *Panics) ([]MarshalPanic, error)
		ArgUnmarshal       func(ctx context.Context, obj *Panics, u []MarshalPanic) (bool, error)
	}
	QueryResolver struct {
		InvalidIdentifier      func(ctx context.Context) (*invalid_packagename.InvalidIdentifier, error)
		Collision              func(ctx context.Context) (*introspection1.It, error)
		MapInput               func(ctx context.Context, input map[string]interface{}) (*bool, error)
		Recursive              func(ctx context.Context, input *RecursiveInputSlice) (*bool, error)
		NestedInputs           func(ctx context.Context, input [][]*OuterInput) (*bool, error)
		NestedOutputs          func(ctx context.Context) ([][]*OuterObject, error)
		Shapes                 func(ctx context.Context) ([]Shape, error)
		ErrorBubble            func(ctx context.Context) (*Error, error)
		ModelMethods           func(ctx context.Context) (*ModelMethods, error)
		Valid                  func(ctx context.Context) (string, error)
		User                   func(ctx context.Context, id int) (*User, error)
		NullableArg            func(ctx context.Context, arg *int) (*string, error)
		DirectiveArg           func(ctx context.Context, arg string) (*string, error)
		DirectiveNullableArg   func(ctx context.Context, arg *int, arg2 *int) (*string, error)
		DirectiveInputNullable func(ctx context.Context, arg *InputDirectives) (*string, error)
		DirectiveInput         func(ctx context.Context, arg InputDirectives) (*string, error)
		InputSlice             func(ctx context.Context, arg []string) (bool, error)
		ShapeUnion             func(ctx context.Context) (ShapeUnion, error)
		Autobind               func(ctx context.Context) (*Autobind, error)
		Panics                 func(ctx context.Context) (*Panics, error)
		ValidType              func(ctx context.Context) (*ValidType, error)
	}
	SubscriptionResolver struct {
		Updated     func(ctx context.Context) (<-chan string, error)
		InitPayload func(ctx context.Context) (<-chan string, error)
	}
	UserResolver struct {
		Friends func(ctx context.Context, obj *User) ([]User, error)
	}
}

func (r *Stub) ForcedResolver() ForcedResolverResolver {
	return &stubForcedResolver{r}
}
func (r *Stub) ModelMethods() ModelMethodsResolver {
	return &stubModelMethods{r}
}
func (r *Stub) Panics() PanicsResolver {
	return &stubPanics{r}
}
func (r *Stub) Query() QueryResolver {
	return &stubQuery{r}
}
func (r *Stub) Subscription() SubscriptionResolver {
	return &stubSubscription{r}
}
func (r *Stub) User() UserResolver {
	return &stubUser{r}
}

type stubForcedResolver struct{ *Stub }

func (r *stubForcedResolver) Field(ctx context.Context, obj *ForcedResolver) (*Circle, error) {
	return r.ForcedResolverResolver.Field(ctx, obj)
}

type stubModelMethods struct{ *Stub }

func (r *stubModelMethods) ResolverField(ctx context.Context, obj *ModelMethods) (bool, error) {
	return r.ModelMethodsResolver.ResolverField(ctx, obj)
}

type stubPanics struct{ *Stub }

func (r *stubPanics) FieldScalarMarshal(ctx context.Context, obj *Panics) ([]MarshalPanic, error) {
	return r.PanicsResolver.FieldScalarMarshal(ctx, obj)
}
func (r *stubPanics) ArgUnmarshal(ctx context.Context, obj *Panics, u []MarshalPanic) (bool, error) {
	return r.PanicsResolver.ArgUnmarshal(ctx, obj, u)
}

type stubQuery struct{ *Stub }

func (r *stubQuery) InvalidIdentifier(ctx context.Context) (*invalid_packagename.InvalidIdentifier, error) {
	return r.QueryResolver.InvalidIdentifier(ctx)
}
func (r *stubQuery) Collision(ctx context.Context) (*introspection1.It, error) {
	return r.QueryResolver.Collision(ctx)
}
func (r *stubQuery) MapInput(ctx context.Context, input map[string]interface{}) (*bool, error) {
	return r.QueryResolver.MapInput(ctx, input)
}
func (r *stubQuery) Recursive(ctx context.Context, input *RecursiveInputSlice) (*bool, error) {
	return r.QueryResolver.Recursive(ctx, input)
}
func (r *stubQuery) NestedInputs(ctx context.Context, input [][]*OuterInput) (*bool, error) {
	return r.QueryResolver.NestedInputs(ctx, input)
}
func (r *stubQuery) NestedOutputs(ctx context.Context) ([][]*OuterObject, error) {
	return r.QueryResolver.NestedOutputs(ctx)
}
func (r *stubQuery) Shapes(ctx context.Context) ([]Shape, error) {
	return r.QueryResolver.Shapes(ctx)
}
func (r *stubQuery) ErrorBubble(ctx context.Context) (*Error, error) {
	return r.QueryResolver.ErrorBubble(ctx)
}
func (r *stubQuery) ModelMethods(ctx context.Context) (*ModelMethods, error) {
	return r.QueryResolver.ModelMethods(ctx)
}
func (r *stubQuery) Valid(ctx context.Context) (string, error) {
	return r.QueryResolver.Valid(ctx)
}
func (r *stubQuery) User(ctx context.Context, id int) (*User, error) {
	return r.QueryResolver.User(ctx, id)
}
func (r *stubQuery) NullableArg(ctx context.Context, arg *int) (*string, error) {
	return r.QueryResolver.NullableArg(ctx, arg)
}
func (r *stubQuery) DirectiveArg(ctx context.Context, arg string) (*string, error) {
	return r.QueryResolver.DirectiveArg(ctx, arg)
}
func (r *stubQuery) DirectiveNullableArg(ctx context.Context, arg *int, arg2 *int) (*string, error) {
	return r.QueryResolver.DirectiveNullableArg(ctx, arg, arg2)
}
func (r *stubQuery) DirectiveInputNullable(ctx context.Context, arg *InputDirectives) (*string, error) {
	return r.QueryResolver.DirectiveInputNullable(ctx, arg)
}
func (r *stubQuery) DirectiveInput(ctx context.Context, arg InputDirectives) (*string, error) {
	return r.QueryResolver.DirectiveInput(ctx, arg)
}
func (r *stubQuery) InputSlice(ctx context.Context, arg []string) (bool, error) {
	return r.QueryResolver.InputSlice(ctx, arg)
}
func (r *stubQuery) ShapeUnion(ctx context.Context) (ShapeUnion, error) {
	return r.QueryResolver.ShapeUnion(ctx)
}
func (r *stubQuery) Autobind(ctx context.Context) (*Autobind, error) {
	return r.QueryResolver.Autobind(ctx)
}
func (r *stubQuery) Panics(ctx context.Context) (*Panics, error) {
	return r.QueryResolver.Panics(ctx)
}
func (r *stubQuery) ValidType(ctx context.Context) (*ValidType, error) {
	return r.QueryResolver.ValidType(ctx)
}

type stubSubscription struct{ *Stub }

func (r *stubSubscription) Updated(ctx context.Context) (<-chan string, error) {
	return r.SubscriptionResolver.Updated(ctx)
}
func (r *stubSubscription) InitPayload(ctx context.Context) (<-chan string, error) {
	return r.SubscriptionResolver.InitPayload(ctx)
}

type stubUser struct{ *Stub }

func (r *stubUser) Friends(ctx context.Context, obj *User) ([]User, error) {
	return r.UserResolver.Friends(ctx, obj)
}
