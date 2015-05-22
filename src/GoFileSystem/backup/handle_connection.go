package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func HandleConnection(conn net.Conn) {
	fmt.Println("Handle Connection")

	defer conn.Close()

	_, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("HandleConnection::READESTRING::Error: " + err.Error())
		os.Exit(1)
	}

	conn.Write([]byte("Message Received\n"))
}
