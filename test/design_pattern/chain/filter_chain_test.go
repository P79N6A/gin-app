/**
 * @description chain
 * @author zhangbingbing@baidu.com
 * @date 2019/3/21
 */
package chain

import (
	"testing"
	"fmt"
)

func TestChain(t *testing.T) {
	DefaultFilterChainInstance.DoFilter()
}

func TestSome(t *testing.T) {
	name1 := "a"
	name2 := []string{"b", "c"}
	Names(name1, name2)
}

func Names(nm string, names ...interface{}) {
	fmt.Println(nm, names)
	fmt.Println(names[0],len(names))
}
