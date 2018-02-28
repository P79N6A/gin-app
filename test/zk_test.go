package main

import (
	"fmt"
	"testing"

	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func TestZK(t *testing.T) {
	// 新建客户端连接
	conn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second)
	must(err)
	defer conn.Close()

	flags := int32(0)
	acl := zk.WorldACL(zk.PermAll)
	path, err := conn.Create("/01", []byte("data"), flags, acl)
	must(err)
	fmt.Printf("create: %+v\n", path)

	data, stat, err := conn.Get("/01")
	must(err)
	fmt.Printf("get: %+v %+v\n", string(data), stat)

	stat, err = conn.Set("/01", []byte("new data"), stat.Version)
	must(err)
	fmt.Printf("set: %+v\n", stat)

	err = conn.Delete("/01", -1)
	must(err)
	fmt.Printf("delete: ok\n")

	exists, stat, err := conn.Exists("/01")
	must(err)
	fmt.Printf("exists: %+v %+v\n", exists, stat)
}
func must(err error) {
	if err != nil {
		panic(err)
	}
}
