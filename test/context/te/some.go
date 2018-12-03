/**
 * @description te
 * @author zhangbingbing@baidu.com
 * @date 2018/12/3
 */
package te

var var1 int = 0
var var2 int = 1

type User struct {
	ID int
}

func NewUser(ID int) *User {
	return &User{ID}
}
func (user *User) GetVar1() int {
	return var1
}
func (user *User) GetVar2() int {
	return var2
}

func (user *User) SetVar1(var11 int) {
	var1 = var11
}
func (user *User) SetVar2(var12 int) {
	var2 = var12
}
