// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package testserver

import (
	"context"

	introspection1 "github.com/99designs/gqlgen/codegen/testserver/introspection"
	"github.com/99designs/gqlgen/codegen/testserver/invalid-packagename"
)

type Stub struct {
	ErrorsResolver struct {
		A func(ctx context.Context, obj *Errors) (*Error, error)
		B func(ctx context.Context, obj *Errors) (*Error, error)
		C func(ctx context.Context, obj *Errors) (*Error, error)
		D func(ctx context.Context, obj *Errors) (*Error, error)
		E func(ctx context.Context, obj *Errors) (*Error, error)
	}
	ForcedResolverResolver struct {
		Field func(ctx context.Context, obj *ForcedResolver) (*Circle, error)
	}
	ModelMethodsResolver struct {
		ResolverField func(ctx context.Context, obj *ModelMethods) (bool, error)
	}
	OverlappingFieldsResolver struct {
		OldFoo func(ctx context.Context, obj *OverlappingFields) (int, error)
	}
	PanicsResolver struct {
		FieldScalarMarshal func(ctx context.Context, obj *Panics) ([]MarshalPanic, error)
		ArgUnmarshal       func(ctx context.Context, obj *Panics, u []MarshalPanic) (bool, error)
	}
	PrimitiveResolver struct {
		Value func(ctx context.Context, obj *Primitive) (int, error)
	}
	PrimitiveStringResolver struct {
		Value func(ctx context.Context, obj *PrimitiveString) (string, error)
		Len   func(ctx context.Context, obj *PrimitiveString) (int, error)
	}
	QueryResolver struct {
		InvalidIdentifier      func(ctx context.Context) (*invalid_packagename.InvalidIdentifier, error)
		Collision              func(ctx context.Context) (*introspection1.It, error)
		MapInput               func(ctx context.Context, input map[string]interface{}) (*bool, error)
		Recursive              func(ctx context.Context, input *RecursiveInputSlice) (*bool, error)
		NestedInputs           func(ctx context.Context, input [][]*OuterInput) (*bool, error)
		NestedOutputs          func(ctx context.Context) ([][]*OuterObject, error)
		Shapes                 func(ctx context.Context) ([]Shape, error)
		ModelMethods           func(ctx context.Context) (*ModelMethods, error)
		User                   func(ctx context.Context, id int) (*User, error)
		NullableArg            func(ctx context.Context, arg *int) (*string, error)
		DirectiveArg           func(ctx context.Context, arg string) (*string, error)
		DirectiveNullableArg   func(ctx context.Context, arg *int, arg2 *int) (*string, error)
		DirectiveInputNullable func(ctx context.Context, arg *InputDirectives) (*string, error)
		DirectiveInput         func(ctx context.Context, arg InputDirectives) (*string, error)
		DirectiveInputType     func(ctx context.Context, arg InnerInput) (*string, error)
		InputSlice             func(ctx context.Context, arg []string) (bool, error)
		ShapeUnion             func(ctx context.Context) (ShapeUnion, error)
		Autobind               func(ctx context.Context) (*Autobind, error)
		DeprecatedField        func(ctx context.Context) (string, error)
		Overlapping            func(ctx context.Context) (*OverlappingFields, error)
		MapStringInterface     func(ctx context.Context, in map[string]interface{}) (map[string]interface{}, error)
		ErrorBubble            func(ctx context.Context) (*Error, error)
		Errors                 func(ctx context.Context) (*Errors, error)
		Valid                  func(ctx context.Context) (string, error)
		Panics                 func(ctx context.Context) (*Panics, error)
		PrimitiveObject        func(ctx context.Context) ([]Primitive, error)
		PrimitiveStringObject  func(ctx context.Context) ([]PrimitiveString, error)
		DefaultScalar          func(ctx context.Context, arg string) (string, error)
		Slices                 func(ctx context.Context) (*Slices, error)
		ScalarSlice            func(ctx context.Context) ([]byte, error)
		Fallback               func(ctx context.Context, arg FallbackToStringEncoding) (FallbackToStringEncoding, error)
		OptionalUnion          func(ctx context.Context) (TestUnion, error)
		ValidType              func(ctx context.Context) (*ValidType, error)
		WrappedStruct          func(ctx context.Context) (*WrappedStruct, error)
		WrappedScalar          func(ctx context.Context) (WrappedScalar, error)
	}
	SubscriptionResolver struct {
		Updated     func(ctx context.Context) (<-chan string, error)
		InitPayload func(ctx context.Context) (<-chan string, error)
	}
	UserResolver struct {
		Friends func(ctx context.Context, obj *User) ([]*User, error)
	}
}

func (r *Stub) Errors() ErrorsResolver {
	return &stubErrors{r}
}
func (r *Stub) ForcedResolver() ForcedResolverResolver {
	return &stubForcedResolver{r}
}
func (r *Stub) ModelMethods() ModelMethodsResolver {
	return &stubModelMethods{r}
}
func (r *Stub) OverlappingFields() OverlappingFieldsResolver {
	return &stubOverlappingFields{r}
}
func (r *Stub) Panics() PanicsResolver {
	return &stubPanics{r}
}
func (r *Stub) Primitive() PrimitiveResolver {
	return &stubPrimitive{r}
}
func (r *Stub) PrimitiveString() PrimitiveStringResolver {
	return &stubPrimitiveString{r}
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

type stubErrors struct{ *Stub }

func (r *stubErrors) A(ctx context.Context, obj *Errors) (*Error, error) {
	return r.ErrorsResolver.A(ctx, obj)
}
func (r *stubErrors) B(ctx context.Context, obj *Errors) (*Error, error) {
	return r.ErrorsResolver.B(ctx, obj)
}
func (r *stubErrors) C(ctx context.Context, obj *Errors) (*Error, error) {
	return r.ErrorsResolver.C(ctx, obj)
}
func (r *stubErrors) D(ctx context.Context, obj *Errors) (*Error, error) {
	return r.ErrorsResolver.D(ctx, obj)
}
func (r *stubErrors) E(ctx context.Context, obj *Errors) (*Error, error) {
	return r.ErrorsResolver.E(ctx, obj)
}

type stubForcedResolver struct{ *Stub }

func (r *stubForcedResolver) Field(ctx context.Context, obj *ForcedResolver) (*Circle, error) {
	return r.ForcedResolverResolver.Field(ctx, obj)
}

type stubModelMethods struct{ *Stub }

func (r *stubModelMethods) ResolverField(ctx context.Context, obj *ModelMethods) (bool, error) {
	return r.ModelMethodsResolver.ResolverField(ctx, obj)
}

type stubOverlappingFields struct{ *Stub }

func (r *stubOverlappingFields) OldFoo(ctx context.Context, obj *OverlappingFields) (int, error) {
	return r.OverlappingFieldsResolver.OldFoo(ctx, obj)
}

type stubPanics struct{ *Stub }

func (r *stubPanics) FieldScalarMarshal(ctx context.Context, obj *Panics) ([]MarshalPanic, error) {
	return r.PanicsResolver.FieldScalarMarshal(ctx, obj)
}
func (r *stubPanics) ArgUnmarshal(ctx context.Context, obj *Panics, u []MarshalPanic) (bool, error) {
	return r.PanicsResolver.ArgUnmarshal(ctx, obj, u)
}

type stubPrimitive struct{ *Stub }

func (r *stubPrimitive) Value(ctx context.Context, obj *Primitive) (int, error) {
	return r.PrimitiveResolver.Value(ctx, obj)
}

type stubPrimitiveString struct{ *Stub }

func (r *stubPrimitiveString) Value(ctx context.Context, obj *PrimitiveString) (string, error) {
	return r.PrimitiveStringResolver.Value(ctx, obj)
}
func (r *stubPrimitiveString) Len(ctx context.Context, obj *PrimitiveString) (int, error) {
	return r.PrimitiveStringResolver.Len(ctx, obj)
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
func (r *stubQuery) ModelMethods(ctx context.Context) (*ModelMethods, error) {
	return r.QueryResolver.ModelMethods(ctx)
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
func (r *stubQuery) DirectiveInputType(ctx context.Context, arg InnerInput) (*string, error) {
	return r.QueryResolver.DirectiveInputType(ctx, arg)
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
func (r *stubQuery) DeprecatedField(ctx context.Context) (string, error) {
	return r.QueryResolver.DeprecatedField(ctx)
}
func (r *stubQuery) Overlapping(ctx context.Context) (*OverlappingFields, error) {
	return r.QueryResolver.Overlapping(ctx)
}
func (r *stubQuery) MapStringInterface(ctx context.Context, in map[string]interface{}) (map[string]interface{}, error) {
	return r.QueryResolver.MapStringInterface(ctx, in)
}
func (r *stubQuery) ErrorBubble(ctx context.Context) (*Error, error) {
	return r.QueryResolver.ErrorBubble(ctx)
}
func (r *stubQuery) Errors(ctx context.Context) (*Errors, error) {
	return r.QueryResolver.Errors(ctx)
}
func (r *stubQuery) Valid(ctx context.Context) (string, error) {
	return r.QueryResolver.Valid(ctx)
}
func (r *stubQuery) Panics(ctx context.Context) (*Panics, error) {
	return r.QueryResolver.Panics(ctx)
}
func (r *stubQuery) PrimitiveObject(ctx context.Context) ([]Primitive, error) {
	return r.QueryResolver.PrimitiveObject(ctx)
}
func (r *stubQuery) PrimitiveStringObject(ctx context.Context) ([]PrimitiveString, error) {
	return r.QueryResolver.PrimitiveStringObject(ctx)
}
func (r *stubQuery) DefaultScalar(ctx context.Context, arg string) (string, error) {
	return r.QueryResolver.DefaultScalar(ctx, arg)
}
func (r *stubQuery) Slices(ctx context.Context) (*Slices, error) {
	return r.QueryResolver.Slices(ctx)
}
func (r *stubQuery) ScalarSlice(ctx context.Context) ([]byte, error) {
	return r.QueryResolver.ScalarSlice(ctx)
}
func (r *stubQuery) Fallback(ctx context.Context, arg FallbackToStringEncoding) (FallbackToStringEncoding, error) {
	return r.QueryResolver.Fallback(ctx, arg)
}
func (r *stubQuery) OptionalUnion(ctx context.Context) (TestUnion, error) {
	return r.QueryResolver.OptionalUnion(ctx)
}
func (r *stubQuery) ValidType(ctx context.Context) (*ValidType, error) {
	return r.QueryResolver.ValidType(ctx)
}
func (r *stubQuery) WrappedStruct(ctx context.Context) (*WrappedStruct, error) {
	return r.QueryResolver.WrappedStruct(ctx)
}
func (r *stubQuery) WrappedScalar(ctx context.Context) (WrappedScalar, error) {
	return r.QueryResolver.WrappedScalar(ctx)
}

type stubSubscription struct{ *Stub }

func (r *stubSubscription) Updated(ctx context.Context) (<-chan string, error) {
	return r.SubscriptionResolver.Updated(ctx)
}
func (r *stubSubscription) InitPayload(ctx context.Context) (<-chan string, error) {
	return r.SubscriptionResolver.InitPayload(ctx)
}

type stubUser struct{ *Stub }

func (r *stubUser) Friends(ctx context.Context, obj *User) ([]*User, error) {
	return r.UserResolver.Friends(ctx, obj)
}
