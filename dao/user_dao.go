package dao

import (
	//依赖包https://github.com/go-sql-driver/mysql
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"../model"
)

//用户名
const userName string = "root"

//密码
const password = "123456"

//host和端口
const (
	host     string = "127.0.0.1"
	port            = 3306
	database        = "test"
)

//数据库连接串
var jdbcUrl string

//导入包完成后自动初始化
func init() {
	jdbcUrl = userName + ":" + password + "@(" + host + ":" + strconv.Itoa(port) + ")/" + database + "?charset=utf8"
}

/**
添加用户
*/
func AddUser(user model.User) (result bool, err error) {
	//打开连接
	db, err := sql.Open("mysql", jdbcUrl)
	//最后关系连接
	defer db.Close()
	checkError(err)

	//插入数据
	stmt, err := db.Prepare("insert user set name=?,password=?,email=?")
	checkError(err)

	res, err := stmt.Exec(user.Name, user.Password, user.Email)
	checkError(err)

	affect, err := res.RowsAffected()

	return affect > 0, err
}

/**
获取用户
*/
func GetUser(id int) (user model.User, err error) {

	//打开连接
	db, err := sql.Open("mysql", jdbcUrl)
	//最后关系连接
	defer db.Close()
	checkError(err)

	//查询数据
	stmt, err := db.Prepare("select * from user where id=?")
	checkError(err)
	fmt.Println("id = ", id)
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Password, &user.Email)
	checkError(err)

	return user, err
}

/**
获取用户列表
*/
func GetUsers() (users []model.User, err error) {
	users = make([]model.User, 0)
	//打开连接
	db, err := sql.Open("mysql", jdbcUrl)
	//最后关系连接
	defer db.Close()
	checkError(err)

	//查询数据
	rows, err := db.Query("select * from user")
	checkError(err)

	//遍历查询结果
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		checkError(err)
		users = append(users, user)

	}

	return users, err

}

/**
删除用户
*/
func DeleteUser(id int) (result bool, err error) {
	//打开连接
	db, err := sql.Open("mysql", jdbcUrl)
	//最后关系连接
	defer db.Close()
	checkError(err)
	//删除数据
	stmt, err := db.Prepare("delete from user where id=?")
	checkError(err)

	res, err := stmt.Exec(id)
	fmt.Println(res, err)
	affect, err := res.RowsAffected()
	checkError(err)

	fmt.Println(affect)

	return affect > 0, err
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
