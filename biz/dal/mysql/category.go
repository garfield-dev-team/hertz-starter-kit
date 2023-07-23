package mysql

import "hertz-starter-kit/pkg/config"

type Category struct {
	*BaseModel
	Name  string `json:"name" gorm:"type:varchar(100);not null"`
	State uint8  `json:"state" gorm:"type:tinyint(3);default:0"` // 状态 0 为禁用、1 为启用

	Articles []*Article `json:"articles" gorm:"many2many:article_category;"`
}

func (*Category) TableName() string {
	return config.Config.Database.TablePrefix + "category"
}

func CreateCategory(category *Category) (uint, error) {
	if err := DB.Create(&category).Error; err != nil {
		return 0, err
	}
	return category.ID, nil
}

// QueryCategoryById .
// 查询某个分类下关联的文章
func QueryCategoryById(id uint) (*Category, error) {
	var category *Category
	result := DB.Model(&category).Preload("Articles").Where("id = ?", id).Limit(1).Find(&category)
	if err := result.Error; err != nil {
		return nil, err
	}
	return category, nil
}

// QueryCategories .
// 查询全部分类列表
func QueryCategories() ([]*Category, error) {
	var categories []*Category
	if err := DB.Model(&Category{}).Preload("Articles").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func UpdateCategory(category *Category) error {
	return DB.Model(&category).Updates(category).Error
}

func DeleteCategoryById(id uint) error {
	var category *Category
	return DB.Delete(&category, id).Error
}
