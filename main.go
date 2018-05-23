// main.go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-app/action"
	"gin-app/library/logger"
	"gin-app/library/config"
	"fmt"
)

var mylog logger.BLog = logger.GetLogger()

func main() {
	fmt.Println(config.GetLogConfig())
	fmt.Println(config.GetAppConfig())

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
		api_user_router.GET("/get", action.GetUser)
		api_user_router.GET("/list", action.ListUser)
		api_user_router.POST("/add", action.AddUser)
		api_user_router.GET("/delete", action.DeleteUser)
	}
	api_people_router := api_router.Group("/people")
	{
		api_people_router.GET("/get", action.GetPeople)
		api_people_router.GET("/list", action.ListPeople)
		api_people_router.POST("/add", action.AddPeople)
		api_people_router.GET("/delete", action.DeletePeople)
	}

	router.Run()
}
