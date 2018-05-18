package dao

import (
	"github.com/go-xorm/xorm"
	"fmt"
	"github.com/mgutz/logxi/v1"
	"github.com/go-xorm/core"
	"gin-app/model"
	_ "github.com/go-sql-driver/mysql"
)

//const (
//	host     string = "127.0.0.1"
//	port     int    = 3306
//	database string = "test"
//)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	if err != nil {
		log.Error("-----------connect database fail", err)
	} else {
		//日志打印SQL
		engine.ShowSQL(true)
		//设置连接池的空闲数大小
		engine.SetMaxIdleConns(5)
		//设置最大打开连接数
		engine.SetMaxOpenConns(5)
		//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
		engine.SetTableMapper(core.SnakeMapper{})
	}
}

func AddPeople(people model.People) (result bool, err error) {
	affect, err := engine.Insert(people)
	return affect > 0, err
}

func GetPeople(id int) (people model.People, err error) {
	result, err := engine.Id(id).Get(&people)
	fmt.Println(result)
	return
}

func GetPeoples() (peoples []model.People, err error) {
	peoples = make([]model.People, 0)
	err = engine.Find(&peoples)
	return
}

func DeletePeople(id int) (result bool, err error) {
	res, err := engine.Delete(&model.People{Id: id})
	return res > 0, err
}
