package service

import (
	"gin-app/model"
	"gin-app/dao"
	"strconv"
)

func init() {

}

func GetPeoples() (peoples []model.People, err error) {
	peoples, err = dao.GetPeoples()
	return
}
func GetPeople(id int) (people model.People, err error) {
	people, err = dao.GetPeople(id)
	return
}
func AddPeople(params map[string]string) (result bool, err error) {
	age, _ := strconv.Atoi(params["age"])
	people := model.People{Name: params["name"], Age: age}
	result, err = dao.AddPeople(people)
	return
}
func DeletePeople(id int) (result bool, err error) {
	result, err = dao.DeletePeople(id)
	return
}
