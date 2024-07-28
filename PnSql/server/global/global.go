package global

import "gorm.io/gorm"

var (
	PVA_DB *gorm.DB
)

type Pva struct {
	System System
	Mysql  Mysql
}

var P_cfg = new(Pva)
