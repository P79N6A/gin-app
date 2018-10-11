package context

import (
	"testing"
	"context"
	"fmt"
	"time"
	"encoding/csv"
	"os"
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
	//ctx, cancel := context.WithCancel(context.Background())
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

func TestCvs(t *testing.T) {
	//data, _ := ioutil.ReadFile("./test.cvm")
	//reader := bytes.NewBuffer(data)
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
