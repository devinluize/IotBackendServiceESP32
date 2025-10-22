package menuserviceimpl

import (
	"IotBackend/api/entities"
	"IotBackend/api/helper"
	payloads "IotBackend/api/payloads/auth"
	MenuPayloads "IotBackend/api/payloads/menu"
	"IotBackend/api/payloads/responses"
	menuRepository "IotBackend/api/repositories/menu"
	"IotBackend/api/service/menu"
	"gorm.io/gorm"
)

func NewProfileServiceImpl(db *gorm.DB, repository menuRepository.ProfileMenuRepository) menu.ProfileService {
	return &ProfileServiceImpl{
		db:         db,
		repository: repository,
	}

}

type ProfileServiceImpl struct {
	db         *gorm.DB
	repository menuRepository.ProfileMenuRepository
}

func (service *ProfileServiceImpl) GetProfileMenu(id int) (payloads.GetUserDetailById, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.GetProfileMenu(trans, id)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *ProfileServiceImpl) UpdateProfileMenu(Request MenuPayloads.ProfilePayloadRequest, userId int) (entities.UserDetail, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.UpdateProfileMenu(trans, Request, userId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *ProfileServiceImpl) CreateProfileMenu(Request MenuPayloads.ProfilePayloadRequest, userId int) (entities.UserDetail, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.CreateProfileMenu(trans, Request, userId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (service *ProfileServiceImpl) GetBmi(userId int) (payloads.UserBmiResponse, *responses.ErrorResponses) {
	trans := service.db.Begin()
	res, err := service.repository.GetBmi(trans, userId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
