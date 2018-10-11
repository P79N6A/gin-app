package tsync

import (
	"testing"
	"sync"
	"net/http"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _,url:=range urls{
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			res,err :=http.Get(url)
			t.Log(res,err)
		}(url)
	}

	wg.Wait()
}

func TestPool(t *testing.T) {
	p:=&sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	a:=p.Get().(int)
	p.Put(1)
	b:=p.Get().(int)
	t.Log(a, b)

}
