package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	fmt.Println("我是分離中間件")
	c.Set("username", "Tommy")

	cCp := c.Copy() //開啟協成 使用c要copy
	go func() {
		time.Sleep(time.Second)
		fmt.Println(cCp.Request.URL.Path)
	}()
}
