/**
 * @description chain
 * @author zhangbingbing@baidu.com
 * @date 2019/3/21
 */
package chain

import "log"

type AFilter struct {
	Filter
}

func (aFilter *AFilter) DoFilter(chain FilterChain, obj ...interface{}) bool {
	log.Println("before a filter do filter...")
	chain.DoFilter(obj)
	log.Println("after a filter do filter...")
	return true
}
