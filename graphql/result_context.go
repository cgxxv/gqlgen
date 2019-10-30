package graphql

import (
	"context"
	"fmt"
	"sync"

	"github.com/vektah/gqlparser/gqlerror"
)

type resultContext struct {
	errors   gqlerror.List
	errorsMu sync.Mutex

	extensions   map[string]interface{}
	extensionsMu sync.Mutex
}

var resultCtx key = "result_context"

func getResultContext(ctx context.Context) *resultContext {
	val, _ := ctx.Value(resultCtx).(*resultContext)
	return val
}

func WithResultContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, resultCtx, &resultContext{})
}

// AddErrorf writes a formatted error to the client, first passing it through the error presenter.
func AddErrorf(ctx context.Context, format string, args ...interface{}) {
	c := getResultContext(ctx)

	c.errorsMu.Lock()
	defer c.errorsMu.Unlock()

	c.errors = append(c.errors, GetRequestContext(ctx).ErrorPresenter(ctx, fmt.Errorf(format, args...)))
}

// AddError sends an error to the client, first passing it through the error presenter.
func AddError(ctx context.Context, err error) {
	c := getResultContext(ctx)

	c.errorsMu.Lock()
	defer c.errorsMu.Unlock()

	c.errors = append(c.errors, GetRequestContext(ctx).ErrorPresenter(ctx, err))
}

// HasFieldError returns true if the given field has already errored
func HasFieldError(ctx context.Context, rctx *ResolverContext) bool {
	c := getResultContext(ctx)

	c.errorsMu.Lock()
	defer c.errorsMu.Unlock()
	path := rctx.Path()

	for _, err := range c.errors {
		if equalPath(err.Path, path) {
			return true
		}
	}
	return false
}

// GetFieldErrors returns a list of errors that occurred in the given field
func GetFieldErrors(ctx context.Context, rctx *ResolverContext) gqlerror.List {
	c := getResultContext(ctx)

	c.errorsMu.Lock()
	defer c.errorsMu.Unlock()
	path := rctx.Path()

	var errs gqlerror.List
	for _, err := range c.errors {
		if equalPath(err.Path, path) {
			errs = append(errs, err)
		}
	}
	return errs
}

func GetErrors(ctx context.Context) gqlerror.List {
	resCtx := getResultContext(ctx)
	resCtx.errorsMu.Lock()
	defer resCtx.errorsMu.Unlock()

	if len(resCtx.errors) == 0 {
		return nil
	}

	errs := resCtx.errors
	cpy := make(gqlerror.List, len(errs))
	for i := range errs {
		errCpy := *errs[i]
		cpy[i] = &errCpy
	}
	return cpy
}

// RegisterExtension allows you to add a new extension into the graphql response
func RegisterExtension(ctx context.Context, key string, value interface{}) {
	c := getResultContext(ctx)
	c.extensionsMu.Lock()
	defer c.extensionsMu.Unlock()

	if c.extensions == nil {
		c.extensions = make(map[string]interface{})
	}

	if _, ok := c.extensions[key]; ok {
		panic(fmt.Errorf("extension already registered for key %s", key))
	}

	c.extensions[key] = value
}

// GetExtensions returns any extensions registered in the current result context
func GetExtensions(ctx context.Context) map[string]interface{} {
	ext := getResultContext(ctx).extensions
	if ext == nil {
		return map[string]interface{}{}
	}

	return ext
}

func GetExtension(ctx context.Context, name string) interface{} {
	ext := getResultContext(ctx).extensions
	if ext == nil {
		return nil
	}

	return ext[name]
}
