package main

import (
	"fmt"
	"net/http"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handle get")
}
