package mysql

import (
	"database/sql"
	"time"
)

// BaseModel 参考 gorm.Model 自定义一个公共的 model
// 其中 ID 为 uint 表示自增主键，为 string 表示 uuid
type BaseModel[T uint | string] struct {
	ID        T            `json:"id" gorm:"type:int(10);primaryKey"`
	CreatedAt time.Time    `json:"createdAt" gorm:"format:2006-01-02 15:04:05"`
	UpdatedAt time.Time    `json:"updatedAt" gorm:"format:2006-01-02 15:04:05"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}
