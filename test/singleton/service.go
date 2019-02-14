/**
 * @description singleton
 * @author zhangbingbing@baidu.com
 * @date {DATE}
 */
package singleton

import (
	"time"
)

type MyService struct {
}

var myService = &MyService{}

func NewMyService() *MyService {
	return &MyService{}
}
func (myService *MyService) SayHi() {
	// fmt.Println("say hi .....")
	time.Sleep(1 * time.Millisecond)
}

func (myService *MyService) SayHello() {
	// fmt.Println("say hello .....")
	time.Sleep(10 * time.Millisecond)
}
