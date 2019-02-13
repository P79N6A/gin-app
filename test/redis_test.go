package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestRedis(t *testing.T) {
	// 新建客户端连接
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer client.Close()
	// 连接服务器测试
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	t.Log(pong, err)

	// 添加key-value
	err = client.Set("key1", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取存在的key
	val, err := client.Get("key1").Result()
	fmt.Println(val, err)
	t.Log(val, err)
	// 获取不存在的key
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
		t.Fail()
	} else {
		fmt.Println("key2", val2)
		t.Log(val2)
	}
}

func TestRedis1(t *testing.T) {
	// 新建客户端连接
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer client.Close()
	rand.Seed(time.Now().Unix())
	uid := 10000
	var mems []redis.Z
	for i := 0; i < 100000000; i++ {
		mems = append(mems, redis.Z{float64(rand.Intn(1000000)), uid + i})
		if i > 0 && i%1000 == 0 {
			err := client.ZAdd("sorted_set", mems...).Err()
			if err != nil {
				t.Error(err)
			}
			mems = []redis.Z{}
			time.Sleep(100 * time.Millisecond)
		}
	}
	res, err := client.ZCard("sorted_set").Result()
	t.Log(res, err)

}
func TestSome(t *testing.T) {
	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	t.Log(arr)
	arr = []int{}
	t.Log(arr)

	var aa = []byte{123, 34, 114, 101, 115, 112, 95, 115, 116, 114, 34, 58, 34, 123, 39, 101, 114, 114, 110, 111, 39, 58, 32, 45, 49, 44, 32, 39, 101, 114, 114, 109, 115, 103, 39, 58, 32, 39, 98, 97, 100, 32, 114, 101, 113, 117, 101, 115, 116, 39, 125, 92, 110, 34, 125,}
	fmt.Println(string(aa))
}
