package context

import (
	"context"
	"encoding/csv"
	"fmt"
	"math"
	"net/url"
	"os"
	"sync"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("go routine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了")
			return
		default:
			fmt.Println(name, "goroutine 监控中")
			time.Sleep(2 * time.Second)
		}
	}
}
func TestContext1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "监控1 ")
	go watch(ctx, "监控2 ")
	go watch(ctx, "监控3 ")
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(5 * time.Second)

}

var key string = "name"

func TestContext2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	varCtx := context.WithValue(ctx, key, "监控 ")
	go watch2(varCtx)
	varCtx1 := context.WithValue(ctx, key, "监控1")
	go watch2(varCtx1)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(5 * time.Second)

	first := 0
	for i := 0; i < 10; i++ {
		fmt.Println(first - (i + 1))
	}

}

func watch2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "监控退出，停止了")
			return
		default:
			fmt.Println(ctx.Value(key), "goroutine 监控中")
			time.Sleep(2 * time.Second)
		}
	}
}

func TestContext3(t *testing.T) {
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("go routine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	fmt.Println("before cancel")
	cancel()
	fmt.Println("after cancel")
	time.Sleep(5 * time.Second)

	fmt.Println("delete", len([]byte("delete")))
}

func TestContext4(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	var datas []map[string]interface{}
	data := map[string]interface{}{"name": "bill", "id": 0}
	for i := 0; i < 10; i++ {
		datas = append(datas, data)
	}

	for j := 0; j < len(datas); j++ {
		wg.Add(1)
		datas[j]["id"] = j
		jj := j
		m := map[string]interface{}{"id": j}
		d := datas[j]
		go func(sctx context.Context, swg *sync.WaitGroup) {
			defer swg.Done()
			select {
			case <-sctx.Done():
				fmt.Println("had canceled...")
				return
			default:
				if d["id"].(int) == 15 {
					cancel()
					return
				}
				time.Sleep(3 * time.Second)
				fmt.Println(d, m, jj)
			}

		}(ctx, &wg)
	}
	wg.Wait()
}

func DeepCopy(value map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for key, val := range value {
		newMap[key] = val
	}
	return newMap
}

func TestCvs(t *testing.T) {
	// data, _ := ioutil.ReadFile("./test.cvm")
	// reader := bytes.NewBuffer(data)
	file, _ := os.Open("./test.cvs")
	defer file.Close()
	cvsReader := csv.NewReader(file)
	str, _ := cvsReader.ReadAll()
	for _, line := range str {
		for _, item := range line {
			fmt.Print(item, " ")
		}
		fmt.Println()
	}
	fmt.Println(str)
}

func TestSome(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[len(arr)-1])
	for idx, val := range arr {
		fmt.Println(-1-idx, val)
	}
	i := 203
	f := float64(i)
	n := 10.1
	fmt.Println(203/10, 203%10, math.Ceil(f/n))

	u, err := url.Parse("http://g.hiphotos.baidu.com/boxapp_novel/wh%3D267%2C357/s94674974b7315c6043c063edbb86e720/1b4c510fd9f9d72a14c0f57cdf2a2834349bbb2f.jpg")
	fmt.Println(u.Host, u.Path, err)
}

func TestContext5(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for j := 0; j < 50; j++ {
		id := j
		wg.Add(1)
		go func(sctx context.Context, swg *sync.WaitGroup) {
			defer swg.Done()
			select {
			case <-sctx.Done():
				fmt.Printf("thread %d cancel\n", id)
			default:
				time.Sleep(2 * time.Second)
				fmt.Printf("thread %d done\n", id)
			}
		}(ctx, &wg)
	}
	// time.Sleep(1 * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestContext6(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	for j := 0; j < 20; j++ {
		time.Sleep(1 * time.Second)
		id := j
		wg.Add(1)
		go func(sctx context.Context, swg *sync.WaitGroup) {
			defer swg.Done()
			select {
			case <-sctx.Done():
				fmt.Printf("thread %d cancel\n", id)
			default:
				time.Sleep(1 * time.Second)
				fmt.Printf("thread %d done\n", id)
			}
		}(ctx, &wg)
	}
	wg.Wait()
}
