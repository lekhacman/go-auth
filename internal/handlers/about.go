package handlers

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func NewHelloHandler(appName string) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "Welcome to %s!", appName)
	}
}
