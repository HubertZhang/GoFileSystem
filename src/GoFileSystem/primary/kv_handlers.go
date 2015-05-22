package main

import (
	"fmt"
	"net/http"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handle get")
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handle delete.")
}

func HandleInsert(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handle Insert")
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handle update")
}
