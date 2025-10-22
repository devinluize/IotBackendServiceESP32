package menu

import (
	"IotBackend/api/entities"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
)

type TimerService interface {
	InsertTimer(payload MenuPayloads.TimerInsertPayload, userId int) (entities.TimerEntity, *responses.ErrorResponses)
	InsertQueueTimer(payload MenuPayloads.TimerQueueInsertResponse) (entities.TimerQueueEntity, *responses.ErrorResponses)
	UpdateQueueTimer(payload MenuPayloads.TimerQueueUpdatePayload) (entities.TimerQueueEntity, *responses.ErrorResponses)
	DeleteTimerQueueTimer(TimerQueueId int) (bool, *responses.ErrorResponses)
	GetTimerByUserId(UserId int) ([]MenuPayloads.GetAllTimerByUserIdResponse, *responses.ErrorResponses)
	GetAllQueueTimer(TimerId int) ([]entities.TimerQueueEntity, *responses.ErrorResponses)
	DeleteTimer(TimerId int) (bool, *responses.ErrorResponses)
}
