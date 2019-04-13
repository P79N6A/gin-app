/**
 * @description chain
 * @author zhangbingbing@baidu.com
 * @date 2019/3/21
 */
package chain

import "log"

type BFilter struct {
	Filter
}

func (aFilter *BFilter) DoFilter(chain FilterChain, obj ...interface{}) bool {
	log.Println("before b filter do filter...")
	chain.DoFilter(obj)
	log.Println("after b filter do filter...")
	return true
}
