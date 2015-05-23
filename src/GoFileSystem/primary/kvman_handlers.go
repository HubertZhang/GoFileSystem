package main

import (
	"net/http"
	"os"
)

func HandleCountKey(w http.ResponseWriter, r *http.Request) {
	rsp := make(chan *Rsp, 1)
	msg := NewMsg(KVMAN_COUNTKEY, "", "", rsp)
	msgChnl <- msg
	body := <- rsp
	if body.err != nil {
		http.Error(w, body.err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body.data)
}

func HandleDump(w http.ResponseWriter, r *http.Request) {
	rsp := make(chan *Rsp, 1)
	msg := NewMsg(KVMAN_DUMP, "", "", rsp)
	msgChnl <- msg
	body := <- rsp
	if body.err != nil {
		http.Error(w, body.err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body.data)
}

func HandleShutdown(w http.ResponseWriter, r *http.Request) {
	rsp := make(chan *Rsp, 1)
	msg := NewMsg(KVMAN_SHUTDOWN, "", "", rsp)
	msgChnl <- msg
	body := <- rsp
	if body.err != nil {
		http.Error(w, body.err.Error(), http.StatusInternalServerError)
	}

	os.Exit(0)
}
