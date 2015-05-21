package main

import (
	"net/http"
)

const (
	GREETING       = 0
	KV_INSERT      = 1
	KV_DELETE      = 2
	KV_GET         = 3
	KV_UPDATE      = 4
	KVMAN_COUNTKEY = 5
	KVMAN_DUMP     = 6
	KVMAN_SHUTDOWN = 7
)

type Msg struct {
	header int
	key    string
	val    string
	w      *http.ReponseWriter
}

func newMsg(hd int, k, v string, ww *http.ResponseWriter) *Msg {
	ret := make(Msg)
	ret.header = hd
	ret.key = k
	ret.val = v
	ret.w = ww

	return &ret
}

var (
	msgChnl = make(chan *Msg, 100)
)
