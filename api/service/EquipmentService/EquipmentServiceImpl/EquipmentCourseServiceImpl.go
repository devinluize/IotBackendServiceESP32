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

type EquipmentCourseServiceImpl struct {
	repository menuRepository.EquipmentCourseRepository
	db         *gorm.DB
	cld        *cloudinary.Cloudinary
}

func NewEquipmentCourseServiceImpl(db *gorm.DB, repository menuRepository.EquipmentCourseRepository, cld *cloudinary.Cloudinary) EquipmentService.EquipmentCourseService {
	return &EquipmentCourseServiceImpl{
		db:         db,
		repository: repository,
		cld:        cld,
	}
}
func (s *EquipmentCourseServiceImpl) GetAllEquipmentCourseByEquipment(equipmentId int) (Equipment.GetAllCourseEquipmentResponse, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.GetAllEquipmentCourseByEquipment(trans, equipmentId, s.cld)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *EquipmentCourseServiceImpl) InsertEquipmentCourse(payload Equipment.InsertEquipmentCourseDataPayload) (entities.EquipmentCourseDataEntity, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.InsertEquipmentCourse(trans, payload)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *EquipmentCourseServiceImpl) GetEquipmentCourse(courseId int, userId int) (Equipment.GetCourseByIdResponse, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.GetEquipmentCourse(trans, courseId, s.cld, userId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *EquipmentCourseServiceImpl) SearchEquipmentByKey(EquipmentKey string, userId int) ([]entities.EquipmentMasterEntities, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.SearchEquipmentByKey(trans, EquipmentKey, userId, s.cld)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *EquipmentCourseServiceImpl) GetEquipmentSearchHistoryByKey(userId int) ([]entities.EquipmentSearchHistoryEntities, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.GetEquipmentSearchHistoryByKey(trans, userId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *EquipmentCourseServiceImpl) DeleteEquipmentSearchHistoryById(equipmentSearchHistoryId int) (bool, *responses.ErrorResponses) {
	trans := s.db.Begin()
	res, err := s.repository.DeleteEquipmentSearchHistoryById(trans, equipmentSearchHistoryId)
	defer helper.CommitOrRollback(trans)
	if err != nil {
		return res, err
	}
	return res, nil
}
