package tsync

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			res, err := http.Get(url)
			t.Log(res, err)
		}(url)
	}

	wg.Wait()
}

func TestPool(t *testing.T) {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	ch := make(chan struct{})
	go func() {
		rand.Seed(time.Now().Unix())
		for {
			select {
			case <-ch:
				return
			default:
				p.Put(rand.Intn(100))
			}
		}
	}()
	time.Sleep(2 * time.Second)
	for {
		id := p.Get().(int)
		log.Println(id)
		if id == 100 {
			ch <- struct{}{}
			break
		}

		time.Sleep(1 * time.Second)
	}

}

type SafeMap struct {
	sync.RWMutex
	Map map[string]interface{}
}

func NewSafeMap() *SafeMap {
	sm := &SafeMap{}
	sm.Map = map[string]interface{}{}
	return sm
}
func (sm *SafeMap) Get(key string) interface{} {
	sm.RLock()
	defer sm.RUnlock()
	return sm.Map[key]
}
func (sm *SafeMap) Put(key string, value interface{}) {
	sm.Lock()
	defer sm.Unlock()
	sm.Map[key] = value
}
func TestMap1(t *testing.T) {
	testMap := NewSafeMap()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		id := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			key := strconv.Itoa(id)
			testMap.Put(key, true)
		}()
	}
	wg.Wait()
	fmt.Println(testMap)
	fmt.Println(len(testMap.Map))
}
func TestMap(t *testing.T) {
	testMap := map[string]interface{}{
	}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			key := strconv.Itoa(i)
			testMap[key] = true
		}()
	}
	wg.Wait()
	fmt.Println(testMap)
}

func TestMap2(t *testing.T) {
	var syncMap sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 500000; i++ {
		id := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			key := strconv.Itoa(id)
			syncMap.Store(key, true)
		}()
	}
	wg.Wait()
	count := 0
	var keys []string
	syncMap.Range(func(key, value interface{}) bool {
		// fmt.Println(key, value)
		count++
		keys = append(keys, key.(string))
		return true
	})
	fmt.Println(keys)
	fmt.Println(count)
}
func TestNum(t *testing.T) {
	var count int
	var syncCount int32
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			count = count + 1
			atomic.AddInt32(&syncCount, 1)
		}()
	}
	wg.Wait()
	fmt.Println(count)
	fmt.Println(syncCount)
}

type ConcurrentList struct {
	sync.RWMutex
	List []int
}

func NewConcurrentList() *ConcurrentList {
	return &ConcurrentList{}
}
func (list *ConcurrentList) Add(value int) {
	list.Lock()
	defer list.Unlock()
	list.List = append(list.List, value)
}
func (list *ConcurrentList) Get(index int) int {
	list.RLock()
	defer list.RUnlock()
	return list.List[index]
}
func (list *ConcurrentList) Size() int {
	list.RLock()
	defer list.RUnlock()
	return len(list.List)
}
func TestSlice(t *testing.T) {
	list := NewConcurrentList()
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		id := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			list.Add(id)
		}()
	}
	wg.Wait()
	fmt.Println(list)
	fmt.Println(list.Size())
}

func TestNum1(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			fmt.Println("****", id)
		}(i)
	}
	wg.Wait()
}

type TestStruct struct {
	First int
}

var (
	mapDemo    = make(map[int]int, 1)
	structDemo = TestStruct{}
)

func mapIncr() {
	mapDemo[1]++
}
func structIncr() {
	structDemo.First++
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapIncr()
	}
}

func BenchmarkStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		structIncr()
	}
}
