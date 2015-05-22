package main

import (
	// "log"
	"fmt"
	"net/http"
)

func HandleCountKey(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handle Count Key")
}
