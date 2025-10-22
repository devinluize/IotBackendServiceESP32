package menu

import (
	"IotBackend/api/entities"
	"IotBackend/api/helper"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
)

type ArticleService interface {
	InsertArticle(payloads MenuPayloads.ArticleInsertPayloads) (entities.ArticleEntities, *responses.ErrorResponses)
	DeleteArticleById(id int) (bool, *responses.ErrorResponses)
	UpdateArticle(payloads MenuPayloads.ArticleUpdatePayloads) (entities.ArticleEntities, *responses.ErrorResponses)
	GetArticleById(id int, userId int) (MenuPayloads.ArticleSelectResponses, *responses.ErrorResponses)
	GetAllArticleWithPagination(paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses)
	GetArticleHistory(userId int) ([]entities.SearchHistoryEntities, *responses.ErrorResponses)
	GetAllArticleWithFilter(paginationResponses helper.Pagination, Key string, userId int) (helper.Pagination, *responses.ErrorResponses)
}
