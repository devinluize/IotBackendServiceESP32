package EquipmentServiceImpl

import (
	entities "IotBackend/api/entities/Equipment"
	"IotBackend/api/helper"
	"IotBackend/api/payloads/Equipment"
	"IotBackend/api/payloads/responses"
	menuRepository "IotBackend/api/repositories/Equipment"
	"IotBackend/api/service/EquipmentService"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
)

type EquipmentBookmarkServiceImpl struct {
	repository menuRepository.EquipmentBookmarkRepository
	db         *gorm.DB
	cld        *cloudinary.Cloudinary
}

func (e *EquipmentBookmarkServiceImpl) GetEquipmentBookmarkByUserId(userId int) ([]Equipment.GetBookmarkEquipmentResponse, *responses.ErrorResponses) {
	trans := e.db.Begin()
	res, err := e.repository.GetEquipmentBookmarkByUserId(trans, userId, e.cld)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (e *EquipmentBookmarkServiceImpl) AddEquipmentBookmark(userId, equipmentCourseId int) (entities.EquipmentBookmark, *responses.ErrorResponses) {
	trans := e.db.Begin()
	res, err := e.repository.AddEquipmentBookmark(trans, userId, equipmentCourseId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (e *EquipmentBookmarkServiceImpl) RemoveEquipmentBookmark(userId, equipmentCourseId int) (bool, *responses.ErrorResponses) {
	trans := e.db.Begin()
	res, err := e.repository.RemoveEquipmentBookmark(trans, userId, equipmentCourseId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}

func NewEquipmentBookmarkServiceImpl(repository menuRepository.EquipmentBookmarkRepository, db *gorm.DB, cld *cloudinary.Cloudinary) EquipmentService.EquipmentBookmarkService {
	return &EquipmentBookmarkServiceImpl{repository: repository, db: db, cld: cld}

}
