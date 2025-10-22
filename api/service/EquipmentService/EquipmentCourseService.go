package EquipmentService

import (
	entities "IotBackend/api/entities/Equipment"
	"IotBackend/api/payloads/Equipment"
	"IotBackend/api/payloads/responses"
)

type EquipmentCourseService interface {
	GetAllEquipmentCourseByEquipment(equipmentId int) (Equipment.GetAllCourseEquipmentResponse, *responses.ErrorResponses)
	InsertEquipmentCourse(payload Equipment.InsertEquipmentCourseDataPayload) (entities.EquipmentCourseDataEntity, *responses.ErrorResponses)
	GetEquipmentCourse(courseId int, userId int) (Equipment.GetCourseByIdResponse, *responses.ErrorResponses)
	SearchEquipmentByKey(EquipmentKey string, userId int) ([]entities.EquipmentMasterEntities, *responses.ErrorResponses)
	GetEquipmentSearchHistoryByKey(userId int) ([]entities.EquipmentSearchHistoryEntities, *responses.ErrorResponses)
	DeleteEquipmentSearchHistoryById(equipmentSearchHistoryId int) (bool, *responses.ErrorResponses)
}
