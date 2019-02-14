/**
 * @description flag_test
 * $go run main.go -name "hello,bill" -age 100
 * $go run main.go -name="hello,bill" -age=100
 * $go run main.go --name="hello,bill" --age=100
 * $go run main.go --name "hello,bill" --age 100
 * Usage of /var/folders/kl/mk1n7t0179929lcb0l87hdg00000gn/T/go-build625526068/b001/exe/main:
 * 		-age int
 * 		      the value of age (default 10)
 * 		-name string
 * 		      the value of name (default "bill")
 * @author zhangbingbing@baidu.com
 * @date 2019-02-14
 */
package main

import (
	"flag"
	"log"
)

var (
	name string
	age  int
)

const (
	defaultName = "bill"
	defaultAge  = 10
	usageName   = "the value of name"
	usageAge    = "the value of age"
)

func init() {
	flag.StringVar(&name, "name", defaultName, usageName)
	flag.IntVar(&age, "age", defaultAge, usageAge)
	flag.Parse()
}
func main() {
	log.Println("name:", name)
	log.Println("age:", age)
}
