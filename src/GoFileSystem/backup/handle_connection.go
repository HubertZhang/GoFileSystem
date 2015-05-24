package main

import (
	"fmt"
	"net"
	"bufio"
	"encoding/binary"
	"encoding/json"
)

type Op struct {
	OpCode int
	Key    string
	Value  string
}

var conn net.Conn = nil

func HandleConnection(con net.Conn) {
	fmt.Println("Handle Connection")
	rsp, err := bufio.NewReader(con).ReadByte()
	if err != nil {
		fmt.Println("HandleConnection::READSTRING::Error: " + err.Error())
	}
	if rsp == 1 {
		fmt.Println("Conn established.")

		con.Write([]byte{1})
		conn = con
		for true {
			msg := getMsg()
			if msg == nil {
				conn.Close()
				return
			}
		}
	} else {
		fmt.Println("Conn fail.")
		return
	}
}


func getMsg() *Op {
	header := make([]byte, 8)

	conn_reader := bufio.NewReader(conn)

	n, err := conn_reader.Read(header)
	if err != nil {
		fmt.Println("HandleConnection::READHEADER::Error: " + err.Error())
		return nil
	}
	if n != 8 {
		fmt.Println("Header Error")
		return nil
	}
	fmt.Println("Header read.")

	body_length, n := binary.Uvarint(header)
	if n < 0 {
		fmt.Println("Length too large")
		return nil
	}
	if n == 0 {
		fmt.Println("Buffer is too small")
		return nil
	}

	body := make([]byte, body_length)
	n, err = conn_reader.Read(body)
	if err != nil {
		fmt.Println("getMsg::READBODY::Error: " + err.Error())
		return nil
	}
	if uint64(n) < body_length {
		fmt.Println("Body Error")
		return nil
	}

	op := new(Op)
	err = json.Unmarshal(body, op)

	fmt.Println(op.OpCode)
	fmt.Println(op.Key)
	fmt.Println(op.Value)

	conn.Write([]byte{1})

	return op

}
