package routers

import (
	"gin-project/controller/api"

	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/api")

	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/userList", api.ApiController{}.UserList)
		apiRouters.GET("/pList", api.ApiController{}.PList)
	}
}
