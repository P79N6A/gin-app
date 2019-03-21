/**
 * @description observer
 * @author zhangbingbing@baidu.com
 * @date 2019/3/13
 */
package observer

import (
	"testing"
	"reflect"
	"log"
	"gin-app/test/design_pattern/observer/event"
	"time"
	"encoding/json"
)

type TestEvent struct {
	Type int
}

func TestListener(t *testing.T) {
	evt := TestEvent{1}
	tt := reflect.TypeOf(evt)
	log.Println(tt, tt.Name() == "TestEvent")
	_, ok := interface{}(evt).(TestEvent)
	log.Println(ok)
}

func TestListener1(t *testing.T) {
	context := NewDefaultApplicationContext()
	evt := &event.AEvent{}
	context.Publish(evt)
	ent := &event.BEvent{}
	context.Publish(ent)
	time.Sleep(1 * time.Second)
}

func TestListener2(t *testing.T) {
	evt := &event.AEvent{}
	DefaultContext.Publish(evt)
	ent := &event.BEvent{}
	DefaultContext.Publish(ent)
	time.Sleep(1 * time.Second)
}

type Shape struct {
	Width  int
	Height int
}

type AShape struct {
	*Shape
	Name string
}

type BShape struct {
	Width  int
	Height int
	Name   string
}

func TestSome(t *testing.T) {
	a := &BShape{}
	a.Name = "aa"
	a.Width = 10
	a.Height = 20
	bt, err := json.Marshal(a)
	log.Println(string(bt), err)

	var bb AShape
	err = json.Unmarshal(bt, &bb)
	log.Println(bb)
}
