package menuRepository

import (
	entities "IotBackend/api/entities/Equipment"
	"IotBackend/api/payloads/Equipment"
	"IotBackend/api/payloads/responses"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
)

type EquipmentCourseRepository interface {
	GetAllEquipmentCourseByEquipment(db *gorm.DB, equipmentId int, cld *cloudinary.Cloudinary) (Equipment.GetAllCourseEquipmentResponse, *responses.ErrorResponses)
	InsertEquipmentCourse(db *gorm.DB, payload Equipment.InsertEquipmentCourseDataPayload) (entities.EquipmentCourseDataEntity, *responses.ErrorResponses)
	GetEquipmentCourse(db *gorm.DB, courseId int, cld *cloudinary.Cloudinary, userId int) (Equipment.GetCourseByIdResponse, *responses.ErrorResponses)
	SearchEquipmentByKey(db *gorm.DB, EquipmentKey string, userId int, cld *cloudinary.Cloudinary) ([]entities.EquipmentMasterEntities, *responses.ErrorResponses)
	GetEquipmentSearchHistoryByKey(db *gorm.DB, userId int) ([]entities.EquipmentSearchHistoryEntities, *responses.ErrorResponses)
	DeleteEquipmentSearchHistoryById(db *gorm.DB, equipmentSearchHistoryId int) (bool, *responses.ErrorResponses)
}
