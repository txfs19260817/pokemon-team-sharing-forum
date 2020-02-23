package util

import (
	"errors"
	"golang.org/x/image/draw"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// resize image to 768 * 432
// https://text.baldanders.info/golang/resize-image/
func Rescale(filePath string) (err error) {
	//open original image file
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	//decode image
	imgSrc, ext, err := image.Decode(file)
	if err != nil {
		return
	}

	//rectange of image
	rctSrc := imgSrc.Bounds()

	//rescale to 768 * 432
	imgDst := image.NewRGBA(image.Rect(0, 0, 768, 432))
	draw.CatmullRom.Scale(imgDst, imgDst.Bounds(), imgSrc, rctSrc, draw.Over, nil)

	//create resized image file
	dst, err := os.Create(filePath) //dst file path is same as src
	if err != nil {
		return
	}
	defer dst.Close()

	//encode resized image
	switch ext {
	case "jpeg":
		if err = jpeg.Encode(dst, imgDst, &jpeg.Options{Quality: 80}); err != nil {
			return
		}
	case "gif":
		if err = gif.Encode(dst, imgDst, nil); err != nil {
			return
		}
	case "png":
		if err = png.Encode(dst, imgDst); err != nil {
			return
		}
	default:
		err = errors.New("format error")
	}
	return nil
}
