package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
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

type IBike interface {
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
	bk = &Bike{}
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

type A struct {
}

func (p *A) Name() string {
	return "I am A"
}

func (p *A) say() {
	fmt.Println(p.Name())
}

func (p *A) sayReal(child interface{}) {
	ref := reflect.ValueOf(child)
	method := ref.MethodByName("Name")
	if method.IsValid() {
		r := method.Call(nil)
		fmt.Println(r[0].String())
	}
}

type B struct {
	A
}

func (c *B) Name() string {
	return "I am B"
}

func TestChild(t *testing.T) {
	b := &B{}
	b.say()
	b.sayReal(b)

	var data []B
	fmt.Println(data, data == nil)
	Console(data)

}

func Console(data interface{}) {
	fmt.Println(data, data == nil, reflect.ValueOf(data).IsNil(), reflect.TypeOf(data))
}

func TestRe1(t *testing.T) {
	tt := &T{}
	fieldA, _ := reflect.TypeOf(tt).Elem().FieldByName("A")
	fieldAPtr := uintptr(unsafe.Pointer(tt)) + fieldA.Offset
	fmt.Println(fieldAPtr, unsafe.Pointer(fieldAPtr))
	intVar := (*int)(unsafe.Pointer(fieldAPtr))
	fmt.Println(intVar, reflect.TypeOf(intVar), reflect.TypeOf(fieldAPtr), tt.A)
	*intVar = 1
	fmt.Println(tt.A)
}
