package MenuPayloads

import "time"

type GetAllBookmarkResponse struct {
	ArticleHeader            string    `json:"article_header"`
	ArticleDateCreated       time.Time `gorm:"column:article_date_created" json:"article_date_created"`
	ArticleCreatedByUserId   int       `gorm:"column:article_created_by_user_id" json:"article_created_by_user_id"`
	ArticleId                int       `gorm:"column:article_id;not null;primaryKey" json:"article_id"`
	BookmarkId               int       `json:"bookmark_id"`
	ArticleHeaderPathContent string    `gorm:"column:article_header_path_content" json:"article_header_path_content"`
}
