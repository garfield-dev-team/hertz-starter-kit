package mysql

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel 参考 gorm.Model 自定义一个公共的 model
type BaseModel struct {
	ID        uint           `json:"id" gorm:"type:bigint;primaryKey"`
	CreatedAt time.Time      `json:"createdAt" gorm:"not null;format:2006-01-02 15:04:05"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"not null;format:2006-01-02 15:04:05"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
