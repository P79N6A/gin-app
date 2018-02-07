package dao

import (
	//依赖包https://github.com/go-sql-driver/mysql
	"database/sql"
	"fmt"

	"../model/user"
)

func getUser() {
	//打开连接
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	checkError(err)

	//查询数据
	rows, err := db.Query("select * from user")
	checkError(err)

	//遍历查询结果
	for rows.Next() {
		var user user.User
		err = rows.Scan(&id, &name, &password, &email)
		checkError(err)
		fmt.Println(id, name, password, email)
	}

}

func getUsers() (users []user.User, err error) {
	users = make([]user.User, 0)
	//打开连接
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	checkError(err)

	//查询数据
	rows, err := db.Query("select * from user")
	checkError(err)

	//遍历查询结果
	for rows.Next() {
		var user user.User
		err = rows.Scan(&user.id, &user.name, &user.password, &user.email)
		checkError(err)
		users = append(users, user)
		fmt.Println(id, name, password, email)
	}
	return users

}

func deleteUser(id int) bool {
	//删除数据
	stmt, err = db.Prepare("delete from user where id=?")
	checkError(err)

	res, err = stmt.Exec(id)
	affect, err = res.RowsAffected()
	checkError(err)

	fmt.Println(affect)
	db.Close()
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
