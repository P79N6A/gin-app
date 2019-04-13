/**
 * @description chain
 * @author zhangbingbing@baidu.com
 * @date 2019/3/21
 */
package chain

import "log"

type FilterChain interface {
	DoFilter(obj ...interface{}) bool
	init()
	addFilter(filter Filter)
}

type DefaultFilterChain struct {
	FilterChain
	pos     int
	count   int
	filters []Filter
}

func NewDefaultFilterChain() *DefaultFilterChain {
	chain := &DefaultFilterChain{}
	chain.init()
	return chain
}
func (chain *DefaultFilterChain) init() {
	chain.addFilter(&AFilter{})
	chain.addFilter(&BFilter{})
}
func (chain *DefaultFilterChain) DoFilter(obj ...interface{}) bool {
	if chain.pos < chain.count {
		filter := chain.filters[chain.pos]
		chain.pos++
		filter.DoFilter(chain, nil)
		return true
	}
	log.Println("all filter done...")
	return true
}
func (chain *DefaultFilterChain) addFilter(filter Filter) {
	chain.count++
	chain.filters = append(chain.filters, filter)
}
