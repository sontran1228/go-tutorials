package main

import (
	"fmt"
	"net/http"
)

// PlayerServer return player information
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
