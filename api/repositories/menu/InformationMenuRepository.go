package menuRepository

import (
	"IotBackend/api/entities"
	"IotBackend/api/helper"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
)

type ArticleMenu interface {
	InsertArticle(tx *gorm.DB, payloads MenuPayloads.ArticleInsertPayloads) (entities.ArticleEntities, *responses.ErrorResponses)
	DeleteArticleById(db *gorm.DB, id int) (bool, *responses.ErrorResponses)
	UpdateArticle(tx *gorm.DB, payloads MenuPayloads.ArticleUpdatePayloads) (entities.ArticleEntities, *responses.ErrorResponses)
	GetAllArticleWithPagination(db *gorm.DB, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses)
	GetArticleById(db *gorm.DB, id int, userId int, cld *cloudinary.Cloudinary) (MenuPayloads.ArticleSelectResponses, *responses.ErrorResponses)
	GetAllArticleWithFilter(db *gorm.DB, paginationResponses helper.Pagination, Key string, userId int, cloudinary *cloudinary.Cloudinary) (helper.Pagination, *responses.ErrorResponses)
	GetArticleHistory(db *gorm.DB, userId int) ([]entities.SearchHistoryEntities, *responses.ErrorResponses)
}
