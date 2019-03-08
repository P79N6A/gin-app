/**
 * @description adapter
 * @author zhangbingbing@baidu.com
 * @date 2019-03-06
 */
package adapter

import (
	"log"
	"sync"
)

type Render interface {
	render()
}

type RenderAdapter struct {
	Render
}

func (adapter *RenderAdapter) GetAdapterInfo() {
	log.Println("call GetAdapterInfo function")
}

type ARenderAdapter struct {
	RenderAdapter
	Context *RenderContext
	Style   string
}

func (a *ARenderAdapter) render() {
	log.Println("render a object")
	log.Println("in context ", a.Context.Name)
	a.Context.Result.Store("age", 30)
	a.GetAdapterInfo()

}

type BRenderAdapter struct {
	RenderAdapter
	Context *RenderContext
}

func (b *BRenderAdapter) render() {
	log.Println("render b object")
	log.Println("in context ", b.Context.Name)
	b.Context.Result.Store("name", "bill")
	b.GetAdapterInfo()
}

type RenderContext struct {
	Name      string
	Result    sync.Map
	renderMap map[string]Render
}

func NewRenderContext(name string) *RenderContext {
	return &RenderContext{
		Name:      name,
		renderMap: map[string]Render{},
	}
}
func (c *RenderContext) GetRender(style string) Render {
	if _, ok := c.renderMap[style]; !ok {
		switch style {
		case "a":
			c.renderMap[style] = &ARenderAdapter{Context: c}
		case "b":
			c.renderMap[style] = &BRenderAdapter{Context: c}
		}
	}

	return c.renderMap[style]
}
func (c *RenderContext) GetResult() {
	res := map[string]interface{}{}
	c.Result.Range(func(key, value interface{}) bool {
		res[key.(string)] = value
		return true
	})
	log.Println(res)
}
