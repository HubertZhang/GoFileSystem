package main

import (
	// "log"
	"fmt"
	"net/http"
)

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handle update")
}
