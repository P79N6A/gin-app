/**
 * @description threadpool
 * @author zhangbingbing@baidu.com
 * @date 2018/11/30
 */
package threadpool

import (
	"fmt"
	"testing"
	"time"
)

type MyWorker struct {
}

func (myWorker *MyWorker) Run() {
	time.Sleep(5 * time.Second)
	fmt.Println("hello,bill")
}
func TestWorkerPool_Start(t *testing.T) {
	pool := New(10)
	for i := 0; i < 10; i++ {
		worker := &MyWorker{}
		pool.Submit(worker)
	}
	pool.Start()
	pool.Cancel()
	pool.AwaitTermination()
	fmt.Println("done......")
	pool.Shutdown()
}
