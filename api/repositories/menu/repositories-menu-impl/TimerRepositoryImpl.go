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

type TimerRepositoryImpl struct {
}

func NewTimerRepositoryImpl() menuRepository.TimerRepository {
	return &TimerRepositoryImpl{}
}

func (repo *TimerRepositoryImpl) InsertTimer(db *gorm.DB, payload MenuPayloads.TimerInsertPayload, userId int) (entities.TimerEntity, *responses.ErrorResponses) {
	TimerEntities := entities.TimerEntity{
		UserId: userId,
		//RemindingHours:   payload.RemindingHours,
		TimerName:        payload.TimerName,
		TimerDescription: payload.TimerDescription,
		//RemindingMinutes: payload.RemindingMinutes,
	}
	err := db.Create(&TimerEntities).First(&TimerEntities).Error
	if err != nil {
		return TimerEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}
	return TimerEntities, nil
}

func (repo *TimerRepositoryImpl) InsertQueueTimer(db *gorm.DB, payload MenuPayloads.TimerQueueInsertResponse) (entities.TimerQueueEntity, *responses.ErrorResponses) {
	timerQueueEntities := entities.TimerQueueEntity{
		TimerId:                    payload.TimerId,
		TimerQueueName:             payload.TimerQueueName,
		TimerQueueRemindingHour:    payload.TimerQueueRemindingHour,
		TimerQueueRemindingMinutes: payload.TimerQueueRemindingMinutes,
		TimerQueueRemindingSecond:  payload.TimerQueueRemindingSecond,
	}
	lastLineNumber := 0
	err := db.Raw(`
				SELECT ISNULL(timer_queue_line_number,0) FROM trx_timer_queue A 
							INNER JOIN	trx_timer B ON A.timer_id = B.timer_id
				WHERE B.timer_id = ?
			`, payload.TimerId).Scan(&lastLineNumber).Error
	if err != nil {
		return timerQueueEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}
	lastLineNumber += 1
	//TimerQueueLineNumber
	timerQueueEntities.TimerQueueLineNumber = lastLineNumber
	err = db.Create(&timerQueueEntities).First(&timerQueueEntities).Error
	if err != nil {
		return timerQueueEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}
	return timerQueueEntities, nil
}
func (repo *TimerRepositoryImpl) UpdateQueueTimer(db *gorm.DB, payload MenuPayloads.TimerQueueUpdatePayload) (entities.TimerQueueEntity, *responses.ErrorResponses) {
	//get first
	timerQueueEntity := entities.TimerQueueEntity{}
	err := db.Model(timerQueueEntity).Where(entities.TimerQueueEntity{TimerQueueId: payload.TimerQueueId}).
		First(&timerQueueEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.TimerQueueEntity{}, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Err:        err,
				Success:    false,
			}
		}
		return timerQueueEntity, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}
	timerQueueEntity.TimerQueueName = payload.TimerQueueName
	timerQueueEntity.TimerQueueRemindingMinutes = payload.TimerQueueRemindingMinutes
	timerQueueEntity.TimerQueueRemindingHour = payload.TimerQueueRemindingHour
	timerQueueEntity.TimerQueueRemindingSecond = payload.TimerQueueRemindingSecond

	err = db.Save(&timerQueueEntity).Error
	if err != nil {
		return timerQueueEntity, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}
	return timerQueueEntity, nil
}
func (repo *TimerRepositoryImpl) DeleteTimerQueueTimer(db *gorm.DB, TimerQueueId int) (bool, *responses.ErrorResponses) {
	err := db.Delete(&entities.TimerQueueEntity{}, entities.TimerQueueEntity{TimerQueueId: TimerQueueId}).Error
	if err != nil {
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}
	return true, nil
}

func (repo *TimerRepositoryImpl) GetTimerByUserId(db *gorm.DB, UserId int) ([]MenuPayloads.GetAllTimerByUserIdResponse, *responses.ErrorResponses) {
	var TimerPayloads []MenuPayloads.GetAllTimerByUserIdResponse
	err := db.Model(&entities.TimerEntity{}).
		Where(entities.TimerEntity{UserId: UserId}).Scan(&TimerPayloads).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return TimerPayloads, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Err:        err,
				Success:    false,
				Data:       err,
			}
		}
		return TimerPayloads, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
		}
	}
	return TimerPayloads, nil
}
func (repo *TimerRepositoryImpl) GetAllQueueTimer(db *gorm.DB, timerId int) ([]entities.TimerQueueEntity, *responses.ErrorResponses) {
	var queueEntities []entities.TimerQueueEntity
	err := db.Model(&entities.TimerQueueEntity{}).Where(entities.TimerQueueEntity{TimerId: timerId}).Scan(&queueEntities).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return queueEntities, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Err:        err,
				Success:    false,
			}
		}
		return queueEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
		}
	}
	return queueEntities, nil
}
func (repo *TimerRepositoryImpl) DeleteTimer(db *gorm.DB, timerId int) (bool, *responses.ErrorResponses) {
	err := db.Delete(&entities.TimerQueueEntity{}, entities.TimerQueueEntity{TimerId: timerId}).Error
	if err != nil {
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
		}
	}
	err = db.Delete(&entities.TimerEntity{}, entities.TimerEntity{TimerId: timerId}).Error
	if err != nil {
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
		}
	}
	return true, nil
}
