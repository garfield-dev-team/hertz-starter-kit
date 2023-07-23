package mysql

import "hertz-starter-kit/pkg/config"

type Article struct {
	*BaseModel
	Title      string `json:"title" gorm:"type:varchar(100);not null"`
	Desc       string `json:"desc" gorm:"type:varchar(255);default:''"`
	CoverImage string `json:"cover_image" gorm:"type:varchar(255);default:''"`
	Content    string `json:"content" gorm:"type:longtext"`

	UserID     uint        `json:"user_id" gorm:"type:bigint;not null"`
	Categories []*Category `json:"categories" gorm:"many2many:article_category;"`
}

func (*Article) TableName() string {
	return config.Config.Database.TablePrefix + "article"
}

func CreateArticle(article *Article) (uint, error) {
	if err := DB.Create(article).Error; err != nil {
		return 0, err
	}
	return article.ID, nil
}

// QueryArticleById .
// 查询指定文章，可以外键关联作者、分类等
func QueryArticleById(id uint) (*Article, error) {
	var article *Article

	result := DB.Where("id = ?", id).Limit(1).Find(&article)
	if err := result.Error; err != nil {
		return nil, err
	}
	return article, nil
}

// QueryArticles .
// 分页查询文章列表，可以外键关联作者、分类等
func QueryArticles(pageNum, pageSize int) ([]*Article, error) {
	var articles []*Article
	// 支持分页查询
	offset := (pageNum - 1) * pageSize
	result := DB.Limit(pageSize).Offset(offset).Find(&articles)

	if err := result.Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func UpdateArticle(article *Article) error {
	return DB.Model(&article).Updates(article).Error
}

func DeleteArticleById(id uint) error {
	var article *Article
	// 删除一条记录时，删除对象需要指定主键，否则会触发 批量删除
	// 这里传递 article 的作用是告诉 GORM 操作哪张表，推荐传递指针类型，没有额外内存分配
	return DB.Delete(&article, id).Error
}
