package main

import (
	"fmt"
	"log"
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func TestMongod(t *testing.T) {
	// 建立一个会话
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
		t.Fail()

	}
	// 延迟最后关系会话
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// 指定数据库与表
	c := session.DB("test").C("people")
	// 播放数据
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	// 查找数据
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
	t.Log(result)
}
