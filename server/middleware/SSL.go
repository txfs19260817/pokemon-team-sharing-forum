package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"server/pkg/setting"
)

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     fmt.Sprintf("%s:%d", setting.URL, setting.HTTPPort), //host
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
