/**
 * @description chain
 * @author zhangbingbing@baidu.com
 * @date 2019/3/21
 */
package chain

type FilterChain interface {
	DoFilter(obj ...interface{}) bool
	init()
	addFilter(filter Filter)
	DoService(obj ...interface{}) bool
}

type FilterChainTemplate struct {
	FilterChain
}
