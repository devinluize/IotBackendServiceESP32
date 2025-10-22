package menuRepository

import (
	"IotBackend/api/entities"
	"IotBackend/api/helper"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	"gorm.io/gorm"
)

type WeightHistoryRepository interface {
	GetWeightNotes(db *gorm.DB, UserId int, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses)
	PostWeightNotes(db *gorm.DB, payloads MenuPayloads.WeightHistoryPayloads, userId int) (entities.WeightHistoryEntities, *responses.ErrorResponses)
	DeleteWeightNotes(db *gorm.DB, UserId int, WeightHistoryId int) (bool, *responses.ErrorResponses)
	GetLastWeightHistory(db *gorm.DB, UserId int) (MenuPayloads.LastWeightResponse, *responses.ErrorResponses)
	GetAllWeightWithDateFilter(db *gorm.DB, userId int, dateParams map[string]string) ([]MenuPayloads.WeightHistoryGetAllResponse, *responses.ErrorResponses)
}
