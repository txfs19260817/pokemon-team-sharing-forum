package upload

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"server/pkg/file"
	"server/pkg/setting"
	"server/pkg/util"
	"strings"
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
	fileDir := fmt.Sprintf("./%s", GetImagePath())
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

	if _, err = io.Copy(out, f); err != nil {
		return "", err
	}

	if err = out.Close(); err != nil {
		return "", err
	}

	if err = util.Rescale(filePathStr); err != nil {
		return "", err
	}

	return GetImageFullUrl(fileName), nil
}

// base64
func DecodeBase64AndSave(b string) (fileUrlStr string, err error) {
	// check
	if len(b) < 10 {
		return "", errors.New("Empty base64 string. ")
	}
	var ext string
	if b[11] == 'j' {
		b = b[23:]
		ext = ".jpg"
	} else if b[11] == 'p' {
		b = b[22:]
		ext = ".png"
	} else if b[11] == 'g' {
		b = b[22:]
		ext = ".gif"
	} else {
		return "", errors.New("Not supported base64 image format. ")
	}

	// decode
	data, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return "", err
	}

	// save
	fileDir := fmt.Sprintf("./%s", GetImagePath())
	err = file.MkDir(fileDir)
	if err != nil {
		return "", err
	}

	fileName := file.Rename(b[:32] + ext)
	filePathStr := fileDir + "/" + fileName
	f, _ := os.OpenFile(filePathStr, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	n, err := f.Write(data)
	if err != nil {
		return "", err
	} else if n == 0 {
		return "", errors.New("Wrote an empty file. ")
	}

	return GetImageFullUrl(fileName), nil
}
