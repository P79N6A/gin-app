/**
 * @description chain
 * @author zhangbingbing@baidu.com
 * @date 2019/3/21
 */
package chain

import "testing"

func TestChain(t *testing.T) {
	DefaultFilterChainInstance.DoFilter()
}