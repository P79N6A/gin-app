/**
 * @description singleton
 * @author zhangbingbing@baidu.com
 * @date 2019-01-31
 */
package singleton

import (
	"testing"
)

// var (
// 	myService *MyService
// )
//
// func init() {
// 	myService = NewMyService()
// }

const COUNT = 100000

func BenchmarkMyService_SayHi(b *testing.B) {
	for i := 0; i < COUNT; i++ {
		NewMyService().SayHi()
	}
}
func BenchmarkMyService_SayHello(b *testing.B) {
	for i := 0; i < COUNT; i++ {
		myService.SayHi()
	}
}
