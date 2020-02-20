package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
	"server/models"
	"server/pkg/e"
	"server/pkg/setting"
	"server/pkg/util"
)

// 获取多个队伍
func GetTeams(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTeams(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTeamTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetTeamById(c *gin.Context) {
	var id = -1
	if arg := c.Param("id"); arg != "" {
		id = com.StrTo(arg).MustInt()
		if id == 0 { // Would get the data whose id=1
			id = -1
		}
	}

	code := e.SUCCESS

	data := models.GetTeamById(id)
	if data.ID == 0 {
		code = e.INVALID_PARAMS
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetTeamByFormat(c *gin.Context)  {
	data := make(map[string]interface{})
	format := c.Param("format")

	code := e.SUCCESS

	data["lists"] = models.GetTeamByFormat(format)
	data["total"] = models.GetTeamTotalByFormat(format)

	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增队伍
func AddTeam(c *gin.Context) {
	var team models.Team
	code := e.SUCCESS
	data := make(map[string]string, 20)

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}()

	if err := c.ShouldBindJSON(&team); err != nil {
		code = e.INVALID_PARAMS
		log.Printf("ERROR: %s\n", err)
		return
	}
	if valid := team.TeamValidator(data); !valid {
		code = e.ERROR
		log.Println("An error occurred when validating the form. See `data` for details")
		return
	}
	if err := models.AddTeam(&team); err != nil {
		code = e.ERROR
		log.Printf("ERROR: %s\n", err)
		return
	}
}

// 修改队伍
func EditTeam(c *gin.Context) {
}

// 删除队伍
func DeleteTeam(c *gin.Context) {
}
