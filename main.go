package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"gin-project/routers"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type User struct {
	Username string `json:"username" form:"username"`
	Age      int    `json:"age" form:"age"`
	Sex      string `json:"sex" form:"sex"`
}

func main() {
	r := gin.Default()

	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(200, "首頁")
		})
		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "新聞")
		})
	}

	routers.ApiRouters(r)
	routers.PingRouters(r)

	//響應jsonp  http://localhost:3000/jsonp?callback=test

	//result: test({"title":"我是一個Title","description":"descs"}); 處理跨域

	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{
			Title:       "我是一個Title",
			Description: "descs",
		}
		c.JSONP(200, a)
	})

	r.GET("/query", func(c *gin.Context) {
		username := c.Query("username")
		// age :=
		newAge, _ := strconv.Atoi(c.Query("age"))
		sex := c.DefaultQuery("sex", "male")
		user := &User{}
		if err := c.ShouldBind(&user); err == nil { //User那邊要綁定form
			fmt.Printf("%#v", user)
			c.JSONP(200, user)
		}
		a := &User{
			Username: username,
			Age:      newAge,
			Sex:      sex,
		}
		c.JSONP(200, a)
	})
	// form
	r.POST("/form", func(c *gin.Context) {

		username := c.PostForm("username")
		password := c.PostForm("password")
		// age := c.DefaultPostForm("age", "20")
		age, _ := strconv.Atoi(c.DefaultPostForm("age", "20"))

		c.JSON(200, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})
	//raw data
	r.POST("/form2", func(c *gin.Context) {

		raw, _ := c.GetRawData()
		var buf bytes.Buffer
		if err := json.Compact(&buf, raw); err != nil {
			fmt.Printf("%T", err)
		}
		data := buf.Bytes()
		fmt.Println(string(data))

		var rawData map[string]interface{}

		if err := json.Unmarshal([]byte(string(data)), &rawData); err != nil { // 將字串解析為 map[string]interface{} 型態
			panic(err)
		}

		c.JSON(200, rawData)
	})

	r.POST("/", func(c *gin.Context) {
		c.String(200, "test22222")
	})

	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}
