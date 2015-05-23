package main

import (
	"net/http"
)

/*
type GreetingHandler struct{}

func (router GreetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Greetings!")
}
*/

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	rsp := make(chan *Rsp, 1)
	msgChnl <- NewMsg(GREETING, "", "", rsp)
	body := <- rsp
	w.Write(body.data)
}
