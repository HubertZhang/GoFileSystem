package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
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
	conn, err = net.Dial("tcp", "http://"+conf.backup_ip+":5000")
	if err != nil {
		fmt.Println("Write::DIAL::Error: " + err.Error())
		return false
	}

	conn.Write([]byte{1})

	header := make([]byte, 8)

	reader := bufio.NewReader(conn)

	n, err := reader.Read(header)
	if err != nil {
		fmt.Println("Write::READBYTE::Error: " + err.Error())
		return false
	}
	if n != 8 {
		fmt.Println("STARUP::Error: Header Error")
		return false
	}

	body_length, n := binary.Uvarint(header)
	if n < 0 {
		fmt.Println("STARTUP::Error: Length too large")
		return false
	}
	if n == 0 {
		fmt.Println("STARTUP::Error: Buffer is too small")
		return false
	}

	body := make([]byte, body_length)
	var readed uint64 = 0
	for readed != body_length {
		n, err = reader.Read(body[readed:])
		if err != nil {
			fmt.Println("getMsg::READBODY::Error: " + err.Error())
			return false
		}
		fmt.Println("Get package of length: " + strconv.Itoa(n))
		readed += uint64(n)
	}

	err = json.Unmarshal(body, &table)
	if err != nil {
		fmt.Println("STARTUP::Error: " + err.Error())
		return false
	}

	go HeartBeat()
	return true
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

func WaitReconnect() bool {
	fmt.Println("StartConn.")
	var err error = nil
	conn, err = net.Dial("tcp", "http://"+conf.backup_ip+":5000")
	if err != nil {
		fmt.Println("Write::DIAL::Error: " + err.Error())
		return false
	}

	conn.Write([]byte{2})

	data, err := json.Marshal(table)
	if err != nil {
		fmt.Println("WAITRECONNECT::Error: " + err.Error())
		return false
	}
	var length uint64 = uint64(len(data))
	header := make([]byte, 8)
	binary.PutUvarint(header, length)
	conn.Write(header)
	conn.Write(data)

	rsp, err := bufio.NewReader(conn).ReadByte()
	if err != nil {
		fmt.Println("WAITRECONNECT::Error: " + err.Error())
		return false
	}
	if rsp == 1 {
		return true
	} else {
		fmt.Println("Wrong Reply")
		return false
	}
}

func HeartBeat() {
	time.Sleep(1 * time.Second)
	if len(msgChnl) == 0 {
		msgChnl <- NewMsg(HEARTBEAT, "", "", nil)
	}
	go HeartBeat()
}
