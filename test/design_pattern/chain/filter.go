/**
 * @description chain
 * @author zhangbingbing@baidu.com
 * @date 2019/3/21
 */
package chain

type Filter interface {
	DoFilter(chain FilterChain, obj ...interface{}) bool
}

type Request struct {
	ModuleID int
	ShowAppIDs []int
	Children []Request
}

type DataItem struct {
	ModuleID int
	ModuleName string
	ModuleType int
}
type Response map[int]interface{}



