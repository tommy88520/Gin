package api

import "github.com/gin-gonic/gin"

type ApiController struct {
}

func (con ApiController) Index(c *gin.Context) {
	c.String(200, "我是一個api街口")

}
func (con ApiController) UserList(c *gin.Context) {
	c.String(200, "我是一個userlist")
}

func (con ApiController) PList(c *gin.Context) {
	c.String(200, "我是一個plist")
}
