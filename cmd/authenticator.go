package main

import (
	"github.com/lekhacman/go-auth/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/about", handlers.NewHelloHandler("Go-Auth"))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
