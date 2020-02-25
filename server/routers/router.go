package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	_ "server/docs"
	"server/middleware"
	"server/pkg/setting"
	"server/routers/api/v1"
	"time"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	if setting.RunMode == "release" {
		// https
		r.Use(middleware.TlsHandler())

		// limit access rate by custom key (here is IP) and rate
		r.Use(limit.NewRateLimiter(func(c *gin.Context) string {
			return c.ClientIP() // limit rate by client ip
		}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
			// limit 1/60 qps/clientIp and permit bursts of at most 10 tokens,
			// and the limiter liveness time duration is 1 hour
			// https://www.cyhone.com/articles/usage-of-golang-rate/
			return rate.NewLimiter(rate.Every(60*time.Second), 10), time.Hour
		}, func(c *gin.Context) {
			if c.Request.Method == http.MethodPost {
				c.AbortWithStatus(429) // handle exceed rate limit request
			}
		}))

	} else if setting.RunMode == "debug" {
		r.Use(middleware.Cors()) // Cross-Origin Resource Sharing
	} else {
		log.Fatalln("Unknown RunMode")
	}

	r.Static(setting.RelativePath, setting.Root)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiv1 := r.Group("/api/v1")
	{
		// 获取队伍列表
		apiv1.GET("/teams", v1.GetTeams)
		// 根据id获取队伍
		apiv1.GET("/teams/:id", v1.GetTeamById)
		// 根据模式获取队伍
		apiv1.GET("/formats/:format", v1.GetTeamByFormat)
		// 根据模式获取队伍
		apiv1.GET("/pokemon/:pokemon", v1.GetTeamByPokemon)
		// 新建队伍
		apiv1.POST("/teams", v1.AddTeam)

		// 上传图片
		apiv1.POST("/upload", v1.UploadImage)
		// 上传base64
		apiv1.POST("/uploadb64", v1.UploadBase64Image)
	}

	return r
}
