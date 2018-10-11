package tnet

import (
	"testing"
	"net/http"
	"net/url"
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
	ck,_:=req.Cookie("name")
	t.Log(ck.String())
	t.Log(http.CanonicalHeaderKey("accept-encoding"),http.CanonicalHeaderKey("if-none-match"))
}

func TestUrl(t *testing.T)  {
	urlStr := "http://yq01-zhangbingbing.epc.baidu.com:8003/searchbox?action=novel&type=selected"
	res:=url.QueryEscape(urlStr)
	t.Log(res)
	res,_ = url.QueryUnescape(res)
	t.Log(res)

	u,_:=url.Parse(urlStr)
	t.Log(u.User,u.Scheme,u.Host,u.Query(),u.RawQuery,u.Path)

	v:=url.Values{}
	v.Set("name", "bill")
	v.Set("age", "30")
	v.Set("email", "bill@126.com")
	t.Log(v.Encode())

	v,_=url.ParseQuery(u.RawQuery)
	t.Log(v)

}
