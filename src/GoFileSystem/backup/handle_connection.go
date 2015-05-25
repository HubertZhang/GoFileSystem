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
	conn_reader := bufio.NewReader(con)
	rsp, err := conn_reader.ReadByte()
	if err != nil {
		fmt.Println("HandleConnection::READSTRING::Error: " + err.Error())
	}
	if rsp == 1 {
		fmt.Println("Conn established.")
		conn = con
		passData()
		serve()
	} else if rsp == 2 {
		fmt.Println("Restore backup")
		conn = con
		restoreBackup(conn_reader)
		serve()
		return
	} else {
		fmt.Println("Reject")
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

	conn.Write([]byte{1})

	return op

}

func passData() {
	body, err := json.Marshal(table)
	if err != nil {
		fmt.Println("PASSDATA::Error: " + err.Error())
		return
	}
	var length uint64 = uint64(len(body))
	header := make([]byte, 8)
	binary.PutUvarint(header, length)

	conn.Write(header)
	conn.Write(body)
}

func restoreBackup(conn_reader *bufio.Reader) bool {
	header := make([]byte, 8)

	n, err := conn_reader.Read(header)
	if err != nil {
		fmt.Println("HandleConnection::READHEADER::Error: " + err.Error())
		return false
	}
	if n != 8 {
		fmt.Println("Header Error")
		return false
	}
	fmt.Println("Header read.")

	body_length, n := binary.Uvarint(header)
	if n < 0 {
		fmt.Println("Length too large")
		return false
	}
	if n == 0 {
		fmt.Println("Buffer is too small")
		return false
	}

	body := make([]byte, body_length)
	n, err = conn_reader.Read(body)
	if err != nil {
		fmt.Println("getMsg::READBODY::Error: " + err.Error())
		return false
	}
	if uint64(n) < body_length {
		fmt.Println("Body Error")
		return false
	}

	err = json.Unmarshal(body, &table)
	if err != nil {
		fmt.Println("RESTORE:Error: " + err.Error())
		return false
	}

	conn.Write([]byte{1})
	return true
}

func serve() {
	for true {
		msg := getMsg()
		if msg == nil {
			conn.Close()
			return
		}
		Perform(msg)
	}
}
