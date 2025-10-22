package MenuImplRepositories

import (
	"IotBackend/api/entities"
	payloads "IotBackend/api/payloads/auth"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	menuRepository "IotBackend/api/repositories/menu"
	"errors"
	"gorm.io/gorm"
	"math"
	"net/http"
)

type ProfileMenuRepositoryImpl struct {
}

func NewProfileMenuRepositoryImpl() menuRepository.ProfileMenuRepository {
	return &ProfileMenuRepositoryImpl{}

}
func (p *ProfileMenuRepositoryImpl) GetProfileMenu(db *gorm.DB, id int) (payloads.GetUserDetailById, *responses.ErrorResponses) {
	Entities := payloads.GetUserDetailById{}
	err := db.Table("user_details A").
		Joins("INNER JOIN mtr_user B ON A.user_id = B.user_id").
		Where(entities.Users{UserId: id}).
		Select("A.*,B.*").
		Order("A.user_detail_id DESC").
		Scan(&Entities).Error
	if err != nil {
		return Entities,
			&responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
				Err:     err,
				Message: err.Error()}
	}
	EntitiesWeight := entities.WeightHistoryEntities{}
	err = db.Model(&entities.WeightHistoryEntities{}).
		Where(&entities.WeightHistoryEntities{UserId: id}).
		Order("user_weight_time DESC").First(&EntitiesWeight).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Entities, nil
		}
		return Entities, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Err:        err,
		}
	}
	Entities.UserWeight = EntitiesWeight.UserWeight

	return Entities, nil
}

func (p *ProfileMenuRepositoryImpl) UpdateProfileMenu(db *gorm.DB, Request MenuPayloads.ProfilePayloadRequest, userId int) (entities.UserDetail, *responses.ErrorResponses) {
	Entities := entities.UserDetail{}
	err := db.Model(&Entities).Where(entities.UserDetail{UserId: userId}).
		Order("user_detail_id DESC").
		Scan(&Entities).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
			Err:     err,
			Message: err.Error()}
	}
	if Request.UserHeight != 0 {
		Entities.UserHeight = Request.UserHeight
	}
	if Request.UserWeight != 0 {
		Entities.UserWeight = Request.UserWeight
	}
	if Request.UserGender != "" {
		Entities.UserGender = Request.UserGender
	}
	if Request.UserProfileDescription != "" {
		Entities.UserProfileDescription = Request.UserProfileDescription
	}
	//if Request.UserProfileImage != "" {
	//	Entities.UserProfileImage = Request.UserProfileImage
	//}
	if Request.UserPhoneNumber != "" {
		Entities.UserPhoneNumber = Request.UserPhoneNumber
	}
	entitiesUser := entities.Users{}
	err = db.Model(&entitiesUser).Where(entities.Users{UserId: userId}).First(&entitiesUser).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
			Err:     err,
			Message: err.Error()}
	}
	if Request.UserEmail != "" {
		entitiesUser.UserEmail = Request.UserEmail
	}
	if Request.UserName != "" {
		entitiesUser.UserName = Request.UserName
	}
	err = db.Save(&entitiesUser).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
			Message: "Error updating profile menu",
			Err:     err,
		}
	}
	err = db.Save(&Entities).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{StatusCode: http.StatusInternalServerError,
			Message: "Error updating profile menu",
			Err:     err,
		}
	}
	return Entities, nil
}
func (p *ProfileMenuRepositoryImpl) CreateProfileMenu(db *gorm.DB, Request MenuPayloads.ProfilePayloadRequest, userId int) (entities.UserDetail, *responses.ErrorResponses) {
	Entities := entities.UserDetail{
		UserId:                 userId,
		UserWeight:             Request.UserWeight,
		UserHeight:             Request.UserHeight,
		UserGender:             Request.UserGender,
		UserProfileDescription: Request.UserProfileDescription,
		//UserProfileImage:       Request.UserProfileImage,
		UserPhoneNumber: Request.UserPhoneNumber,
	}
	err := db.Create(&Entities).Error
	if err != nil {
		return Entities, &responses.ErrorResponses{StatusCode: http.StatusBadRequest,
			Message: "Error creating profile menu",
			Err:     err,
		}
	}
	return Entities, nil
}
func (p *ProfileMenuRepositoryImpl) GetBmi(db *gorm.DB, userId int) (payloads.UserBmiResponse, *responses.ErrorResponses) {
	//get user weight first with user id
	userBmiResponse := payloads.UserBmiResponse{}
	var lastWeight float64
	err := db.Model(&entities.WeightHistoryEntities{}).
		Where(entities.WeightHistoryEntities{UserId: userId}).
		Select("user_weight").Limit(1).Order("user_weight_time desc").
		Scan(&lastWeight).Error
	if err != nil {
		return userBmiResponse, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to get user weight by user id",
		}
	}
	//get profile
	profile := entities.UserDetail{}
	err = db.Model(&profile).Where(entities.UserDetail{UserId: userId}).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userBmiResponse, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Err:        err,
				Message:    "user detail is not found",
			}
		}
		return userBmiResponse, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to get user detail",
			Err:        err,
		}
	}
	userHeightMeter := profile.UserHeight / 100

	Bmi := lastWeight / (math.Pow(userHeightMeter, 2))
	userBmiResponse.Bmi = Bmi
	userBmiResponse.UserId = userId
	userBmiResponse.UserHeight = userHeightMeter
	userBmiResponse.UserWeight = lastWeight
	return userBmiResponse, nil
}
