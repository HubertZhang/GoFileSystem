package main

import (
	// "log"
	"fmt"
	"net/http"
)

func HandleInsert(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handle Insert")
}
