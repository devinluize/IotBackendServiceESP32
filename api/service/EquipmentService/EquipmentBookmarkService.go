package EquipmentService

import (
	entities "IotBackend/api/entities/Equipment"
	"IotBackend/api/payloads/Equipment"
	"IotBackend/api/payloads/responses"
)

type EquipmentBookmarkService interface {
	AddEquipmentBookmark(userId, equipmentCourseId int) (entities.EquipmentBookmark, *responses.ErrorResponses)
	RemoveEquipmentBookmark(userId, equipmentCourseId int) (bool, *responses.ErrorResponses)
	GetEquipmentBookmarkByUserId(userId int) ([]Equipment.GetBookmarkEquipmentResponse, *responses.ErrorResponses)
}
