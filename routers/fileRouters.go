package routers

import (
	"gin-project/controller/file"

	"github.com/gin-gonic/gin"
)

func FileRouters(r *gin.Engine) {
	pingRouters := r.Group("/file")
	{
		pingRouters.POST("/upload", file.FileController{}.Upload)
		pingRouters.POST("/mutilateUpload", file.FileController{}.MutilateUpload)
		pingRouters.POST("/dateUploadFile", file.FileController{}.DateUploadedFile)

	}
}
