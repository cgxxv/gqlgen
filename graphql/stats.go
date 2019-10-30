package graphql

import (
	"context"
	"fmt"
	"time"
)

type Stats struct {
	OperationStart time.Time
	Parsing        TraceTiming
	Validation     TraceTiming
}

type TraceTiming struct {
	Start time.Time
	End   time.Time
}

var ctxTraceStart key = "trace_start"

// StartOperationTrace captures the current time and stores it in context. This will eventually be added to request
// context but we want to grab it as soon as possible. For transports that can only handle a single graphql query
// per http requests you dont need to call this at all, the server will do it for you. For transports that handle
// multiple (eg batching, subscriptions) this should be called before decoding each request.
func StartOperationTrace(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxTraceStart, Now())
}

// GetStartTime should only be called by the handler package, it will be set into request context
// as Stats.Start
func GetStartTime(ctx context.Context) time.Time {
	t, ok := ctx.Value(ctxTraceStart).(time.Time)
	if !ok {
		panic(fmt.Sprintf("missing start time: %T", ctx.Value(ctxTraceStart)))
	}
	return t
}

// Now is time.Now, except in tests. Then it can be whatever you want it to be.
var Now = time.Now
