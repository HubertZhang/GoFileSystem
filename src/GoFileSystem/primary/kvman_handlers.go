package main

import (
	"fmt"
	"net/http"
)

func HandleCountKey(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handle Count Key")
}

func HandleDump(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handle dump.")
}

func HandleShutdown(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handle shutdown")
}
