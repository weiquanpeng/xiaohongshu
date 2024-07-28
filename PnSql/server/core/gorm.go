package core

import (
	"github.com/xiaohongshu/PnSql/server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	m := global.P_cfg.Mysql
	if m.Dbname == "" || m.Username == "" || m.Password == "" || m.Host == "" || m.Port == "" {
		panic("数据库配置参数不全")
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         128,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(50)
		sqlDB.SetMaxOpenConns(1000)
		return db
	}

}
