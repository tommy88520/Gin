package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
}

func (con ApiController) Index(c *gin.Context) {
	username, _ := c.Get("username")
	fmt.Println(username)
	v, ok := username.(string)
	if ok {
		c.String(200, "我是一個api街口"+v)
	} else {
		c.String(200, "獲取失敗")
	}

}
func (con ApiController) UserList(c *gin.Context) {
	c.String(200, "我是一個userlist")
}

func (con ApiController) PList(c *gin.Context) {
	c.String(200, "我是一個plist")
}
