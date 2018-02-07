package service

import (
	"../dao"
	"../model"
)

func init() {}

func GetUsers() (users []model.User, err error) {
	users, err = dao.GetUsers()
	return users, err
}

func GetUser(id int) (user model.User, err error) {
	user, err = dao.GetUser(id)
	return user, err
}

func AddUser(params map[string]string) (result bool, err error) {
	user := model.User{Name: params["name"], Password: params["password"], Email: params["email"]}
	result, err = dao.AddUser(user)
	return result, err

}

func DeleteUser(id int) (result bool, err error) {
	result, err = dao.DeleteUser(id)
	return result, err
}
