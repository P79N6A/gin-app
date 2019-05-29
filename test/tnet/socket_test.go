package tnet

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"
)

func TestNet(t *testing.T) {
	t.Log("aaa")
	resp, err := http.Get("http://www.baidu.com")
	t.Log(resp, err, resp.Status, resp.Body)
	resp, err = http.PostForm("http://www.baidu.com", url.Values{"name": {"value"}})
	t.Log(resp, err)
}

func TestNet1(t *testing.T) {
	client := &http.Client{CheckRedirect: nil}
	res, err := client.Get("http://www.baidu.com")
	req, err := http.NewRequest("get", "http://www.baidu.com", nil)
	req.Header.Add("If-None-Match", "sss")
	res, err = client.Do(req)
	t.Log(res, err)
	ck, _ := req.Cookie("name")
	t.Log(ck.String())
	t.Log(http.CanonicalHeaderKey("accept-encoding"), http.CanonicalHeaderKey("if-none-match"))
}

func TestUrl(t *testing.T) {
	urlStr := "http://yq01-zhangbingbing.epc.baidu.com:8003/searchbox?action=novel&type=selected"
	res := url.QueryEscape(urlStr)
	t.Log(res)
	res, _ = url.QueryUnescape(res)
	t.Log(res)

	u, _ := url.Parse(urlStr)
	t.Log(u.User, u.Scheme, u.Host, u.Query(), u.RawQuery, u.Path)

	v := url.Values{}
	v.Set("name", "bill")
	v.Set("age", "30")
	v.Set("email", "bill@126.com")
	t.Log(v.Encode())

	v, _ = url.ParseQuery(u.RawQuery)
	t.Log(v)

}

// 服务端没启动服务或不可用，则报connection refused
func TestTcp(t *testing.T) {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		log.Println("dial error:", err)
		// err:dial tcp :8899: getsockopt: connection refused
		return
	}
	defer conn.Close()
	log.Println("dial ok")

}
func TestTCPServer(t *testing.T) {
	l, err := net.Listen("tcp", ":8899")
	if err != nil {
		log.Println("error listen:", err)
		return
	}
	defer l.Close()
	log.Println("listen ok")

	var i int
	for {
		time.Sleep(10 * time.Second)
		if _, err := l.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		}
		i++
		log.Printf("%d: accept a new connection\n", i)
	}
}
func TestBacklog(t *testing.T) {
	// server

	// client
	for i := 1; i < 1000; i++ {

		conn, err := net.Dial("tcp", "10.64.49.20:8888")
		if err != nil {
			log.Printf("%d: dial error:%s", i, err)
			continue
		}
		log.Println(i, ":connect to server ok")
		// data := make([]byte, 65536)
		// var total int
		// for {
		// 	n, err := conn.Write(data)
		// 	if err != nil {
		// 		total += n
		// 		log.Printf("write %d bytes, error:%s\n", n, err)
		// 		// server closed write tcp 127.0.0.1:62373->127.0.0.1:8899: write: broken pipe
		// 		break
		// 	}
		// 	total += n
		// 	log.Printf("write %d bytes this time, %d bytes in total\n", n, total)
		// }
		go handleConnection(conn)
	}
	time.Sleep(10000 * time.Second)

}

func TestWriteConnectionResetByPeer(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", "10.64.49.20:8888")
			if err != nil {
				fmt.Printf("dial error:%s", err)
				return
			}
			fmt.Println("connect to server ok")
			data := []byte("hello ")
			// res := make([]byte, 100)
			// n, err := conn.Read(res)
			// fmt.Println("receive server response: ", string(res), n, err)
			// time.Sleep(5 * time.Second)
			fmt.Println("begin writing...")
			// 在写入数据的过程中服务端断开连接
			for i := 0; i < 1000; i++ {
				time.Sleep(100 * time.Millisecond)
				n, err := conn.Write(data)
				if err != nil {
					fmt.Printf("write %d bytes, error:%s\n", n, err)
					return
				}
				fmt.Println("write ", n, err)

			}
			conn.Close()
			fmt.Println("request done...")
		}()
	}
	wg.Wait()
	// time.Sleep(10000 * time.Second)
}

func handleConnection(c net.Conn) {
	defer c.Close()
	time.Sleep(10 * time.Second)
	fmt.Println("client handle done....")
}

// 连接超过2s报连接超时
func TestDialTimeout(t *testing.T) {
	log.Println("begin dial...")
	conn, err := net.DialTimeout("tcp", ":8899", 2*time.Second)
	if err != nil {
		log.Println("dial error:", err)
		// err:dial tcp :8899: i/o timeout
		return
	}
	defer conn.Close()
	log.Println("dial ok")

}
