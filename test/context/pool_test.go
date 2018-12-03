/**
 * @description context
 * @author zhangbingbing@baidu.com
 * @date 2018/12/1
 */
package context

import (
	"fmt"
	"log"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

type MyWorker struct {
}

func (w *MyWorker) Work() {
	time.Sleep(1 * time.Second)
	log.Println("*****worker done.....")
}

func TestPool(t *testing.T) {
	pool := NewGoroutinePool(10, 30)
	for i := 0; i < 100; i++ {
		worker := &MyWorker{}
		pool.Execute(worker)
	}
	// time.Sleep(50 * time.Millisecond)
	// pool.Shutdown()
	// fmt.Println("-------", all, done)
	// time.Sleep(100 * time.Second)
	pool.AwaitTermination()

}

func TestPool1(t *testing.T) {
	pool := NewGoroutinePool(10, 30)
	for i := 0; i < 100; i++ {
		id := i
		worker := &Worker{func() {
			time.Sleep(2 * time.Second)
			log.Println(id, "worker done...")
		}}
		pool.Execute(worker)
	}
	// time.Sleep(50 * time.Millisecond)
	// pool.Shutdown()
	// fmt.Println("-------", all, done)
	// time.Sleep(100 * time.Second)
	pool.AwaitTermination()

}

func TestChan(t *testing.T) {
	ch := make(chan MyWorker, 10)
	fmt.Println(ch, len(ch), ch == nil)
	close(ch)
	fmt.Println(ch, len(ch), ch == nil)

}

func TestAtom(t *testing.T) {
	var count int32 = 0
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			atomic.AddInt32(&count, 1)
		}()

	}
	for j := 0; j < 10; j++ {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("***", atomic.LoadInt32(&count))
		}()
	}

	time.Sleep(10 * time.Second)
}
