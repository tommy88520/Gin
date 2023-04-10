package file

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

type FileController struct{}

func (con FileController) Upload(c *gin.Context) {

	file, err := c.FormFile("file")
	dst := path.Join("./static/upload", file.Filename)
	if err == nil {
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": "tommy",
		"dst":      dst,
	})
	fmt.Println("file")
}

func (con FileController) MutilateUpload(c *gin.Context) {

	file, err := c.FormFile("file")
	dst := path.Join("./static/upload", file.Filename)
	if err == nil {
		c.SaveUploadedFile(file, dst)
	}

	file2, err2 := c.FormFile("file2")
	dst2 := path.Join("./static/upload", file2.Filename)
	if err2 == nil {
		c.SaveUploadedFile(file2, dst2)
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": "tommy",
		"dst":      dst,
		"dst2":     dst,
	})
}

func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

func GetUnix() int64 {
	return time.Now().Unix()
}
func (con FileController) DateUploadedFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err == nil {

		extName := path.Ext(file.Filename) //副檔名
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}

		if _, ok := allowExtMap[extName]; !ok {
			fmt.Println("檔案錯誤")
			c.String(200, "Invalid")
			return
		}

		day := GetDay()
		dir := "./static/upload/" + day //創建資料夾
		err := os.MkdirAll(dir, 0666)

		if err != nil {
			fmt.Println("err")
			c.String(200, "failed")
			return
		}

		// fileName := strconv.FormatInt(GetUnix(), 10) + extName
		//todo dir need fix
		dst := path.Join("./static/upload/", file.Filename)
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"username": "tommy",
			"dst":      dst,
		})

	}
}
