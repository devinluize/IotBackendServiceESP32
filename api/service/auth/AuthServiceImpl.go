package auth

import (
	"IotBackend/api/entities"
	"IotBackend/api/helper"
	payloads "IotBackend/api/payloads/auth"
	"IotBackend/api/payloads/responses"
	"IotBackend/api/repositories/user"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	DB   *gorm.DB
	Repo userrepositories.UsersRepository
}

func NewAuthServiceImpl(db *gorm.DB, Repo userrepositories.UsersRepository) AuthService {
	return &AuthServiceImpl{
		DB:   db,
		Repo: Repo,
	}
}

//func ConvertToDetailEntities(payloads payloads.RegisterPayloads)entities.UserDetail{
//	return entities.UserDetail{
//		UserDetailId:           0,
//		UserId:                 0,
//		UserWeight:             0,
//		UserHeight:             0,
//		UserGender:             "",
//		UserProfileDescription: "",
//		UserProfileImage:       "",
//		UserPhoneNumber:        "",
//	}
//}
//func ConverToEntities(payloads payloads.RegisterPayloads) entities.Users {
//	return entities.Users{
//		UserName:     payloads.Username,
//		UserEmail:    payloads.Useremail,
//		UserPassword: payloads.Userpasword,
//		UserDetails:
//		//IsVIP:        payloads.IsVIP,
//	}
//}

func ConverToEntitiesLogin(payloads payloads.LoginPaylods) entities.Users {
	return entities.Users{
		//Username:    payloads.Username,
		UserEmail:    payloads.Useremail,
		UserPassword: payloads.Userpasword,
		//IsVIP:       payloads.IsVIP,
	}
}
func (a *AuthServiceImpl) Register(payloads payloads.RegisterPayloads) responses.ErrorResponses {
	tx := a.DB.Begin()
	str := a.Repo.Register(payloads, tx)
	defer helper.CommitOrRollback(tx)
	return str
}
func (a *AuthServiceImpl) LoginAuth(payloads payloads.LoginPaylods) (responses.ErrorResponses, entities.Users) {
	tx := a.DB.Begin()
	str, data := a.Repo.Login(ConverToEntitiesLogin(payloads), tx)
	defer helper.CommitOrRollback(tx)
	return str, data
}
