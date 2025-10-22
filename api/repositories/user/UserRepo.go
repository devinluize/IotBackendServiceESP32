package userrepositories

import (
	"IotBackend/api/entities"
	payloads "IotBackend/api/payloads/auth"
	"IotBackend/api/payloads/responses"
	"gorm.io/gorm"
)

type UsersRepository interface {
	Register(payloads payloads.RegisterPayloads, DB *gorm.DB) responses.ErrorResponses
	Login(requestData entities.Users, DB *gorm.DB) (responses.ErrorResponses, entities.Users)
	//InsertProfile(	payloads)
}
