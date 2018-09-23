package main

import (
	"net"
	"fmt"
	"os"
	"encoding/json"
	"bufio"
	"encoding/binary"
)

func sendMessage(conn net.Conn) {
	//words := "hello, world!"
	//conn.Write([]byte(words))
	readerAndWriter:=bufio.NewReadWriter(bufio.NewReader(conn),bufio.NewWriter(conn))
	message:=map[string]interface{}{"name":"bill", "age":30, "email":"bill@126.com", "password":"1111"}
	data,_:=json.Marshal(message)
	bodyLength:=len(data)
	header:=make([]byte, 4)
	binary.BigEndian.PutUint32(header[0:4],uint32(bodyLength))
	readerAndWriter.Write(header)
	var msg map[string]interface{}
	json.Unmarshal(data,&msg)
	fmt.Println(msg)

	//conn.Write(data)
	readerAndWriter.Write(data)
	readerAndWriter.Flush()
	fmt.Println("send over")
	reply:=make([]byte, 1024)
	//conn.Read(reply)
	readerAndWriter.Read(reply)
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
