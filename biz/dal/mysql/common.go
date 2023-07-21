package mysql

import (
	"database/sql"
	"time"
)

// BaseModel 参考 gorm.Model 自定义一个公共的 model
type BaseModel struct {
	ID        uint         `json:"id" gorm:"type:bigint;primaryKey"`
	CreatedAt time.Time    `json:"createdAt" gorm:"format:2006-01-02 15:04:05"`
	UpdatedAt time.Time    `json:"updatedAt" gorm:"format:2006-01-02 15:04:05"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}
