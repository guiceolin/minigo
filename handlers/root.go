package handlers

import (
	"fmt"
	"net/http"
)

func RootHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Mini!")
	}
}
