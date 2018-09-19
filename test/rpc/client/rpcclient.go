package main

import (
	"net/rpc"
	"github.com/mgutz/logxi/v1"
	"fmt"
	"../server"
	"net"
	"net/http"
)

func main() {
	// 注册服务
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	// 调用服务
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dailing:", err)
	}

	args := &server.Args{7, 8}
	var reply int
	// 同步调用
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	quotient := new(server.Quotient)
	// 异步调用
	divCall := client.Go("Arith.Divide", args, &quotient, nil)
	replyCall := <-divCall.Done
	fmt.Println(replyCall, quotient.Rem, quotient.Quo)

}
