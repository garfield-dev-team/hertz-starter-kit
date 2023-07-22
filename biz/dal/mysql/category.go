package mysql

type Category struct {
	BaseModel
	Name  string `json:"name" gorm:"type:varchar(100);not null"`
	State uint8  `json:"state" gorm:"type:tinyint(3);default:0"` // 状态 0 为禁用、1 为启用

	Articles []Article `json:"articles" gorm:"many2many:article_category;"`
}
