package action

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-app/service"
	"strconv"
)

//var mylog logger.BLog = logger.GetLogger()
/**
获取用户处理器
*/
func ListUser(c *gin.Context) {

	/*
		users := make([]map[string]string, 0)
		m := map[string]string{
			"name":     "bill",
			"password": "111",
			"email":    "bill@email.com",
		}
		users = append(users, m)

		users = append(users, map[string]string{
			"name":     "bing",
			"password": "222",
			"email":    "bing@email.com",
		})
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    users,
		})
	*/
	users, err := service.GetUsers()
	mylog.Info(users)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    users,
		})
	} else {
		mylog.Fatal(err)
	}
}

/**
删除用户处理器
*/
func DeleteUser(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	fmt.Println("id: ", id)
	userId, err := strconv.Atoi(id)
	result, err := service.DeleteUser(userId)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    result,
		})
	} else {
		mylog.Fatal(err)
	}
}

/**
添加用户处理器
*/
func AddUser(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	params := make(map[string]string)
	params["name"] = name
	params["password"] = password
	params["email"] = email
	res, err := service.AddUser(params)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    res,
		})
	} else {
		mylog.Fatal(err)
	}

}

/**
获取用户处理器
*/
func GetUser(c *gin.Context) {

	/*
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data": map[string]string{
				"name":     "bill",
				"password": "111",
				"email":    "bill@email.com",
			},
		})
	*/
	id := c.DefaultQuery("id", "0")
	fmt.Println("id: ", id)
	userId, err := strconv.Atoi(id)
	user, err := service.GetUser(userId)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    user,
		})
	} else {
		mylog.Fatal(err)
	}
}
