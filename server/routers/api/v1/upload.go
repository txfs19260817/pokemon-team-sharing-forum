package v1

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"server/pkg/e"
	"server/pkg/setting"
	"strings"
)

func RootPath() string {
	path, _ := os.Getwd()
	return path
}

func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := ""

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}()

	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		code = e.INVALID_PARAMS
		log.Printf("ERROR (Parsing filename): %s\n", err)
		return
	}
	filename := header.Filename
	log.Println(header.Filename)
	out, err := os.Create("./assets/" + filename)
	if err != nil {
		code = e.ERROR
		log.Printf("ERROR (Creating file): %s\n", err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		code = e.ERROR
		log.Printf("ERROR (Copying file): %s\n", err)
		return
	}
	data = filepath.ToSlash(filepath.Join(setting.HTTPURL, "assets", filename))
	data = strings.Replace(data,"http:/", "http://", 1)

}
