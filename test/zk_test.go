package main

import (
	"fmt"
	"testing"

	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func connect() *zk.Conn {
	conn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second)
	must(err)
	return conn
}
func TestZK(t *testing.T) {
	// 新建客户端连接
	conn := connect()
	defer conn.Close()

	// 设置访问权限
	flags := int32(zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)

	var zpath string = "/02"
	// 创建节点
	path, err := conn.Create(zpath, []byte("hello,bill"), flags, acl)
	must(err)
	fmt.Printf("create: %+v\n", path)

	// 获取节点数据
	data, stat, err := conn.Get(zpath)
	must(err)
	fmt.Printf("get: %+v %+v\n", string(data), stat)

	// 设置节点数据
	stat, err = conn.Set(zpath, []byte("hello,bing"), stat.Version)
	must(err)
	fmt.Printf("set: %+v\n", stat)
	display(conn, zpath)

	// 删除节点
	err = conn.Delete(zpath, -1)
	must(err)
	fmt.Printf("delete: ok\n")

	// 是否存在节点
	exists, stat, err := conn.Exists(zpath)
	must(err)
	fmt.Printf("exists: %+v %+v\n", exists, stat)

	// 创建一个父节点
	zPath1 := "/dir"
	_, err = conn.Create(zPath1, []byte("parent_path"), 0, acl)
	must(err)
	// 创建其下的子节点
	for i := 1; i <= 3; i++ {
		key := fmt.Sprintf(zPath1+"/key%d", i)
		data := []byte(fmt.Sprintf("data-child-%d", i))
		path, err = conn.Create(key, data, flags, acl)
		must(err)
		fmt.Printf("%+v\n", path)
	}

	// 获取父节点数据
	data, _, err = conn.Get(zPath1)
	fmt.Printf("/dir: %s\n", string(data))
	// 获取父节点下的所有子节点
	children, _, err := conn.Children(zPath1)
	must(err)
	for _, name := range children {
		// 获取子节点数据
		data, _, err = conn.Get(zPath1 + "/" + name)
		must(err)
		fmt.Printf("/dir/%s: %s\n", name, string(data))
		err = conn.Delete(zPath1+"/"+name, 0)
	}
	err = conn.Delete(zPath1, 0)
	must(err)

	watch()

}
func watch() {
	// 监测一节点变化
	conn := connect()
	defer conn.Close()

	flags := int32(zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)
	snapshots, errors := mirror(conn, "/mirror")
	go func() {
		for {
			select {
			case snapshot := <-snapshots:
				fmt.Printf("%+v\n", snapshot)
			case err := <-errors:
				panic(err)
			}
		}
	}()

	conn2 := connect()
	time.Sleep(time.Second)

	_, err := conn2.Create("/mirror/one", []byte("one"), flags, acl)
	must(err)
	time.Sleep(time.Second)

	_, err = conn2.Create("/mirror/two", []byte("two"), flags, acl)
	must(err)
	time.Sleep(time.Second)

	err = conn2.Delete("/mirror/two", 0)
	must(err)
	time.Sleep(time.Second)

	_, err = conn2.Create("/mirror/three", []byte("three"), flags, acl)
	must(err)
	time.Sleep(time.Second)

	conn2.Close()
	time.Sleep(10 * time.Second)
}
func display(conn *zk.Conn, path string) {
	data, stat, err := conn.Get(path)
	must(err)
	fmt.Printf("get: %+v %+v\n", string(data), stat)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func mirror(conn *zk.Conn, path string) (chan []string, chan error) {
	snapshots := make(chan []string)
	errors := make(chan error)
	go func() {
		for {
			snapshot, _, events, err := conn.ChildrenW(path)
			if err != nil {
				errors <- err
				return
			}
			snapshots <- snapshot
			evt := <-events
			if evt.Err != nil {
				errors <- evt.Err
				return
			}
		}
	}()
	return snapshots, errors
}
