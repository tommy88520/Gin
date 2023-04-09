package ping

import "github.com/gin-gonic/gin"

type PingController struct {
	BaseController
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (con PingController) Index(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"success": true,
		"msg":     "ping",
	})
}
func (con PingController) Second(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"msg":     "ping",
	})
}

func (con PingController) Third(c *gin.Context) {
	a := &Article{
		Title:       "我是一個Title",
		Description: "descs",
	}
	c.JSON(200, a)
}

func (con PingController) Fourth(c *gin.Context) {
	con.success(c) //還可以繼承BaseController
}
