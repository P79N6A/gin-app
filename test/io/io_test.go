package io

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestAbc(logger *testing.T) {
	infile, err := os.Open("in.txt")
	if err != nil {
		logger.Error("open in.txt file error")
		return
	}
	defer infile.Close()
	values := make([]int, 0)

	buf := bufio.NewReader(infile)
	for {
		line, _, err1 := buf.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		value, err1 := strconv.Atoi(string(line))
		values = append(values, value)
	}
	logger.Log(values)
	outfile, err := os.Create("out.txt")
	if len(values) > 0 {
		if err != nil {
			logger.Error("create file out.txt error")
			return
		}
		defer outfile.Close()
		sort.Ints(values)
		for _, value := range values {
			str := strconv.Itoa(value)
			outfile.WriteString(str)
		}
	}
	logger.Log(values)
	logger.Log("hello,test")
	myLog := MyLog{log.New(outfile, "---", 1)}
	myLog.Println("hello")
	logger.Log(runtime.NumCPU())
}

type MyLog struct {
	*log.Logger
}

func TestStr(t *testing.T) {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func TestAI(t *testing.T) {
	reader := bufio.NewReader(os.Stdin)
	for {
		c, err := reader.ReadString('\n')
		if err == nil {
			c = strings.Replace(c, "吗", "", -1)
			c = strings.Replace(c, "?", "!", -1)
			fmt.Println(c)
		}
	}

}

type Log struct {
	App string
}

func (l *Log) Info(message ...string) {
	fmt.Println(l.App, message)
}

type MyService struct {
	*Log
}

func TestService(t *testing.T) {
	service := &MyService{&Log{"Test"}}
	service.Info("hello,bill",service.App)
	log.Println()
}
