package global

import (
	"time"
)

type PvaModel struct {
	ID        uint      `gorm:"primary" json:"id"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
