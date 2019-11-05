package graphql

import (
	"context"
)

type key string

const resolverCtx key = "resolver_context"

type ResolverContext struct {
	Parent *ResolverContext
	// The name of the type this field belongs to
	Object string
	// These are the args after processing, they can be mutated in middleware to change what the resolver will get.
	Args map[string]interface{}
	// The raw field
	Field CollectedField
	// The index of array in path.
	Index *int
	// The result object of resolver
	Result interface{}
	// IsMethod indicates if the resolver is a method
	IsMethod bool
}

func (r *ResolverContext) Path() []interface{} {
	var path []interface{}
	for it := r; it != nil; it = it.Parent {
		if it.Index != nil {
			path = append(path, *it.Index)
		} else if it.Field.Field != nil {
			path = append(path, it.Field.Alias)
		}
	}

	// because we are walking up the chain, all the elements are backwards, do an inplace flip.
	for i := len(path)/2 - 1; i >= 0; i-- {
		opp := len(path) - 1 - i
		path[i], path[opp] = path[opp], path[i]
	}

	return path
}

func GetResolverContext(ctx context.Context) *ResolverContext {
	if val, ok := ctx.Value(resolverCtx).(*ResolverContext); ok {
		return val
	}
	return nil
}

func WithResolverContext(ctx context.Context, rc *ResolverContext) context.Context {
	rc.Parent = GetResolverContext(ctx)
	return context.WithValue(ctx, resolverCtx, rc)
}

func equalPath(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
