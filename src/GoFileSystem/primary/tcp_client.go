package main

import (
	"fmt"
	"net"
	"bufio"
	"encoding/binary"
)

var conn net.Conn = nil

type Op struct {
	OpCode int
	Key    string
	Value  string
}

func NewOp(code int, key string, value string) *Op {
	rtn := new(Op)
	rtn.OpCode = code
	rtn.Key = key
	rtn.Value = value
	return rtn
}

func StartConn() bool {
	fmt.Println("StartConn.")
	var err error = nil
	conn, err = net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Write::DIAL::Error: " + err.Error())
		return false
	}

	conn.Write([]byte{1})

	rsp, err := bufio.NewReader(conn).ReadByte()
	if err != nil {
		fmt.Println("Write::READBYTE::Error: " + err.Error())
		return false
	}
	if rsp == 1 {
		fmt.Println("Started.")
		return true
	} else {
		fmt.Println("Fail")
		return false
	}
}

func Write(msg []byte) bool {
	var length uint64 = uint64(len(msg))
	header := make([]byte, 8)
	binary.PutUvarint(header, length)

	conn.Write(header)
	conn.Write(msg)

	rsp, err := bufio.NewReader(conn).ReadByte()
	if err != nil {
		fmt.Println("Write::READSTRING::Error: " + err.Error())
		return false
	}
	if rsp == 1 {
		return true
	} else {
		return false
	}
}
