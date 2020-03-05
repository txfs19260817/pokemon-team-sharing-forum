package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"server/pkg/util"
	"time"
)

type Team struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	CreatedAt    time.Time `json:"created_at"`
	Format       string    `json:"format"`
	Pokemon1     string    `json:"pokemon1"`
	Pokemon2     string    `json:"pokemon2"`
	Pokemon3     string    `json:"pokemon3"`
	Pokemon4     string    `json:"pokemon4"`
	Pokemon5     string    `json:"pokemon5"`
	Pokemon6     string    `json:"pokemon6"`
	RentalImgUrl string    `json:"rentalImgUrl"`
	Showdown     string    `json:"showdown"`
	Description  string    `json:"description"`
	State        int       `json:"state"`
}

// 查询
// 获取队伍列表
func GetTeams(pageNum int, pageSize int, maps interface{}) (teams []Team) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Order("created_at desc").Find(&teams)

	return
}

// 获取队伍数目
func GetTeamTotal(maps interface{}) (count int) {
	db.Model(&Team{}).Where(maps).Count(&count)

	return
}

// 根据id获取队伍
func GetTeamById(id int) (team Team) {
	db.Where(&Team{ID: id}).First(&team)

	return
}

// 根据宝可梦获取队伍列表
func GetTeamsByPokemon(pageNum int, pageSize int, maps interface{}, pokemon string) (teams []Team) {
	db.Where(maps). // state
		Where(map[string]interface{}{"pokemon1": pokemon}).
		Or(map[string]interface{}{"pokemon2": pokemon}).
		Or(map[string]interface{}{"pokemon3": pokemon}).
		Or(map[string]interface{}{"pokemon4": pokemon}).
		Or(map[string]interface{}{"pokemon5": pokemon}).
		Or(map[string]interface{}{"pokemon6": pokemon}).
		Offset(pageNum).Limit(pageSize).Order("created_at desc").Find(&teams)

	return
}

// 根据宝可梦获取队伍数目
func GetTeamTotalByPokemon(maps interface{}, pokemon string) (count int) {
	db.Model(&Team{}).Where(maps).
		Where(map[string]interface{}{"pokemon1": pokemon}).
		Or(map[string]interface{}{"pokemon2": pokemon}).
		Or(map[string]interface{}{"pokemon3": pokemon}).
		Or(map[string]interface{}{"pokemon4": pokemon}).
		Or(map[string]interface{}{"pokemon5": pokemon}).
		Or(map[string]interface{}{"pokemon6": pokemon}).
		Count(&count)

	return
}

// 统计每种宝可梦数量
func GetPokemonNumByFormat(format string) map[string]int {
	type Result struct {
		P   string
		Cnt int
	}

	var results []Result
	resultsMap := make(map[string]int)
	for i := 1; i <= 6; i++ {
		// select pokemon%i as p, count(*) as cnt from ptsf_team group by pokemon%i;
		db.Table("ptsf_team").
			Select(fmt.Sprintf("pokemon%d as p, count(*) as cnt", i)).
			Where("format = ? AND state = ?", format, 1).
			Group(fmt.Sprintf("pokemon%d", i)).
			Scan(&results)
		// slice to map
		tempMap := make(map[string]int)
		for _, r := range results {
			tempMap[r.P] = r.Cnt
		}
		// merge temp_map to results_map
		util.UnionKeys(resultsMap, tempMap)
	}
	return resultsMap
}

// 增加
func AddTeam(team *Team) error {
	dbc := db.Create(&team)

	return dbc.Error
}

// 创建队伍时设置时间的Callback
func (team *Team) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("CreatedAt", time.Now())
}
