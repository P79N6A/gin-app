package main

import (
	"os"
	"encoding/json"
	"testing"
)

func Test_Json(t *testing.T) {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			t.Error(err)
			return
		}
		for k := range v {
			if k != "title" {
				v[k] = nil
			}
		}
		if err := enc.Encode(&v); err != nil {
			t.Error(err)
		}
	}
}
