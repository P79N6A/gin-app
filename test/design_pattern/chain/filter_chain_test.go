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
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestChain(t *testing.T) {
	chain := NewDefaultFilterChain()
	chain.DoFilter()
}

func TestSome(t *testing.T) {
	name1 := "a"
	name2 := []string{"b", "c"}
	Names(name1, name2)
	t1 := time.Now()
	var randSigleton = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		log.Println("---", randSigleton.Intn(10))
	}
	for i := 0; i < 10; i++ {
		log.Println("***", rand.Intn(10))
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		log.Println("###", rand.Intn(10))
	}
	elapsed := time.Since(t1)
	log.Println("take ", elapsed.Seconds()*1000)
	log.Println(time.Now().UnixNano())
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
	ModuleID   int   `json:"module_id"`
	ModuleType int   `json:"module_type"`
	ShowAppIDs []int `json:"show_app_ids"`
	Ext        []Ext `json:"ext"`
}

type Module1 struct {
	ModuleID   int                      `json:"module_id"`
	ModuleType int                      `json:"module_type"`
	ShowAppIDs []int                    `json:"show_app_ids"`
	Ext        []map[string]interface{} `json:"ext"`
}

type Ext struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Age  int    `json:"age"`
}

func TestThread(t *testing.T) {
	var sm sync.Map
	exts := map[int]*Ext{}
	for i := 0; i < 20; i++ {
		exts[i] = &Ext{Name: fmt.Sprintf("test %d", i), ID: i, Age: i + 10}
	}
	var wg sync.WaitGroup
	for _, ext := range exts {
		// extbak := ext
		wg.Add(1)
		go func(et *Ext) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			sm.Store(et.ID, et)
			fmt.Printf("%+v\n", et)
		}(ext)
	}
	wg.Wait()
	idx := 0
	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		idx++
		return true
	})
	delete(exts, 1)
	fmt.Println(idx)
	fmt.Println(exts)
}
func TestSome1(t *testing.T) {
	mm := Module{ModuleID: 1, ModuleType: 1, ShowAppIDs: []int{1, 2, 3}, Ext: []Ext{Ext{Name: "xxx", ID: 1, Age: 30}}}
	ss, err := json.Marshal(mm)
	log.Println(string(ss), err)
	var mm1 Module1
	err = json.Unmarshal(ss, &mm1)
	log.Printf("%+v\n", mm1)

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

	mp := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
	}
	for k, v := range mp {
		if k == 1 {
			delete(mp, k)
		} else {
			log.Println(k, v)
		}
	}
	log.Println(mp)
	var request, response interface{}
	log.Printf("%+v %+v\n", request, response)
	filterA(&request, &response)
	log.Printf("%+v %+v\n", request, response)
	filterB(&request, &response)
	log.Printf("%+v %+v\n", request, response)

	ms := []Module{{}, {}}
	log.Println(len(ms))
	Oper(&ms)
	log.Println(len(ms))

}

func Oper(list *[]Module) {
	*list = append(*list, Module{})
}

func filterA(request, response *interface{}) {
	*request = "aaaa"
	*response = "bbb"

}
func filterB(request, response *interface{}) {
	*request = Module{}
	*response = DataType{}

}
