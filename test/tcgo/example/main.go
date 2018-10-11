package main

/*
#cgo CFLAGS: -I./foo
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
