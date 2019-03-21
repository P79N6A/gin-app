/**
 * @description chain
 * @author zhangbingbing@baidu.com
 * @date 2019/3/21
 */
package chain

type FilterChain interface {
	Filter
}

type DefaultFilterChain struct {
	FilterChain
}

var (
	filters                    []Filter
	DefaultFilterChainInstance = &DefaultFilterChain{}
)

func init() {
	addFilter(&AFilter{})
	addFilter(&BFilter{})
}
func (filterChain *DefaultFilterChain) DoFilter() bool {
	for _, filter := range filters {
		ok := filter.DoFilter()
		if !ok {
			break
		}
	}
	return true
}
func addFilter(filter Filter) {
	filters = append(filters, filter)
}
