package auth

import (
	configenv "IotBackend/api/config"
	"IotBackend/api/controller/auth"
	"IotBackend/api/helper"
	payloads "IotBackend/api/payloads/auth"
	"IotBackend/api/payloads/responses"
	auth2 "IotBackend/api/service/auth"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type AuthControllerImpl struct {
	AuthService auth2.AuthService
}

func NewAuthController(AuthService auth2.AuthService) auth.AuthController {
	return &AuthControllerImpl{AuthService: AuthService}
}

// Register List Via Header
//
//	@Security		BearerAuth
//	@Summary		register by heade
//	@Description	register by heade
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		payloads.RegisterPayloads	true	"Insert Header Request"
//	@Success		200		{object}	 responses.ErrorResponses
//	@Router			/api/user/register [post]
func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request) {
	payloadsData := payloads.RegisterPayloads{}
	helper.ReadFromRequestBody(request, &payloadsData)

	data := controller.AuthService.Register(payloadsData)
	if data.Success == false {
		helper.ReturnError(writer, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        data.Err,
			Message:    data.Message,
		})
	}
	loginPayloads := payloads.LoginPaylods{}
	loginPayloads.Useremail = payloadsData.Useremail
	loginPayloads.Userpasword = payloadsData.Userpasword
	loginReq, datas := controller.AuthService.LoginAuth(loginPayloads)

	if loginReq.Success == false {
		helper.WriteToResponseBody(writer, loginReq)
		return
	}
	expTime := time.Now().Add(time.Hour * 1000)

	claims := configenv.JWTClaim{
		UserName:  datas.UserName,
		UserEmail: datas.UserEmail,
		UserId:    datas.UserId,
		//IsVIP:    data.IsVIP,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "devin",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, erronToken := tokenAlgo.SignedString(configenv.JWT_KEY)
	if erronToken != nil {
		helper.WriteToResponseBody(writer, responses.ApiResponseError{
			Message: "parse failed",
			Success: false,
			Err:     nil,
		})
	}
	helper.WriteToResponseBody(writer, payloads.LoginRespons{
		UserName:  datas.UserName,
		UserEmail: datas.UserEmail,
		//IsVIP:     data.IsVIP,
		Token: token,
	})
	return
	//payloads.HandleSuccessStandarRespons(writer, data, "", http.StatusOK)
	//helper.WriteToResponseBody(writer, data)
	//err := validation.ValidationForm(writer, request, &payloadsData)

}

// AuthLogin List Via Header
//
//	@Security		BearerAuth
//	@Summary		Login by header
//	@Description	Login by header
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		payloads.LoginPaylods	true	"Insert Header Request"
//	@Success		200		{object}	 responses.ErrorResponses
//	@Router			/api/user/login2 [post]
func (controller *AuthControllerImpl) AuthLogin(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	loginPayloads := payloads.LoginPaylods{}
	helper.ReadFromRequestBody(request, &loginPayloads)
	loginReq, data := controller.AuthService.LoginAuth(loginPayloads)

	if loginReq.Success == false {
		helper.WriteToResponseBody(writer, loginReq)
		return
	}
	expTime := time.Now().Add(time.Hour * 1000)

	claims := configenv.JWTClaim{
		UserName:  data.UserName,
		UserEmail: data.UserEmail,
		UserId:    data.UserId,
		//IsVIP:    data.IsVIP,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "devin",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, erronToken := tokenAlgo.SignedString(configenv.JWT_KEY)
	if erronToken != nil {
		helper.WriteToResponseBody(writer, responses.ApiResponseError{
			Message: "parse failed",
			Success: false,
			Err:     nil,
		})
	}
	helper.WriteToResponseBody(writer, payloads.LoginRespons{
		UserName:  data.UserName,
		UserEmail: data.UserEmail,
		//IsVIP:     data.IsVIP,
		Token: token,
	})
	return
	//panic("implement me")
}
