package auth

import (
	"IotBackend/api/entities"
	payloads "IotBackend/api/payloads/auth"
	"IotBackend/api/payloads/responses"
)

type AuthService interface {
	LoginAuth(requestPayloads payloads.LoginPaylods) (responses.ErrorResponses, entities.Users)
	Register(payloads payloads.RegisterPayloads) responses.ErrorResponses
}
