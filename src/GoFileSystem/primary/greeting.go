package main

import (
	"fmt"
	"net/http"
)

/*
type GreetingHandler struct{}

func (router GreetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Greetings!")
}
*/

func GreetingHandler(w http.ResponseWriter, r *http.Request) {

	rtn_str := Write("Hello, world")
	fmt.Println("Received Message:" + rtn_str)

	fmt.Fprint(w, "Greetings")
	msgChnl <- NewMsg(GREETING, "", "", &w)
}
