package main

import (
	// "log"
	"fmt"
	"net/http"
)

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handle delete.")
}
