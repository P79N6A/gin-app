package reflect

import (
	"testing"
	"reflect"
)

func TestReflect(t *testing.T) {
	var x float64 = 3.4
	t.Log("type:", reflect.TypeOf(x))

	v := reflect.ValueOf(x)
	t.Log("type:", v.Type())
	t.Log("kind is float64:", v.Kind() == reflect.Float64)
	t.Log("value:", v.Float())

	t.Log("can set:", v.CanSet())

	p := reflect.ValueOf(&x)
	t.Log("type of p:", p.Type())
	t.Log("p can set:", p.CanSet())
	v = p.Elem()
	t.Log("v can set:", v.CanSet())
	t.Log("value:", v.Float())
	v.SetFloat(8.1)
	t.Log("value after set:", v.Float())
	t.Log(v.Interface())

	ts := T{111, "hello"}
	s := reflect.ValueOf(&ts).Elem()
	typeofT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		t.Logf("%d: %s %s = %v\n", i, typeofT.Field(i).Name, f.Type(), f.Interface())
	}
}

type T struct {
	A int
	B string
}
