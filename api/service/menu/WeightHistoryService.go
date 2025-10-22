package menu

import (
	"IotBackend/api/entities"
	"IotBackend/api/helper"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
)

type WeightHistoryService interface {
	GetWeightNotes(UserId int, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses)
	PostWeightNotes(payloads MenuPayloads.WeightHistoryPayloads, userId int) (entities.WeightHistoryEntities, *responses.ErrorResponses)
	DeleteWeightNotes(UserId int, WeightHistoryId int) (bool, *responses.ErrorResponses)
	GetLastWeightHistory(UserId int) (MenuPayloads.LastWeightResponse, *responses.ErrorResponses)
	GetAllWeightWithDateFilter(userId int, dateParams map[string]string) ([]MenuPayloads.WeightHistoryGetAllResponse, *responses.ErrorResponses)
}
