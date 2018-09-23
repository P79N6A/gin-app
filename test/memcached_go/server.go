package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
	"encoding/binary"
	"io"
	"encoding/json"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	checkError(err)
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn.RemoteAddr().String(), " tcp connect success")
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	/*
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		fmt.Println(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
		conn.Write([]byte("had receive client message"))

	}
	*/
	readerAndWriter := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	header := make([]byte, 4)
	readerAndWriter.Read(header)
	fmt.Println(string(header))
	bodyLength := binary.BigEndian.Uint32(header)
	fmt.Println("body length:", bodyLength)
	bodyData:=make([]byte,bodyLength)
	io.ReadFull(readerAndWriter,bodyData)
	fmt.Println("body: ", string(bodyData))
	var jstr map[string]interface{}
	json.Unmarshal(bodyData,&jstr)
	fmt.Println(jstr)
}
func checkError(err error) {
	if err != nil {
		fmt.Printf("fatal error: %s", err.Error())
		os.Exit(1)
	}
}
