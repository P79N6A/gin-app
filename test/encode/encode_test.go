package encode

import (
	"testing"
	"encoding/base64"
	"os"
)

func TestEncode(t *testing.T) {

	str := "c29tZSBkYXRhIHdpdGggACBhbmQg77u/"
	data,err:=base64.StdEncoding.DecodeString(str)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("%q\n",data)

	input := []byte("foo\x00bar")
	str = base64.StdEncoding.EncodeToString(input)
	t.Log(str)

	data,_=base64.StdEncoding.DecodeString(str)
	t.Logf("%q\n",data)

	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
	encoder.Close()

	res:=base64.URLEncoding.EncodeToString([]byte("http://yq01-zhangbingbing.epc.baidu" +
		".com:8003/searchbox?action=novel&type=selected"))
	t.Log(res)
}
