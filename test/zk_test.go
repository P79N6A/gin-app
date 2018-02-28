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

	// 设置访问权限
	flags := int32(0)
	acl := zk.WorldACL(zk.PermAll)

	// 创建节点
	path, err := conn.Create("/01", []byte("data"), flags, acl)
	must(err)
	fmt.Printf("create: %+v\n", path)

	// 获取节点数据
	data, stat, err := conn.Get("/01")
	must(err)
	fmt.Printf("get: %+v %+v\n", string(data), stat)

	// 设置节点数据
	stat, err = conn.Set("/01", []byte("new data"), stat.Version)
	must(err)
	fmt.Printf("set: %+v\n", stat)

	// 删除节点
	err = conn.Delete("/01", -1)
	must(err)
	fmt.Printf("delete: ok\n")

	// 是否存在节点
	exists, stat, err := conn.Exists("/01")
	must(err)
	fmt.Printf("exists: %+v %+v\n", exists, stat)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
