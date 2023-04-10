package routers

import (
	"gin-project/controller/ping"

	"github.com/gin-gonic/gin"
)

func PingRouters(r *gin.Engine) {
	pingRouters := r.Group("/ping")
	{
		pingRouters.GET("/", ping.PingController{}.Index)
		pingRouters.GET("/2", ping.PingController{}.Second)

		pingRouters.GET("/3", ping.PingController{}.Third)
		pingRouters.GET("/4", ping.PingController{}.Fourth)

	}
}
