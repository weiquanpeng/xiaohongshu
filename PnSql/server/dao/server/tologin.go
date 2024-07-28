package server

import (
	"github.com/xiaohongshu/PnSql/server/dao/model"
	"github.com/xiaohongshu/PnSql/server/global"
)

type AuthLogin struct{}

func (a *AuthLogin) CreateUsers(user *model.Users) (err error) {
	err = global.PVA_DB.Create(user).Error
	return err
}

func (a *AuthLogin) GetUsers(user string, passwd string) (b bool) {
	var line model.Users
	global.PVA_DB.Where("username = ? AND password = ?", user, passwd).First(&line)
	if line.Username == "" {
		return false
	}
	return true
}
