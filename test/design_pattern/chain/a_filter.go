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

func (aFilter *AFilter) DoFilter(obj ...interface{}) bool {
	log.Println("a filter do filter...")
	return true
}
