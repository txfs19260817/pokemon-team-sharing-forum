package routers

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
	"server/pkg/setting"
	"server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(middleware.Cors()) // Cross-Origin Resource Sharing
	r.Use(gin.Recovery())

	r.Static("/assets", "./assets")

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		// 获取标签列表
		apiv1.GET("/teams", v1.GetTeams)
		// 新建标签
		apiv1.POST("/teams", v1.AddTeam)
		// 更新指定标签
		apiv1.PUT("/teams/:id", v1.EditTeam)
		// 删除指定标签
		apiv1.DELETE("/teams/:id", v1.DeleteTeam)
		// 获取模式列表
		apiv1.GET("/formats", v1.GetFormats)
		// 上传图片
		apiv1.POST("/upload", v1.UploadImage)
	}

	return r
}
