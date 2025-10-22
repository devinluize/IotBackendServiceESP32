package MenuPayloads

import "time"

type ArticleBodyDetail struct {
	ArticleBodyParagraph    string `gorm:"column:article_body_paragraph" json:"article_body_paragraph"`
	ArticleImageContentPath string `gorm:"column:article_image_content_path" json:"article_image_content_path"`
}
type ArticleInsertPayloads struct {
	//ArticleId                int                     `json:"article_id"`
	ArticleHeader            string              `json:"article_header"`
	ArticleHeaderPathContent string              `gorm:"column:article_header_path_content" json:"article_header_path_content"`
	ArticleBodyParagraph     []ArticleBodyDetail `json:"article_body_paragraph"`
	//ArticleTypeId            int                     `json:"article_type_id"`
}
type ArticleSelectResponses struct {
	ArticleHeader            string    `json:"article_header"`
	ArticleDateCreated       time.Time `gorm:"column:article_date_created" json:"article_date_created"`
	ArticleCreatedByUserId   int       `gorm:"column:article_created_by_user_id" json:"article_created_by_user_id"`
	ArticleId                int       `gorm:"column:article_id;not null;primaryKey" json:"article_id"`
	ArticleHeaderPathContent string    `gorm:"column:article_header_path_content" json:"article_header_image_path"`

	ArticleBodyContent []ArticleBodyDetail `json:"article_body_content"`
	IsBookmark         bool                `json:"is_bookmark"`
	//ArticleTypeId          int                     `json:"article_type_id"`
}
type ArticleSelectResponseHeader struct {
	ArticleHeader          string    `json:"article_header"`
	ArticleDateCreated     time.Time `gorm:"column:article_date_created" json:"article_date_created"`
	ArticleCreatedByUserId int       `gorm:"column:article_created_by_user_id" json:"article_created_by_user_id"`
	ArticleId              int       `gorm:"column:article_id;not null;primaryKey" json:"article_id"`
	BookmarkId             int       `json:"bookmark_id"`
}

type ArticleUpdatePayloads struct {
	ArticleId          int                 `json:"article_id"`
	ArticleBodyContent []ArticleBodyDetail `json:"article_body_content"`
}
