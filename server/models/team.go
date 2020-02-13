package models

import (
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm"
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
	Introduction string    `json:"introduction"`
	State        int       `json:"state"`
}

// tools
func (team *Team) TeamValidator(err map[string]string) bool {
	valid := validation.Validation{}
	valid.MaxSize(team.Title, 64, "Title").Message("标题最长为64字符")
	valid.MaxSize(team.Author, 64, "Author").Message("作者名最长为64字符")
	// TODO: format限制在一个数组内
	// TODO: 验证6只pm名字合法性
	// TODO: 验证showdown语法合法性
	valid.MaxSize(team.Introduction, 300, "Introduction").Message("简介最长为300字符，过长建议附外部链接")
	valid.Range(team.State, 0, 1, "State").Message("状态只允许0或1")
	if valid.HasErrors() {
		for _, e := range valid.Errors {
			err[e.Key] = e.Message
		}
		return false
	}
	return true
}


// 查询
// 获取队伍列表
func GetTeams(pageNum int, pageSize int, maps interface{}) (teams []Team) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&teams)

	return
}

// 获取队伍数目
func GetTeamTotal(maps interface{}) (count int) {
	db.Model(&Team{}).Where(maps).Count(&count)

	return
}

func AddTeam(team *Team) error {
	dbc := db.Create(&team)

	return dbc.Error
}

// 增加
// 创建队伍时设置时间的Callback
func (team *Team) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("CreatedAt", time.Now())
}
