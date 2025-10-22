package menu

import (
	"IotBackend/api/entities"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
)

type CalendarService interface {
	InsertCalendar(payloads MenuPayloads.CalendarInsertPayload) (entities.EventEntity, *responses.ErrorResponses)
	GetCalendarByUserId(userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses)
	UpdateCalendar(payloads MenuPayloads.CalendarUpdatePayload) (entities.EventEntity, *responses.ErrorResponses)
	DeleteCalendarById(calendarId int) (entities.EventEntity, *responses.ErrorResponses)
	GetCalendarByDate(date string, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses)
}
