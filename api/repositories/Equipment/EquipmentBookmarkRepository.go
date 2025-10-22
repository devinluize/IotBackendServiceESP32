package menuRepository

import (
	entities "IotBackend/api/entities/Equipment"
	"IotBackend/api/payloads/Equipment"
	"IotBackend/api/payloads/responses"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
)

type EquipmentBookmarkRepository interface {
	AddEquipmentBookmark(db *gorm.DB, userId, equipmentCourseId int) (entities.EquipmentBookmark, *responses.ErrorResponses)
	RemoveEquipmentBookmark(db *gorm.DB, userId, equipmentCourseId int) (bool, *responses.ErrorResponses)
	GetEquipmentBookmarkByUserId(db *gorm.DB, userId int, cld *cloudinary.Cloudinary) ([]Equipment.GetBookmarkEquipmentResponse, *responses.ErrorResponses)
}
