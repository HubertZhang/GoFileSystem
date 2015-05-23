package main

import (
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
	rsp    chan *Rsp
}

type Rsp struct {
	data []byte
	err  error
}

func NewMsg(hd int, k string, v string, rsp chan *Rsp) *Msg {
	ret := new(Msg)
	ret.header = hd
	ret.key = k
	ret.val = v
	ret.rsp = rsp

	return ret
}

func NewRsp(data []byte, err error) *Rsp {
	rsp := new(Rsp)
	rsp.data = data
	rsp.err = err
	return rsp
}

var (
	msgChnl = make(chan *Msg, 100)
)
