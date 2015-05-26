package main

import (
	// "log"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

func main() {
	err := init_config()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

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

	go waitForMsg()

	if !StartConn() {
		for !WaitReconnect() {
			time.Sleep(1 * time.Second)
		}
		go HeartBeat()
	}

	http.ListenAndServe(conf.Primary_ip+":"+conf.Http_port, r)
}
