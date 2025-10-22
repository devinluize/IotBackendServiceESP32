package MenuImplRepositories

import (
	"IotBackend/api/entities"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	menuRepository "IotBackend/api/repositories/menu"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type CalendarRepositoryImpl struct {
}

func NewEventRepositoryImpl() menuRepository.EventRepository {
	return &CalendarRepositoryImpl{}
}
func (repo *CalendarRepositoryImpl) InsertEvent(db *gorm.DB, payloads MenuPayloads.CalendarInsertPayload) (entities.EventEntity, *responses.ErrorResponses) {
	CalendarEntities := entities.EventEntity{
		EventName:     payloads.EventName,
		EventDate:     payloads.EventDate,
		UserId:        payloads.UserId,
		EventTimeFrom: payloads.EventTimeFrom,
		EventTimeTo:   payloads.EventTimeTo,
	}
	err := db.Create(&CalendarEntities).Error
	if err != nil {
		return CalendarEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to create calendar entity",
			Success:    false,
			Err:        err,
		}
	}
	return CalendarEntities, nil
}
func (repo *CalendarRepositoryImpl) GetEventByUserId(db *gorm.DB, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses) {
	CalendarEntities := entities.EventEntity{}
	var ResponseGetById []MenuPayloads.CalendarGetByIdResponse
	err := db.Model(&CalendarEntities).Where(entities.EventEntity{UserId: userId}).Scan(&ResponseGetById).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ResponseGetById, &responses.ErrorResponses{}
		}
		return ResponseGetById, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to get calendar by user id",
			Success:    false,
			Err:        err,
		}
	}
	return ResponseGetById, nil
}
func (repo *CalendarRepositoryImpl) UpdateEvent(db *gorm.DB, payloads MenuPayloads.CalendarUpdatePayload) (entities.EventEntity, *responses.ErrorResponses) {
	CalendarEntities := entities.EventEntity{}
	err := db.Model(&CalendarEntities).Where(entities.EventEntity{EventId: payloads.EventId}).
		First(&CalendarEntities).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return CalendarEntities, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Message:    "calendar not found",
				Err:        err,
				Success:    false,
			}
		}
		return CalendarEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to update calendar",
			Success:    false,
			Err:        err,
		}
	}
	CalendarEntities.EventTimeTo = payloads.EventTimeTo
	CalendarEntities.EventTimeFrom = payloads.EventTimeFrom
	CalendarEntities.EventName = payloads.EventName
	err = db.Save(&CalendarEntities).Error
	if err != nil {
		return CalendarEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to update calendar",
			Success:    false,
			Err:        err,
		}
	}
	return CalendarEntities, nil
}
func (repo *CalendarRepositoryImpl) DeleteEventById(db *gorm.DB, calendarId int) (entities.EventEntity, *responses.ErrorResponses) {
	var CalendarEntities entities.EventEntity
	//get cek first if there is the data
	err := db.Model(&CalendarEntities).Where(entities.EventEntity{EventId: calendarId}).First(&CalendarEntities).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return CalendarEntities, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Message:    "calendar to delete is not found",
				Err:        err,
				Success:    false,
			}
		}
		return CalendarEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to delete calendar",
			Success:    false,
			Err:        err,
		}
	}
	err = db.Delete(&CalendarEntities).Error
	if err != nil {
		return CalendarEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to delete calendar",
			Success:    false,
			Err:        err,
		}
	}
	return CalendarEntities, nil
}
func (repo *CalendarRepositoryImpl) GetEventById(db *gorm.DB, date string, userId int) ([]MenuPayloads.CalendarGetByIdResponse, *responses.ErrorResponses) {
	calendarEntities := entities.EventEntity{}
	var calendarResponse []MenuPayloads.CalendarGetByIdResponse
	err := db.Model(&calendarEntities).
		Where("CONVERT(DATE, event_date) = ?", date).
		Where(entities.EventEntity{UserId: userId}).
		Order("event_time_from ASC").
		Scan(&calendarResponse).Error
	if err != nil {
		return calendarResponse, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to get calendar by date",
		}
	}
	return calendarResponse, nil
}
