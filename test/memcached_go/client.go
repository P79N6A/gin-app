package main

import (
	"net"
	"fmt"
	"os"
	"encoding/json"
)

func sendMessage(conn net.Conn) {
	//words := "hello, world!"
	//conn.Write([]byte(words))
	message:=map[string]interface{}{"name":"bill", "age":30, "email":"bill@126.com", "password":"1111"}
	data,_:=json.Marshal(message)
	var msg map[string]interface{}
	json.Unmarshal(data,&msg)
	fmt.Println(msg)

	conn.Write(data)
	fmt.Println("send over")
	reply:=make([]byte, 1024)
	conn.Read(reply)
	fmt.Println(string(reply))
}

func main() {

	server := "127.0.0.1:8000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Printf("fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Printf("fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	sendMessage(conn)
}
