// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

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

	api_router := router.Group("/api")
	api_user_router := api_router.Group("/user")
	{
		api_user_router.GET("/get", getUser)
		api_user_router.GET("/list", listUser)
	}

	router.Run()
}

func getUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": map[string]string{
			"name":     "bill",
			"password": "111",
			"email":    "bill@email.com",
		},
	})
}

func listUser(c *gin.Context) {

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
}
