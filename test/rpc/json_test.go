package rpc

import (
	"testing"
	"encoding/json"
)

func TestJson(t *testing.T) {
	b := []byte(`{
		"Name":"bill",
		"Age":30,
		"Email":"bill@126.com"
	}`)
	var r interface{}
	err := json.Unmarshal(b, &r)
	if err == nil {
		t.Log(r)
		info, ok := r.(map[string]interface{})
		if ok {
			t.Log(info["Name"])
			for k, v := range info {
				t.Log(k, v)
			}
		}
	}

	var user User
	err = json.Unmarshal(b, &user)
	t.Log(user)

}

type User struct {
	Name  string
	Age   int
	Email string
}
