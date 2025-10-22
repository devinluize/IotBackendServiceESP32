package menuserviceimpl

import (
	"IotBackend/api/entities"
	"IotBackend/api/helper"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	menuRepository "IotBackend/api/repositories/menu"
	"IotBackend/api/service/menu"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
)

type ArticleServiceImpl struct {
	repo menuRepository.ArticleMenu
	db   *gorm.DB
	cld  *cloudinary.Cloudinary
}

func NewArticleServiceImpl(repo menuRepository.ArticleMenu, db *gorm.DB, cld *cloudinary.Cloudinary) menu.ArticleService {
	return &ArticleServiceImpl{repo: repo, db: db, cld: cld}
}
func (service *ArticleServiceImpl) InsertArticle(payloads MenuPayloads.ArticleInsertPayloads) (entities.ArticleEntities, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.InsertArticle(trans, payloads)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *ArticleServiceImpl) DeleteArticleById(id int) (bool, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.DeleteArticleById(trans, id)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *ArticleServiceImpl) UpdateArticle(payloads MenuPayloads.ArticleUpdatePayloads) (entities.ArticleEntities, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.UpdateArticle(trans, payloads)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *ArticleServiceImpl) GetArticleById(id int, userId int) (MenuPayloads.ArticleSelectResponses, *responses.ErrorResponses) {
	trans := service.db.Begin()

	res, err := service.repo.GetArticleById(trans, id, userId, service.cld)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (service *ArticleServiceImpl) GetAllArticleWithPagination(paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.GetAllArticleWithPagination(trans, paginationResponses)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (service *ArticleServiceImpl) GetAllArticleWithFilter(paginationResponses helper.Pagination, Key string, userId int) (helper.Pagination, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.GetAllArticleWithFilter(trans, paginationResponses, Key, userId, service.cld)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (service *ArticleServiceImpl) GetArticleHistory(userId int) ([]entities.SearchHistoryEntities, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repo.GetArticleHistory(trans, userId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
