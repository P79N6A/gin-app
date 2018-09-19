package main

import (
	"os"
	"encoding/json"
	"testing"
	"reflect"
	"fmt"
	"strconv"
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

type ICoupon interface {
	parse(val []interface{})
	dealDetail()
	Id() int
}

type Coupon1 struct {
}

func (this *Coupon1) Id() int {
	return 1
}
func (this *Coupon1) parse(val []interface{}) {
	for idx, item := range val {
		fmt.Println(idx, item.([]interface{})[0], item.([]interface{})[1], item.([]interface{})[2], item.([]interface{})[3])
	}
}
func (this *Coupon1) dealDetail() {

}

type Coupon2 struct {
}

func (this *Coupon2) Id() int {
	return 8
}
func (this *Coupon2) parse(val []interface{}) {
	for idx, item := range val {
		fmt.Println(idx, item.([]interface{})[0], item.([]interface{})[1], item.([]interface{})[2], item.
		([]interface{})[3])
	}
}
func (this *Coupon2) dealDetail() {

}

func getCouponHandler(couponId int) ICoupon {
	var handlerMap = make(map[int]ICoupon)
	handlerMap[1] = &Coupon1{}
	handlerMap[8] = &Coupon2{}
	return handlerMap[couponId]

}
func TestBean(t *testing.T) {
	var detail string = "{\"8\":[[0,1533683940,1533687540,37]],\"1\":[[300,\"20180907\",\"20181007\",\"37\"],[300,\"20180907\",\"20181007\",\"37\"],[300,\"20180907\",\"20181007\",\"37\"],[300,\"20180907\",\"20181007\",\"37\"],[300,\"20180907\",\"20181007\",\"37\"],[300,\"20180907\",\"20181007\",\"37\"]]}"
	var detailJson map[string]interface{}
	json.Unmarshal([]byte(detail), &detailJson)
	t.Log(detailJson)
	for key, val := range detailJson {
		t.Log(key, val, reflect.TypeOf(val))
		couponId, _ := strconv.Atoi(key)
		handler := getCouponHandler(couponId)
		t.Log(handler)
		handler.parse(val.([]interface{}))
	}
}
