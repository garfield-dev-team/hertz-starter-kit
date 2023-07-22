package mysql

type Article struct {
	BaseModel
	Title      string `json:"title" gorm:"type:varchar(100);not null"`
	Desc       string `json:"desc" gorm:"type:varchar(255);default:''"`
	CoverImage string `json:"coverImage" gorm:"type:varchar(255);default:''"`
	Content    string `json:"content" gorm:"type:longtext;default:''"`

	Author     User       `json:"author"`
	Categories []Category `json:"categories" gorm:"many2many:article_category;"`
}
