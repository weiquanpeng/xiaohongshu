package model

import "github.com/xiaohongshu/PnSql/server/global"

type Users struct {
	global.PvaModel
	Username string `gorm:"size:100;not null;default:'';comment:用户名" json:"username"`
	Password string `gorm:"size:100;not null;default:'';comment:密码" json:"password"`
	Power    uint   `gorm:"not null;default:0;comment:权限" json:"power"`
}

func (Users) TableName() string {
	return "users"
}
