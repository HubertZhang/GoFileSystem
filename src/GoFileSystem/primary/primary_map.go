package main

import (
	"fmt"
)

var (
	table = make(map[string]string)
)

func waitForMsg() {
	for true {
		msg := *(<-msgChnl)
		switch msg.header {
		case GREETING:
			{
				// do nothing
			}
		case KV_INSERT:
			{
				table[msg.key] = msg.val
			}
		case KV_DELETE:
			{
				delete(table, msg.key)
			}
		case KV_GET:
			{
				fmt.Fprintln(*msg.w, table[msg.key])
			}
		case KV_UPDATE:
			{
				table[msg.key] = msg.val
			}
		case KVMAN_COUNTKEY:
			{
				fmt.Fprintln(*msg.w, len(table))
			}
		}
	}
}
