package main

import (
	"encoding/json"
	"fmt"
)

var table = make(map[string]string)

func waitForMsg() {
	for true {
		msg := <-msgChnl
		switch msg.header {
		case GREETING:
			{
				data := struct {
					success bool
					value string
				} {
					true,
					"Hello, world",
				}
				rsp, err := json.Marshal(data)
				if err != nil {
					msg.rsp <- NewRsp(nil, err)
				}
				msg.rsp <- NewRsp(rsp, nil)
			}
		case KV_INSERT:
			{
				fmt.Println("Perform insert")
				perform_insert(msg)
			}
		case KV_DELETE:
			{
				delete(table, msg.key)
			}
		case KV_GET:
			{
				fmt.Println("Perform get.")
				perform_get(msg)
			}
		case KV_UPDATE:
			{
				table[msg.key] = msg.val
			}
		case KVMAN_COUNTKEY:
			{
			}
		}
	}
}


func perform_get(msg *Msg) {

	val, ok := table[msg.key]
	if ok {
		fmt.Println("Value: " + val)
		data := struct {
			Success bool `json:"success"`
			Value string `json:"value"`
		} {
			true,
			val,
		}
		rsp, err := json.Marshal(data)
		if err != nil {
			msg.rsp <- NewRsp(nil, err)
		}
		msg.rsp <- NewRsp(rsp, nil)
	} else {
		fmt.Println("Fail")
		data := struct {
			Success bool `json:"success"`
			Value string `json:"value"`
		} {
			false,
			"",
		}
		rsp, err := json.Marshal(data)
		if err != nil {
			msg.rsp <- NewRsp(nil, err)
			return
		}
		msg.rsp <- NewRsp(rsp, nil)
	}

}


func perform_insert(msg *Msg) {
	fmt.Println("Insert " + msg.key + " with " + msg.val)
	table[msg.key] = msg.val

	data := struct {
		Success bool `json:"success"`
	} {
		true,
	}
	rsp, err := json.Marshal(data)
	if err != nil {
		msg.rsp <- NewRsp(nil, err)
	}
	for k, v := range table {
		fmt.Println(k + " : " + v)
	}
	msg.rsp <- NewRsp(rsp, nil)
}

