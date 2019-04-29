/**
 * @description chain
 * @author zhangbingbing@baidu.com
 * @date 2019/3/21
 */
package chain

import "log"

type DefaultFilterChain struct {
	FilterChainTemplate
	pos     int
	count   int
	filters []Filter
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
	chain.DoService(obj)
	return true
}
func (chain *DefaultFilterChain) addFilter(filter Filter) {
	chain.count++
	chain.filters = append(chain.filters, filter)
}

func NewDefaultFilterChain() *DefaultFilterChain {
	chain := &DefaultFilterChain{}
	chain.FilterChain = chain
	chain.init()
	return chain
}

func (chain *DefaultFilterChain) DoService(obj ...interface{}) bool {
	log.Println("chain do service ...")
	return true
}
