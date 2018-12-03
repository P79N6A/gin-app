/**
 * @description context
 * @author zhangbingbing@baidu.com
 * @date 2018/12/1
 */
package context

import (
	"log"
	"sync"
	"sync/atomic"
)

type Workable interface {
	Work()
}

type Worker struct {
	action func()
}

func (worker *Worker) Work() {
	worker.action()
}

var once sync.Once
var defaultMaxWorkerCount = 100
var defaultMaxGoroutineCount = 10

type GoroutinePool struct {
	maxGoroutineCount  int
	maxWorkerCount     int
	workerQueue        chan Workable
	done               chan struct{}
	waitGroup          sync.WaitGroup
	currentWorkerCount int32
}

func NewGoroutinePool(maxGoroutineCount int, maxWorkerCount int) *GoroutinePool {
	if maxGoroutineCount == 0 {
		maxGoroutineCount = defaultMaxGoroutineCount
	}
	if maxWorkerCount == 0 {
		maxWorkerCount = defaultMaxWorkerCount
	}
	pool := &GoroutinePool{
		maxGoroutineCount:  maxGoroutineCount,
		workerQueue:        make(chan Workable, maxWorkerCount),
		done:               make(chan struct{}),
		currentWorkerCount: 0,
	}
	return pool
}
func (pool *GoroutinePool) watch() {
	for {
		select {
		case <-pool.done:
			pool.shutdown()
			return
		default:
			if atomic.LoadInt32(&pool.currentWorkerCount) == 0 {
				log.Println("all worker done....")
				pool.shutdown()
			}
		}
	}
}
func (pool *GoroutinePool) start() {
	pool.waitGroup.Add(pool.maxGoroutineCount)
	for i := 0; i < pool.maxGoroutineCount; i++ {
		go pool.doWork()
	}
	go pool.watch()
}
func (pool *GoroutinePool) Execute(worker Workable) {
	pool.workerQueue <- worker
	atomic.AddInt32(&pool.currentWorkerCount, 1)
	log.Println("add one worker")
	once.Do(pool.start)
}

func (pool *GoroutinePool) doWork() {
	for {
		select {
		case <-pool.done:
			log.Println("close goroutine....")
			pool.waitGroup.Done()
			pool.shutdown()
			return
		case worker := <-pool.workerQueue:
			if worker != nil {
				worker.Work()
				atomic.AddInt32(&pool.currentWorkerCount, -1)
			}
		}
	}
}

func (pool *GoroutinePool) Shutdown() {
	pool.shutdown()
}

func (pool *GoroutinePool) shutdown() {
	pool.done <- struct{}{}
	once.Do(func() {
		close(pool.workerQueue)
	})
}

func (pool *GoroutinePool) AwaitTermination() {
	pool.waitGroup.Wait()
}
