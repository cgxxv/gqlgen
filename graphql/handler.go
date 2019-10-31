package graphql

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/vektah/gqlparser/gqlerror"
)

type (
	Resolver            func(ctx context.Context) (res interface{}, err error)
	Writer              func(Status, *Response)
	OperationMiddleware func(ctx context.Context, next OperationHandler, writer Writer)
	OperationHandler    func(ctx context.Context, writer Writer)
	ResponseHandler     func(ctx context.Context) *Response
	ResponseStream      func() *Response
	ResponseMiddleware  func(ctx context.Context, next ResponseHandler) *Response
	FieldMiddleware     func(ctx context.Context, next Resolver) (res interface{}, err error)

	RawParams struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
		Extensions    map[string]interface{} `json:"extensions"`
	}

	GraphExecutor interface {
		CreateRequestContext(ctx context.Context, params *RawParams) (*RequestContext, gqlerror.List)
		DispatchRequest(ctx context.Context, writer Writer)
	}

	// HandlerExtension interface is entirely optional, see the list of possible hook points below
	// Its important to understand the lifecycle of a graphql request and the terminology we use in gqlgen
	// before working with these
	//
	// +--- REQUEST   POST /graphql --------------------------------------------+
	// | +- OPERATION query OpName { viewer { name } } -----------------------+ |
	// | |  RESPONSE  { "data": { "viewer": { "name": "bob" } } }             | |
	// | +- OPERATION subscription OpName2 { chat { message } } --------------+ |
	// | |  RESPONSE  { "data": { "chat": { "message": "hello" } } }          | |
	// | |  RESPONSE  { "data": { "chat": { "message": "byee" } } }           | |
	// | +--------------------------------------------------------------------+ |
	// +------------------------------------------------------------------------+
	HandlerExtension interface{}

	// RequestParameterMutator is called before creating a request context. allows manipulating the raw query
	// on the way in.
	RequestParameterMutator interface {
		MutateRequestParameters(ctx context.Context, request *RawParams) *gqlerror.Error
	}

	// RequestContextMutator is called after creating the request context, but before executing the root resolver.
	RequestContextMutator interface {
		MutateRequestContext(ctx context.Context, rc *RequestContext) *gqlerror.Error
	}

	// OperationInterceptor is called for each incoming query, for basic requests the writer will be invoked once,
	// for subscriptions it will be invoked multiple times.
	OperationInterceptor interface {
		InterceptOperation(ctx context.Context, next OperationHandler, writer Writer)
	}

	// ResponseInterceptor is called around each graphql operation response. This can be called many times for a single
	// operation the case of subscriptions.
	ResponseInterceptor interface {
		InterceptResponse(ctx context.Context, next ResponseHandler) *Response
	}

	// FieldInterceptor called around each field
	FieldInterceptor interface {
		InterceptField(ctx context.Context, next Resolver) (res interface{}, err error)
	}

	// Transport provides support for different wire level encodings of graphql requests, eg Form, Get, Post, Websocket
	Transport interface {
		Supports(r *http.Request) bool
		Do(w http.ResponseWriter, r *http.Request, exec GraphExecutor)
	}
)

type Status int

const (
	StatusOk Status = iota
	StatusParseError
	StatusValidationError
	StatusResolverError
)

func (w Writer) Errorf(format string, args ...interface{}) {
	w(StatusResolverError, &Response{
		Errors: gqlerror.List{{Message: fmt.Sprintf(format, args...)}},
	})
}

func (w Writer) Error(msg string) {
	w(StatusResolverError, &Response{
		Errors: gqlerror.List{{Message: msg}},
	})
}

func (w Writer) GraphqlErr(err ...*gqlerror.Error) {
	w(StatusResolverError, &Response{
		Errors: err,
	})
}

func (p *RawParams) AddUpload(upload Upload, key, path string) *gqlerror.Error {
	if !strings.HasPrefix(path, "variables.") {
		return gqlerror.Errorf("invalid operations paths for key %s", key)
	}

	var ptr interface{} = p.Variables
	parts := strings.Split(path, ".")

	// skip the first part (variables) because we started there
	for i, p := range parts[1:] {
		last := i == len(parts)-2
		if ptr == nil {
			return gqlerror.Errorf("path is missing \"variables.\" prefix, key: %s, path: %s", key, path)
		}
		if index, parseNbrErr := strconv.Atoi(p); parseNbrErr == nil {
			if last {
				ptr.([]interface{})[index] = upload
			} else {
				ptr = ptr.([]interface{})[index]
			}
		} else {
			if last {
				ptr.(map[string]interface{})[p] = upload
			} else {
				ptr = ptr.(map[string]interface{})[p]
			}
		}
	}

	return nil
}
