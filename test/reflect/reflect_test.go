package reflect

import (
	"fmt"
	"reflect"
	"testing"
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

type INovel interface {
	GetName()
}

type IBike interface{
	GetColor()
}

type Book struct {
}

func (book Book) GetName() {
	fmt.Println("test book")
}

type Bike struct {
}

func (bike *Bike) GetColor() {
	fmt.Println("green")
}

func Test2(t *testing.T) {
	objMap := make(map[string]interface{})
	objMap["book"] = Book{}
	objMap["bike"] = Bike{}
	t1 := reflect.TypeOf(objMap["book"])
	book := reflect.New(t1)
	obj := book.Elem().Interface().(Book)
	obj.GetName()

	t2 := reflect.TypeOf(objMap["bike"])
	obj1 := reflect.New(t2)
	bike := obj1.Elem().Interface().(Bike)
	bike.GetColor()

	var novel INovel
	novel = Book{}
	novel.GetName()

	var bk IBike
	bk=&Bike{}
	bk.GetColor()
}

func TestSome(t *testing.T) {
	var UIDArr []int
	for idx := 0; idx < 1000; idx++ {
		if idx > 0 && idx%100 == 0 && len(UIDArr) > 0 {
			fmt.Println(UIDArr)
			UIDArr = []int{}
		}
		UIDArr = append(UIDArr, idx)
	}
	if len(UIDArr) > 0 {
		fmt.Println(UIDArr)

	}
}