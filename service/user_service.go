package service

import (
	"../dao"
	"../model"
)

func init() {}

/**
获取用户列表
*/
func GetUsers() (users []model.User, err error) {
	users, err = dao.GetUsers()
	return users, err
}

/**
获取用户
*/
func GetUser(id int) (user model.User, err error) {
	user, err = dao.GetUser(id)
	return user, err
}

/**
添加用户
*/
func AddUser(params map[string]string) (result bool, err error) {
	user := model.User{Name: params["name"], Password: params["password"], Email: params["email"]}
	result, err = dao.AddUser(user)
	return result, err

}

/**
根据id删除用户
*/
func DeleteUser(id int) (result bool, err error) {
	result, err = dao.DeleteUser(id)
	return result, err
}
