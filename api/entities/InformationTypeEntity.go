package entities

type ArticleType struct {
	ArticleTypeId   int    `gorm:"column:article_type_id;primaryKey;not null" json:"article_type_id"`
	ArticleTypeName string `gorm:"article_type_name" json:"article_type_name"`
}

func (*ArticleType) TableName() string {
	return "mtr_article_type"
}
