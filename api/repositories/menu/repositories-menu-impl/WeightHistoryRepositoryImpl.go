package MenuImplRepositories

import (
	"IotBackend/api/entities"
	"IotBackend/api/helper"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	menuRepository "IotBackend/api/repositories/menu"
	"errors"
	"gorm.io/gorm"
	"math"
	"net/http"
)

type WeightHistoryRepositoryImpl struct {
}

func NewWeightHistoryRepositoryImpl() menuRepository.WeightHistoryRepository {
	return &WeightHistoryRepositoryImpl{}

}

func (controller *WeightHistoryRepositoryImpl) GetWeightNotes(db *gorm.DB, UserId int, paginationResponses helper.Pagination) (helper.Pagination, *responses.ErrorResponses) {
	Entities := entities.WeightHistoryEntities{}
	var ScannPl []MenuPayloads.WeightHistoryPayloads
	err := db.Model(&entities.WeightHistoryEntities{}).Where("user_id = ?", UserId).
		Scopes(helper.Paginate(&Entities, &paginationResponses, db)).
		Order("user_weight_time desc").Where("user_id = ?", UserId).
		Scan(&ScannPl).Error
	if err != nil {
		return paginationResponses, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed On Get Pagination Get Weight Notes",
			Err:        err,
			Success:    false,
			Data:       Entities,
		}
	}
	paginationResponses.Rows = ScannPl
	return paginationResponses, nil
}

func (controller *WeightHistoryRepositoryImpl) PostWeightNotes(db *gorm.DB, payloads MenuPayloads.WeightHistoryPayloads, userId int) (entities.WeightHistoryEntities, *responses.ErrorResponses) {
	WeightHistoryEntities := entities.WeightHistoryEntities{
		//UserId:         payloads.UserId,
		UserId:         userId,
		UserWeight:     payloads.UserWeight,
		UserWeightTime: payloads.UserWeightTime,
	}
	profile := entities.UserDetail{}
	err := db.Model(&profile).Where(entities.UserDetail{UserId: userId}).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return WeightHistoryEntities, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Err:        err,
				Message:    "user detail is not found",
			}
		}
		return WeightHistoryEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to get user detail",
			Err:        err,
		}
	}
	if profile.UserHeight != 0 {
		heightMeterSquare := math.Pow(profile.UserHeight/100, 2)
		WeightHistoryEntities.UserBmi = payloads.UserWeight / heightMeterSquare
	}

	err = db.Create(&WeightHistoryEntities).Scan(&WeightHistoryEntities).Error
	if err != nil {
		return WeightHistoryEntities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        nil,
			Success:    false,
		}
	}
	return WeightHistoryEntities, nil
	//123
}

func (controller *WeightHistoryRepositoryImpl) DeleteWeightNotes(db *gorm.DB, UserId int, WeightHistoryId int) (bool, *responses.ErrorResponses) {
	WeightHistoryEntities := entities.WeightHistoryEntities{}

	err := db.Model(&WeightHistoryEntities).Where(entities.WeightHistoryEntities{WeightHistoryId: WeightHistoryId, UserId: UserId}).First(&WeightHistoryEntities).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &responses.ErrorResponses{
				StatusCode: http.StatusBadRequest,
				Message:    "Record Not Found",
				Err:        err,
				Success:    false,
				Data:       nil,
			}
		}
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
			Success:    false,
			Data:       nil}
	}
	err = db.Delete(&WeightHistoryEntities).Error
	if err != nil {
		return false, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError, Message: err.Error(), Err: err}
	}
	return true, nil
}
func (controller *WeightHistoryRepositoryImpl) GetLastWeightHistory(db *gorm.DB, UserId int) (MenuPayloads.LastWeightResponse, *responses.ErrorResponses) {
	var response MenuPayloads.LastWeightResponse
	Entities := entities.WeightHistoryEntities{}
	err := db.Model(&entities.WeightHistoryEntities{}).
		Where(&entities.WeightHistoryEntities{UserId: UserId}).
		Order("user_weight_time DESC").First(&Entities).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, nil
		}
		return response, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
		}
	}
	response = MenuPayloads.LastWeightResponse{
		UserWeightTime: Entities.UserWeightTime,
		UserId:         UserId,
		UserWeight:     Entities.UserWeight,
	}
	return response, nil
}
func (controller *WeightHistoryRepositoryImpl) GetAllWeightWithDateFilter(db *gorm.DB, userId int, dateParams map[string]string) ([]MenuPayloads.WeightHistoryGetAllResponse, *responses.ErrorResponses) {
	if dateParams["date_from"] == "" {
		dateParams["date_from"] = "19000101"
	}
	if dateParams["date_to"] == "" {
		dateParams["date_to"] = "99991212"
	}
	var responseData []MenuPayloads.WeightHistoryGetAllResponse
	strDateFilter := "user_weight_time >='" + dateParams["date_from"] + "' AND user_weight_time <= '" + dateParams["date_to"] + "'"

	err := db.Model(&entities.WeightHistoryEntities{}).
		Where(strDateFilter).
		Where(entities.WeightHistoryEntities{UserId: userId}).
		Order("user_weight_time DESC").
		Scan(&responseData).Error
	if err != nil {
		return responseData, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to get user weight data",
		}
	}
	return responseData, nil

}
