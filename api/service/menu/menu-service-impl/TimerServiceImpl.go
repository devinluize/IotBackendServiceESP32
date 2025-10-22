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

type TimerServiceImpl struct {
	repository menuRepository.TimerRepository
	DB         *gorm.DB
}

func (t *TimerServiceImpl) DeleteTimer(TimerId int) (bool, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.DeleteTimer(db, TimerId)
	if err != nil {
		return res, err
	}
	return res, nil
}

func NewTimerServiceImpl(repository menuRepository.TimerRepository, db *gorm.DB) menu.TimerService {
	return &TimerServiceImpl{DB: db, repository: repository}
}

func (t *TimerServiceImpl) InsertTimer(payload MenuPayloads.TimerInsertPayload, userId int) (entities.TimerEntity, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.InsertTimer(db, payload, userId)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (t *TimerServiceImpl) InsertQueueTimer(payload MenuPayloads.TimerQueueInsertResponse) (entities.TimerQueueEntity, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.InsertQueueTimer(db, payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (t *TimerServiceImpl) UpdateQueueTimer(payload MenuPayloads.TimerQueueUpdatePayload) (entities.TimerQueueEntity, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.UpdateQueueTimer(db, payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (t *TimerServiceImpl) DeleteTimerQueueTimer(TimerQueueId int) (bool, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.DeleteTimerQueueTimer(db, TimerQueueId)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (t *TimerServiceImpl) GetTimerByUserId(timerId int) ([]MenuPayloads.GetAllTimerByUserIdResponse, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.GetTimerByUserId(db, timerId)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (t *TimerServiceImpl) GetAllQueueTimer(timerQueueId int) ([]entities.TimerQueueEntity, *responses.ErrorResponses) {
	db := t.DB.Begin()
	defer helper.CommitOrRollback(db)
	res, err := t.repository.GetAllQueueTimer(db, timerQueueId)
	if err != nil {
		return res, err
	}
	return res, nil
}
