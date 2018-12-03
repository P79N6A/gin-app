/**
 * @description context
 * @author zhangbingbing@baidu.com
 * @date 2018/12/3
 */
package context

import (
	"fmt"
	"gin-app/test/context/te"
	"testing"
)

func TestSome1(t *testing.T) {
	u1:=te.NewUser(1)
	u2:=te.NewUser(2)
	fmt.Println(u1.GetVar1(),u2.GetVar1(),u1.GetVar2(),u1.GetVar2())
	u1.SetVar1(3)
	u2.SetVar2(4)
	fmt.Println(u1.GetVar1(),u2.GetVar1(),u1.GetVar2(),u1.GetVar2())
}