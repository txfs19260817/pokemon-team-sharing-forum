package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/e"
)

// Supported formats

var formats = [...]string{
	"VGC 2020",
	"VGC 2019 Ultra Series",
	"VGC 2019 Moon Series",
	"VGC 2019 Sun Series",
	"VGC 2018",
	"VGC 2017",
	"VGC 2016",
	"VGC 2015",
	"VGC 2014",
	"VGC 2013",
	"VGC 2012",
	"VGC 2011",
	"VGC 2010",
	"[Gen 8] Ubers",
	"[Gen 8] OU",
	"[Gen 8] UU",
	"Others",
}

func GetFormats(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": formats,
	})

}
