package menuRepository

import (
	"IotBackend/api/entities"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	"gorm.io/gorm"
)

type TimerRepository interface {
	InsertTimer(db *gorm.DB, payload MenuPayloads.TimerInsertPayload, userId int) (entities.TimerEntity, *responses.ErrorResponses)
	InsertQueueTimer(db *gorm.DB, payload MenuPayloads.TimerQueueInsertResponse) (entities.TimerQueueEntity, *responses.ErrorResponses)
	UpdateQueueTimer(db *gorm.DB, payload MenuPayloads.TimerQueueUpdatePayload) (entities.TimerQueueEntity, *responses.ErrorResponses)
	DeleteTimerQueueTimer(db *gorm.DB, TimerQueueId int) (bool, *responses.ErrorResponses)
	GetTimerByUserId(db *gorm.DB, UserId int) ([]MenuPayloads.GetAllTimerByUserIdResponse, *responses.ErrorResponses)
	GetAllQueueTimer(db *gorm.DB, TimerId int) ([]entities.TimerQueueEntity, *responses.ErrorResponses)
	DeleteTimer(db *gorm.DB, timerId int) (bool, *responses.ErrorResponses)
}
