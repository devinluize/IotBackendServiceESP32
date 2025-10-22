package menu

import (
	"IotBackend/api/entities"
	payloads "IotBackend/api/payloads/auth"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
)

type ProfileService interface {
	GetProfileMenu(id int) (payloads.GetUserDetailById, *responses.ErrorResponses)
	UpdateProfileMenu(Request MenuPayloads.ProfilePayloadRequest, userId int) (entities.UserDetail, *responses.ErrorResponses)
	CreateProfileMenu(Request MenuPayloads.ProfilePayloadRequest, userId int) (entities.UserDetail, *responses.ErrorResponses)
	GetBmi(userId int) (payloads.UserBmiResponse, *responses.ErrorResponses)
}
