package menuserviceimpl

import (
	"IotBackend/api/entities"
	"IotBackend/api/helper"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	menuRepository "IotBackend/api/repositories/menu"
	"IotBackend/api/service/menu"
	"gorm.io/gorm"
)

type WeightHistoryServiceImpl struct {
	db   *gorm.DB
	repo menuRepository.WeightHistoryRepository
}

func NewWeightHistoryServiceImpl(db *gorm.DB, repo menuRepository.WeightHistoryRepository) menu.WeightHistoryService {
	return &WeightHistoryServiceImpl{db: db, repo: repo}

}
func (service *WeightHistoryServiceImpl) GetWeightNotes(UserId int, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses) {
	tx := service.db.Begin()
	res, err := service.repo.GetWeightNotes(tx, UserId, paginationResponses)
	defer helper.CommitOrRollback(tx)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *WeightHistoryServiceImpl) PostWeightNotes(payloads MenuPayloads.WeightHistoryPayloads, userId int) (entities.WeightHistoryEntities, *responses.ErrorResponses) {
	tx := service.db.Begin()
	res, err := service.repo.PostWeightNotes(tx, payloads, userId)
	defer helper.CommitOrRollback(tx)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *WeightHistoryServiceImpl) DeleteWeightNotes(UserId int, WeightHistoryId int) (bool, *responses.ErrorResponses) {
	tx := service.db.Begin()
	res, err := service.repo.DeleteWeightNotes(tx, UserId, WeightHistoryId)
	defer helper.CommitOrRollback(tx)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (service *WeightHistoryServiceImpl) GetLastWeightHistory(UserId int) (MenuPayloads.LastWeightResponse, *responses.ErrorResponses) {
	tx := service.db.Begin()
	res, err := service.repo.GetLastWeightHistory(tx, UserId)
	defer helper.CommitOrRollback(tx)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (service *WeightHistoryServiceImpl) GetAllWeightWithDateFilter(userId int, dateParams map[string]string) ([]MenuPayloads.WeightHistoryGetAllResponse, *responses.ErrorResponses) {
	tx := service.db.Begin()
	res, err := service.repo.GetAllWeightWithDateFilter(tx, userId, dateParams)
	defer helper.CommitOrRollback(tx)
	if err != nil {
		return res, err
	}
	return res, nil
}
