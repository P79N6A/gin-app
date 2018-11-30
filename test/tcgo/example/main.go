package main

/*
//-I指定头文件目录
#cgo CFLAGS: -I./foo
//-L 指定引用库的目录,-l指定库名称
#cgo LDFLAGS: -L./foo -lfoo
#include "foo.h"
 */
import "C"
import "fmt"

func main() {
    a:=C.int(1)
    b:=C.int(2)
    value:=C.add(a, b)
    fmt.Printf("%v\n", value)

}
