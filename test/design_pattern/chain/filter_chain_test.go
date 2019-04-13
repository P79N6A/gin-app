/**
 * @description chain
 * @author zhangbingbing@baidu.com
 * @date 2019/3/21
 */
package chain

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestChain(t *testing.T) {
	chain:=NewDefaultFilterChain()
	chain.DoFilter()
}

func TestSome(t *testing.T) {
	name1 := "a"
	name2 := []string{"b", "c"}
	Names(name1, name2)
}

type AA struct {
	ID   int
	Name string
	Data interface{}
}

type DataType struct {
	ItemID   int
	ItemName string
	Ext      struct {
		ExtID   int
		ExtName string
	}
}

func Names(nm string, names ...interface{}) {
	fmt.Println(nm, names)
	fmt.Println(names[0], len(names))

}

type Module struct {
	ModuleID   int
	ModuleType int
	ShowAppIDs []int
	Children   []Module
}

func TestSome1(t *testing.T) {
	var resourceID int64 = 14332122
	log.Println(0 << 48)
	log.Println(0<<48 + resourceID)
	log.Println(1 << 48)
	log.Println(1<<48 + resourceID)
	log.Println(2 << 48)
	log.Println(2<<48 + resourceID)
	log.Println(0x0000000000000)
	log.Println(0x0000000000000 | resourceID)
	log.Println(0x1000000000000)
	log.Println(0x1000000000000 | resourceID)
	log.Println(0x2000000000000)
	log.Println(0x2000000000000 | resourceID)
	aa := AA{}
	log.Printf("%+v\n", aa)
	// aa.Data = DataType{}
	// log.Printf("%+v\n", aa)
	res, err := json.Marshal(aa)
	log.Println(string(res), err)
	m := Module{}
	log.Printf("%+v\n", m)

}
