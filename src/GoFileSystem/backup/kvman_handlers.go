package main

import (
	"net/http"
	"os"
	"encoding/json"
)

func HandleCountKey(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Result int `json:"result"`
	} {
		len(table),
	}
	rsp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(rsp)
}

func HandleDump(w http.ResponseWriter, r *http.Request) {
	data := make([][2]string, len(table))
	counter := 0
	for k, v := range table {
		data[counter] = [2]string{k, v}
		counter += 1
	}
	rsp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(rsp)
}

func HandleShutdown(w http.ResponseWriter, r *http.Request) {
	if conn != nil {
		conn.Close()
	}
	os.Exit(0)
}
