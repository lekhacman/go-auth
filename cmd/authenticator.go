package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/lekhacman/go-auth/internal/handlers"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	router := fasthttprouter.New()
	router.GET("/", handlers.NewHelloHandler("Go Auth"))

	log.Fatal(fasthttp.ListenAndServe(":8000", router.Handler))
}
