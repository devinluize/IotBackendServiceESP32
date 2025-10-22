package menuRepository

import (
	"IotBackend/api/entities"
	payloads "IotBackend/api/payloads/auth"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	"gorm.io/gorm"
)

type ProfileMenuRepository interface {
	GetProfileMenu(db *gorm.DB, id int) (payloads.GetUserDetailById, *responses.ErrorResponses)
	UpdateProfileMenu(db *gorm.DB, Request MenuPayloads.ProfilePayloadRequest, userId int) (entities.UserDetail, *responses.ErrorResponses)
	CreateProfileMenu(db *gorm.DB, Request MenuPayloads.ProfilePayloadRequest, userId int) (entities.UserDetail, *responses.ErrorResponses)
	GetBmi(db *gorm.DB, userId int) (payloads.UserBmiResponse, *responses.ErrorResponses)
}
