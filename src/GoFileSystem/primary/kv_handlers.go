package main

import (
	"net/http"
	"encoding/json"
)

type Req struct {
	Key   string
	Value string
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	k := r.URL.Query().Get("key")
	rsp := make(chan *Rsp, 1)
	msg := NewMsg(KV_GET, k, "", rsp)
	msgChnl <- msg
	body := <- rsp
	if body.err != nil {
		// http.Error(w, body.err.Error(), http.StatusInternalServerError)
		body.data = returnError()
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body.data)
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	var req Req
	req.Key = r.PostFormValue("key")
	// req.Value = r.PostFormValue("value")
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&req)
	// if err != nil {
	// 	// http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(returnError())
	// 	return
	// }
	k := req.Key

	// k := r.URL.Query().Get("key")
	rsp := make(chan *Rsp, 1)
	msg := NewMsg(KV_DELETE, k, "", rsp)
	msgChnl <- msg
	body := <- rsp
	if body.err != nil {
		// http.Error(w, body.err.Error(), http.StatusInternalServerError)
		body.data = returnError()
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body.data)
}

func HandleInsert(w http.ResponseWriter, r *http.Request) {
	var req Req
	req.Key = r.PostFormValue("key")
	req.Value = r.PostFormValue("value")
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&req)
	// if err != nil {
	// 	// http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(returnError())
	// 	return
	// }
	k := req.Key
	v := req.Value
	// k := r.URL.Query().Get("key")
	// v := r.URL.Query().Get("value")
	rsp := make(chan *Rsp, 1)
	msg := NewMsg(KV_INSERT, k, v, rsp)
	msgChnl <- msg
	body := <- rsp
	if body.err != nil {
		// http.Error(w, body.err.Error(), http.StatusInternalServerError)
		body.data = returnError()
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body.data)
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var req Req
	req.Key = r.PostFormValue("key")
	req.Value = r.PostFormValue("value")
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&req)
	// if err != nil {
	// 	// http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(returnError())
	// 	return
	// }
	k := req.Key
	v := req.Value

	// k := r.URL.Query().Get("key")
	// v := r.URL.Query().Get("value")
	rsp := make(chan *Rsp, 1)
	msg := NewMsg(KV_UPDATE, k, v, rsp)
	msgChnl <- msg
	body := <- rsp
	if body.err != nil {
		// http.Error(w, body.err.Error(), http.StatusInternalServerError)
		body.data = returnError()
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body.data)
}

func returnError() []byte {
	data := struct {
		Success bool `json:"success"`
	} {
		false,
	}

	rsp, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return rsp
}
