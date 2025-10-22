package entities

import "time"

type ArticleEntities struct {
	ArticleId                int                   `gorm:"column:article_id;not null;primaryKey" json:"article_id"`
	ArticleHeader            string                `gorm:"column:article_header;size:255" json:"article_header"`
	ArticleHeaderPathContent string                `gorm:"column:article_header_path_content;size:255" json:"article_header_path_content"`
	ArticleDateCreated       time.Time             `gorm:"column:article_date_created" json:"article_date_created"`
	ArticleCreatedByUserId   int                   `gorm:"column:article_created_by_user_id" json:"article_created_by_user_id"`
	ArticleBody              []ArticleBodyEntities `gorm:"foreignKey:ArticleId;references:ArticleId" json:"article_body"`
	//ArticleTypeId            int                       `gorm:"column:article_type_id" json:"article_type_id"`
	//ArticleType              ArticleType
}

func (*ArticleEntities) TableName() string {
	return "mtr_article"
}
