// main.go
package main

import (
	"net/http"

	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)
import userService "./service"

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"title": "Main Website",
			"body":  "this is a go page!!!",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s", name)
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)

	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})

	})

	//api路由组
	api_router := router.Group("/api")
	api_user_router := api_router.Group("/user")
	{
		api_user_router.GET("/get", getUser)
		api_user_router.GET("/list", listUser)
		api_user_router.POST("/add", addUser)
		api_user_router.GET("/delete", deleteUser)
	}

	router.Run()
}

/**
删除用户处理器
*/
func deleteUser(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	fmt.Println("id: ", id)
	userId, err := strconv.Atoi(id)
	result, err := userService.DeleteUser(userId)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    result,
		})
	} else {
		log.Fatal(err)
	}
}

/**
添加用户处理器
*/
func addUser(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	params := make(map[string]string)
	params["name"] = name
	params["password"] = password
	params["email"] = email
	res, err := userService.AddUser(params)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    res,
		})
	} else {
		log.Fatal(err)
	}

}

/**
获取用户处理器
*/
func getUser(c *gin.Context) {

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
	user, err := userService.GetUser(userId)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    user,
		})
	} else {
		log.Fatal(err)
	}
}

/**
获取用户处理器
*/
func listUser(c *gin.Context) {

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
	users, err := userService.GetUsers()
	fmt.Println(users, err)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    users,
		})
	} else {
		log.Fatalln(err)
	}
}
