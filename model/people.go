package model

type People struct {
	Id int `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name string `json:"name" xorm:"name"`
	Age int `json:"age" xorm:"age"`
}