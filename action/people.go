package action

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-app/service"
	"gin-app/library/logger"
)

var mylog logger.BLog = logger.GetLogger()
/**
删除用户处理器
*/
func DeletePeople(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	fmt.Println("id: ", id)
	peopleId, err := strconv.Atoi(id)
	result, err := service.DeletePeople(peopleId)
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
func AddPeople(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	params := make(map[string]string)
	params["name"] = name
	params["age"] = age
	res, err := service.AddPeople(params)
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
func GetPeople(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	fmt.Println("id: ", id)
	peopleId, err := strconv.Atoi(id)
	user, err := service.GetPeople(peopleId)
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

/**
获取用户处理器
*/
func ListPeople(c *gin.Context) {
	users, err := service.GetPeoples()
	mylog.Debug(users)
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
