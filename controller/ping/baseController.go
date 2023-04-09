package ping

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (con BaseController) success(c *gin.Context) {
	c.String(200, "success")
}
