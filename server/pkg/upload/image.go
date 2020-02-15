package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"server/pkg/file"
	"server/pkg/setting"
	"strings"
	"time"
)

func GetImageFullUrl(name string) string {
	return setting.ImagePrefixUrl + "/" + GetImagePath() + name
}

func GetImagePath() string {
	return setting.ImageSavePath
}

func CheckImageExt(fileName string) bool {
	return strings.Contains(setting.ImageAllowExts, strings.ToLower(file.GetExt(fileName)))
}

//func CheckImageSize(f multipart.File) bool {
// // 坑，读文件到EOF后面没法Copy了
//	size, err := file.GetSize(f)
//	if err != nil {
//		log.Println(err)
//		return false
//	}
//
//	return size <= setting.ImageMaxSize
//}

func CheckImageSize(h *multipart.FileHeader) bool {
	return h.Size <= int64(setting.ImageMaxSize)
}


// Save a file
func SaveImage(f multipart.File, fileName string) (fileUrlStr string, err error) {
	fileDir := fmt.Sprintf("./%s%s", GetImagePath(), time.Now().Format("20060102"))
	err = file.MkDir(fileDir)
	if err != nil {
		return "", err
	}

	fileName = file.Rename(fileName)
	filePathStr := fileDir + "/" + fileName
	out, err := os.Create(filePathStr)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, f)
	if err != nil {
		return "", err
	}
	return GetImageFullUrl(fileName), nil
}
