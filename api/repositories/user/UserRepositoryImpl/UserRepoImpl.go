package UserRepositoryImpl

import (
	"IotBackend/api/entities"
	payloads "IotBackend/api/payloads/auth"
	"IotBackend/api/payloads/responses"
	"IotBackend/api/repositories/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"runtime"
	"strconv"
)

type AuthRepoImpl struct{}

func NewAuthRepoImpl() userrepositories.UsersRepository {
	return &AuthRepoImpl{}
}

func (a *AuthRepoImpl) Register(requestData payloads.RegisterPayloads, DB *gorm.DB) responses.ErrorResponses {
	//TODO implement me
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(requestData.Userpasword), bcrypt.DefaultCost)
	requestData.Userpasword = string(hashPassword)
	//err2 := DB.AutoMigrate(&requestData)
	//if err2 != nil {
	//	fmt.Println("Failed to auto-migrate:", err2)
	//}
	isDuplicate := 0
	errorUnique := DB.Model(&entities.Users{}).Where(entities.Users{UserEmail: requestData.Useremail}).
		Select("1").Scan(&isDuplicate).Error
	if errorUnique != nil {
		return responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "error when checking duplicate email",
			Err:        errorUnique,
			Success:    false,
			//Data:       nil,
		}
	}
	if isDuplicate == 1 {
		return responses.ErrorResponses{Success: false,
			Err:        errorUnique,
			StatusCode: http.StatusBadRequest, Message: "email is already exist"}
	}
	entitiesUser := entities.Users{
		//UserId:       0,
		UserName:     requestData.Username,
		UserEmail:    requestData.Useremail,
		UserPassword: requestData.Userpasword,
		//UserDetail:   entities.UserDetail{},
	}
	err := DB.Create(&entitiesUser).Scan(&entitiesUser)
	if err.Error != nil {
		return responses.ErrorResponses{
			Success: false,
			Message: err.Error.Error(),
			Data:    nil,
		}
	}
	detailEntities := entities.UserDetail{
		UserDetailId:           0,
		UserId:                 entitiesUser.UserId,
		UserWeight:             requestData.UserWeight,
		UserHeight:             requestData.UserHeight,
		UserGender:             requestData.UserGender,
		UserProfileDescription: "",
		//UserProfileImage:       "",
		UserPhoneNumber: requestData.UserPhoneNumber,
	}
	errs := DB.Create(&detailEntities).Scan(&detailEntities).Error
	if errs != nil {
		_, file, line, _ := runtime.Caller(1)
		return responses.ErrorResponses{
			Success: false,
			Message: "Failed To Create in file and line " + file + strconv.Itoa(line),
			Data:    requestData,
		}
	}
	err = DB.Model(&entitiesUser).Where(entities.Users{UserId: entitiesUser.UserId}).Scan(&entitiesUser)
	return responses.ErrorResponses{
		Success: true,
		Message: "Register Success",
		Data:    entitiesUser,
	}

}

func (a *AuthRepoImpl) Login(requestData entities.Users, DB *gorm.DB) (responses.ErrorResponses, entities.Users) {
	var user entities.Users
	//err := DB.Where("UserName = ?", requestData.Username).First(&user).Error
	////err := DB.Debug().Where("UserEmail = ?", requestData.Useremail).
	//Offset(0).
	//	Limit(1).First(&user).Error
	data := DB.Raw("SELECT TOP 1 * FROM mtr_user A WHERE A.user_email = ?", requestData.UserEmail).Scan(&user).Error
	if data != nil {
		return responses.ErrorResponses{
			Success: false,
			Message: data.Error(),
			Data:    nil,
		}, user
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(requestData.UserPassword))
	if err != nil {
		return responses.ErrorResponses{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}, user
	}
	return responses.ErrorResponses{
		Success: true,
		Message: "success Login",
		Data:    user,
	}, user
}
