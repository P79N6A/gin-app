package main

import (
	"fmt"
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
)

func TestMem(t *testing.T) {
	// 连接服务器
	mc := memcache.New("127.0.0.1:11211")

	// 写数据 &memcache.Item{Key,Value,Flags,Expiration,casid}
	mc.Set(&memcache.Item{Key: "key_one", Value: []byte("hello")})
	mc.Set(&memcache.Item{Key: "key_two", Value: []byte("bill")})

	// 获取数据
	val, err := mc.Get("key_one")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	fmt.Printf("-- %s", val.Value)
	t.Log(val)
	// 获取多个key,返回map
	it, err := mc.GetMulti([]string{"key_one", "key_two"})
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}

	for k, v := range it {
		fmt.Printf("## %s => %s\n", k, v.Value)
		t.Log(v)
	}
}
