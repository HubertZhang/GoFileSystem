package main

import (
	// "log"
	"fmt"
	"net/http"
)

func HandleDump(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handle dump.")
}
