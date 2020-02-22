package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/models"
	"server/pkg/e"
	"server/pkg/setting"
	"server/pkg/upload"
)


// @Summary Upload an image 上传图片
// @Accept mpfd
// @Produce json
// @Param url formData string true "Url"
// @Success 200 {object} string "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {object} string "{"code":400,"data":{},"msg":"请求参数错误"}"
// @Failure 500 {object} string "{"code":500,"data":{},"msg":"fail"}"
// @Router /api/v1/upload [post]
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

	// handle request
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		code = e.INVALID_PARAMS
		data = fmt.Sprintf("ERROR (Parsing filename): %s\n", err)
		log.Println(data)
		return
	}

	// validate file
	filename := header.Filename
	log.Println("Uploaded: " + header.Filename)
	if ! upload.CheckImageExt(filename) {
		code = e.ERROR
		data = fmt.Sprintf("ERROR (Validating file): Only " + setting.ImageAllowExts + " are supported image types.")
		log.Println(data)
		return
	}
	if ! upload.CheckImageSize(header) {
		code = e.ERROR
		data = fmt.Sprintf("ERROR (Validating file): The uploaded image is too large (no larger %d than MB).\n", setting.ImageMaxSize)
		log.Println(data)
		return
	}

	// save file
	data, err = upload.SaveImage(file, filename)
	if err != nil {
		code = e.ERROR
		data = fmt.Sprintf("ERROR (Saving file): %s\n", err)
		log.Println(data)
		return
	}
}

// @Summary Upload a base64 image 上传base64图片
// @Accept json
// @Produce json
// @Param base64 body string true "Base64"
// @Success 200 {object} string "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {object} string "{"code":400,"data":{},"msg":"请求参数错误"}"
// @Failure 500 {object} string "{"code":500,"data":{},"msg":"fail"}"
// @Router /api/v1/uploadb64 [post]
func UploadBase64Image(c *gin.Context) {
	code := e.SUCCESS
	data := ""

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}()

	var base64 models.Base64
	err := c.BindJSON(&base64)
	if err != nil {
		code = e.INVALID_PARAMS
		data = fmt.Sprintf("ERROR (Parsing JSON): %s\n", err)
		log.Println(data)
		return
	}

	data, err = upload.DecodeBase64AndSave(base64.Base64Str)
	if err != nil {
		code = e.ERROR
		data = fmt.Sprintf("ERROR (Saving file): %s\n", err)
		log.Println(data)
		return
	}
}
