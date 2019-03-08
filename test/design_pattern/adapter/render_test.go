/**
 * @description adapter
 * @author zhangbingbing@baidu.com
 * @date 2019-03-06
 */
package adapter

import (
	"log"
	"testing"
)

type A interface {
	Say()
	Listen()
}
type B interface {
	A
}
type C struct {
	A
}

func (c *C) Say() {
	log.Println("c say")
}

// func (c *C) Listen() {
// 	log.Println("c listen")
// }
type D struct {
}

func (d *D) Say() {
	log.Println("d say")
}
func (d *D) Listen() {

}

func TestInterface(t *testing.T) {
	var a A
	a = &C{}
	// a.Listen()
	a.Say()

	a = &D{}
	a.Say()
	// a.Listen()

}
func TestSome(t *testing.T) {
	context := NewRenderContext("test")
	context.GetRender("a").render()
	context.GetRender("b").render()
	context.GetResult()
}
