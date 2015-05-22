package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func Write(msg string) string {
	conn, err := net.Dial("tcp", "localhost:5000")
	defer conn.Close()

	if err != nil {
		fmt.Println("Write::DIAL::Error: " + err.Error())
		os.Exit(1)
	}
	fmt.Fprint(conn, msg + "\n")

	str, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Write::READSTRING::Error: " + err.Error())
		os.Exit(1)
	}
	return str
}
