package handlers

import (
	"fmt"
	"net/http"
)

func NewHelloHandler(appName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to %s!", appName)
	}
}
