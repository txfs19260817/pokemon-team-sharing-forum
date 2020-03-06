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
	"server/pkg/validator"
	"strings"
)

// @Summary Get multiple teams 获取多个队伍
// @Produce  json
// @Param state query int false "State"
// @Param page query int false "State"
// @Success 200 {object} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/teams [get]
func GetTeams(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}
	maps["state"] = state

	code := e.SUCCESS

	data["lists"] = models.GetTeams(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTeamTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary Get a team by id 根据id获取1个队伍
// @Produce  json
// @Param id path int true "ID"
// @Param state query int false "State"
// @Param page query int false "State"
// @Success 200 {object} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/teams/{id} [get]
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

// @Summary Get multiple teams by format 根据模式获取多个队伍
// @Produce  json
// @Param format path string true "Format"
// @Param state query int false "State"
// @Param page query int false "State"
// @Success 200 {object} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/formats/{format} [get]
func GetTeamByFormat(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	maps["format"] = c.Param("format")

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}
	maps["state"] = state

	code := e.SUCCESS

	data["lists"] = models.GetTeams(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTeamTotal(maps)

	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary Get multiple teams by Pokemon 根据宝可梦获取多个队伍
// @Produce  json
// @Param pokemon path string true "Pokemon"
// @Param state query int false "State"
// @Param page query int false "State"
// @Success 200 {object} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/pokemon/{pokemon} [get]
func GetTeamByPokemon(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	pokemon := c.Param("pokemon")
	pokemon = strings.Title(pokemon)

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}
	maps["state"] = state

	code := e.SUCCESS

	data["lists"] = models.GetTeamsByPokemon(util.GetPage(c), setting.PageSize, maps, pokemon)
	data["total"] = models.GetTeamTotalByPokemon(maps, pokemon)

	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary Add a team 新增队伍
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param format body string true "Format"
// @Param pokemon1 body string true "Pokemon1"
// @Param pokemon2 body string false "Pokemon2"
// @Param pokemon3 body string false "Pokemon3"
// @Param pokemon4 body string false "Pokemon4"
// @Param pokemon5 body string false "Pokemon5"
// @Param pokemon6 body string false "Pokemon6"
// @Param rental_img_url body string false "RentalImgUrl"
// @Param showdown body string false "Showdown"
// @Param description body string false "Description"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} string "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {object} string "{"code":400,"data":{},"msg":"请求参数错误"}"
// @Failure 500 {object} string "{"code":500,"data":{},"msg":"fail"}"
// @Router /api/v1/teams [post]
func AddTeam(c *gin.Context) {
	var team models.Team
	code := e.SUCCESS
	data := ""

	defer func() {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}()

	//if err := c.ShouldBindJSON(&team); err != nil {
	//	code = e.INVALID_PARAMS
	//	log.Printf("ERROR: %s\n", err)
	//	return
	//}
	data, valid := validator.TeamValidator(&team, c.Request)
	if !valid {
		code = e.ERROR
		log.Printf("Validation error: %s", data)
		return
	}
	if err := models.AddTeam(&team); err != nil {
		code = e.ERROR
		log.Printf("ERROR: %s\n", err)
		return
	}
}

func GetPokemonStat(c *gin.Context)  {
	code := e.SUCCESS
	data := make(map[string]int)
	format := c.Param("format")
	if format == "" {
		code = e.INVALID_PARAMS
	} else{
		data = models.GetPokemonNumByFormat(format)
	}

	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}