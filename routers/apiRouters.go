package routers

import (
	"gin-project/controller/api"
	"gin-project/middlewares"

	"github.com/gin-gonic/gin"
)

//	func innerMiddleWare(c *gin.Context) {
//		fmt.Println("1234566")
//	}
func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/api", middlewares.InitMiddleware)

	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/userList", api.ApiController{}.UserList)
		apiRouters.GET("/pList", api.ApiController{}.PList)
	}
}
