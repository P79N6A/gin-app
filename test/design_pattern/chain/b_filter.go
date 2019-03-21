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

func (aFilter *BFilter) DoFilter() bool{
	log.Println("b filter do filter...")
	return true
}
