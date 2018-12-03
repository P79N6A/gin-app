/**
 * @description threadpool
 * @author zhangbingbing@baidu.com
 * @date 2018/11/30
 */
package threadpool

import (
	"context"
	"fmt"
	"sync"
)

type Worker interface {
	Run()
}

type WorkerPool struct {
	queueWorkerChannel chan Worker
	shutdownChannel    chan int
	maxWorker          int32
	wg                 sync.WaitGroup
	context            context.Context
	cancel             context.CancelFunc
}

func New(max int32) *WorkerPool {
	pool := &WorkerPool{
		maxWorker:          max,
		queueWorkerChannel: make(chan Worker, max),
		shutdownChannel:    make(chan int),
	}
	ctx, cancel := context.WithCancel(context.Background())
	pool.context = ctx
	pool.cancel = cancel
	return pool
}
func (pool *WorkerPool) Start() {
	go func() {
	done:
		for {
			select {
			case <-pool.shutdownChannel:
				// pool.wg.Done()
				// close(pool.shutdownChannel)
				// close(pool.queueWorkerChannel)
				fmt.Println("****pool is done****")
				break done
			case worker := <-pool.queueWorkerChannel:
				go func(ctx context.Context, wg *sync.WaitGroup) {
					defer pool.wg.Done()
					select {
					case <-ctx.Done():
						pool.Shutdown()
						fmt.Println("***cancel***")
						return
					default:
						worker.Run()
					}
				}(pool.context, &pool.wg)
			}
		}
	}()
}

func (pool *WorkerPool) Submit(worker Worker) {
	pool.queueWorkerChannel <- worker
	pool.wg.Add(1)
}
func (pool *WorkerPool) Shutdown() {
	pool.shutdownChannel <- 1
}

func (pool *WorkerPool) AwaitTermination() {
	pool.wg.Wait()
}

func (pool *WorkerPool) Cancel() {
	pool.cancel()
}
