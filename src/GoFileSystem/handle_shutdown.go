package main

import (
	// "log"
	"fmt"
	"net/http"
)

func HandleShutdown(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handle shutdown")
}
