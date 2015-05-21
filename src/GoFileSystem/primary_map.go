package main

import (
	"fmt"
	"msg_queue"
)

var (
	table = make(map[string]string)
)

func waitForMsg() {
	for true {
		msg := *(<-msg_queue.msgChnl)
		switch msg.header {
		case msg_queue.GREETING:
			{
				// do nothing
			}
		case msg_queue.KV_INSERT:
			{
				table[msg.key] = msg.val
			}
		case msg_queue.KV_DELETE:
			{
				delete(table, msg.key)
			}
		case msg_queue.KV_GET:
			{
				fmt.Fprintln(msg.w, table[msg.key])
			}
		case msg_queue.KV_UPDATE:
			{
				table[msg.key] = msg.val
			}
		case msg_queue.COUNTKEY:
			{
				fmt.Fprintln(msg.w, len(table))
			}
		}
	}
}
