package handlers

import (
	"fmt"
	"net/http"
)

func (e Env) RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Mini!")
}
