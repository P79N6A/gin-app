package io

import (
	"testing"
	"os"
	"bufio"
	"io"
	"strconv"
	"log"
	"sort"
	"runtime"
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
