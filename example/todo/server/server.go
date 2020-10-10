package main

import (
	"context"
	"errors"
	"log"
	"runtime/debug"

	"github.com/99designs/gqlgen/example/todo"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/valyala/fasthttp"
)

func main() {
	srv := handler.NewDefaultServer(todo.NewExecutableSchema(todo.New()))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		// send this panic somewhere
		log.Print(err)
		debug.PrintStack()
		return errors.New("user message on panic")
	})

	gqlHandler := srv.Handler()
	playground := playground.Handler("Todo", "/query")

	h := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/query":
			gqlHandler(ctx)
		case "/":
			playground(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	log.Fatal(fasthttp.ListenAndServe(":8081", h))
}
