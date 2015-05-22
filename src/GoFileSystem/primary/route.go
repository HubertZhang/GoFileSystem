package main

import (
	// "log"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// var greeting_handler GreetingHandler
	// http.Handle("/greeting", greeting_handler)
	r := mux.NewRouter()
	r.HandleFunc("/greeting", GreetingHandler)
	r.HandleFunc("/kv/insert", HandleInsert).Methods("POST")
	r.HandleFunc("/kv/delete", HandleDelete).Methods("POST")
	r.HandleFunc("/kv/get", HandleGet).Methods("GET")
	r.HandleFunc("/kv/update", HandleUpdate).Methods("POST")
	r.HandleFunc("/kvman/countkey", HandleCountKey).Methods("GET")
	r.HandleFunc("/kvman/dump", HandleDump).Methods("GET")
	r.HandleFunc("/kvman/shutdown", HandleShutdown).Methods("GET")

	go primary_map.waitForMsg()
	http.ListenAndServe("localhost:4000", r)
}
