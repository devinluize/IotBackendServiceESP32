package entities

type ArticleBodyEntities struct {
	ArticleBodyId           int    `gorm:"column:article_body_id;primaryKey;not null" json:"article_body_id"`
	ArticleId               int    `gorm:"column:article_id" json:"article_id"`
	ArticleBodyParagraph    string `gorm:"column:article_body_paragraph" json:"article_body_paragraph"`
	ArticleImageContentPath string `gorm:"column:article_image_content_path;size:255" json:"article_image_content_path"`
}

func (*ArticleBodyEntities) TableName() string {
	return "mtr_article_detail_body"
}
