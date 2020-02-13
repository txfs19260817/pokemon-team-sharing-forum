package routers

import (
	"github.com/gin-gonic/gin"
	"server/pkg/setting"
	"server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/teams", v1.GetTeams)
		//新建标签
		apiv1.POST("/teams", v1.AddTeam)
		//更新指定标签
		apiv1.PUT("/teams/:id", v1.EditTeam)
		//删除指定标签
		apiv1.DELETE("/teams/:id", v1.DeleteTeam)
	}

	return r
}