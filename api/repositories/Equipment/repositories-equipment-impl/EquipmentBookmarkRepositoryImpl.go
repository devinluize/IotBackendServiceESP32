package repositoriesEquipmentImpl

import (
	entities "IotBackend/api/entities/Equipment"
	"IotBackend/api/payloads/Equipment"
	"IotBackend/api/payloads/responses"
	menuRepository "IotBackend/api/repositories/Equipment"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
	"net/http"
)

type EquipmentBookmarkRepositoryImpl struct {
}

func NewEquipmentBookmarkRepositoryImpl() menuRepository.EquipmentBookmarkRepository {
	return &EquipmentBookmarkRepositoryImpl{}
}
func (e *EquipmentBookmarkRepositoryImpl) AddEquipmentBookmark(db *gorm.DB, userId, equipmentCourseId int) (entities.EquipmentBookmark, *responses.ErrorResponses) {
	entitiesEq := entities.EquipmentBookmark{
		EquipmentCourseId: equipmentCourseId,
		UserId:            userId,
	}

	isExist := 0
	errExist := db.Model(&entitiesEq).Where(entities.EquipmentBookmark{EquipmentCourseId: equipmentCourseId, UserId: userId}).
		Select("1").Scan(&isExist).Error
	if errExist != nil {
		return entitiesEq, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        errExist,
			Message:    "failed to check duplicate bookmark",
		}
	}
	if isExist == 1 {
		return entitiesEq, &responses.ErrorResponses{
			StatusCode: http.StatusBadRequest,
			Message:    "cannot insert duplicate bookmark",
			Err:        errors.New("cannot insert duplicate bookmark"),
			Success:    false,
			Data:       nil,
		}
	}

	err := db.Create(&entitiesEq).First(&entitiesEq).Error
	if err != nil {
		return entitiesEq, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to create bookmark",
		}
	}
	return entitiesEq, nil
}

func (e *EquipmentBookmarkRepositoryImpl) RemoveEquipmentBookmark(db *gorm.DB, userId, equipmentCourseId int) (bool, *responses.ErrorResponses) {
	entitiesEq := entities.EquipmentBookmark{}

	err := db.Delete(&entitiesEq, entities.EquipmentBookmark{EquipmentCourseId: equipmentCourseId, UserId: userId}).Error
	if err != nil {
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to remove equipment bookmark",
		}
	}
	return true, nil
}
func (e *EquipmentBookmarkRepositoryImpl) GetEquipmentBookmarkByUserId(db *gorm.DB, userId int, cld *cloudinary.Cloudinary) ([]Equipment.GetBookmarkEquipmentResponse, *responses.ErrorResponses) {
	var response []Equipment.GetBookmarkEquipmentResponse
	var equipmentBookmark []entities.EquipmentBookmark
	err := db.Model(&entities.EquipmentBookmark{}).
		Where(entities.EquipmentBookmark{UserId: userId}).
		Scan(&equipmentBookmark).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to get equipment bookmark",
		}
	}

	for _, bookmark := range equipmentBookmark {
		//get equipment names
		equipmentCourseEntities := entities.EquipmentCourseDataEntity{}
		err = db.Model(&equipmentCourseEntities).
			Where(entities.EquipmentCourseDataEntity{EquipmentCourseDataId: bookmark.EquipmentCourseId}).
			First(&equipmentCourseEntities).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response, &responses.ErrorResponses{
					StatusCode: http.StatusInternalServerError,
					Err:        err,
					Message:    "course data is not found",
				}
			}
			return response, &responses.ErrorResponses{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
				Message:    "failed to get equipment course entities",
			}
		}
		//get equipment data
		equipmentMaster := entities.EquipmentMasterEntities{}
		err = db.Model(&equipmentMaster).
			Where(entities.EquipmentMasterEntities{EquipmentId: equipmentCourseEntities.EquipmentMasterId}).
			First(&equipmentMaster).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response, &responses.ErrorResponses{
					StatusCode: http.StatusNotFound,
					Err:        err,
					Message:    "equipment master is not found",
				}
			}
		}
		urls, _ := cld.Image(equipmentMaster.EquipmentPhotoPath)
		//res.SortOf = url
		equipmentMaster.EquipmentPhotoPath = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
			"dlrd9z1mk",          // Replace with your Cloudinary cloud name
			urls.AssetType,       // e.g., "image"
			urls.DeliveryType,    // e.g., "upload"
			urls.PublicID+".jpg", // Add appropriate file extension
		)
		resp := Equipment.GetBookmarkEquipmentResponse{
			UserId:              bookmark.UserId,
			EquipmentName:       equipmentMaster.EquipmentName,
			EquipmentId:         equipmentMaster.EquipmentId,
			EquipmentCourseId:   equipmentCourseEntities.EquipmentCourseDataId,
			EquipmentCourseName: equipmentCourseEntities.EquipmentCourseDataName,
			EquipmentPhotoPath:  equipmentMaster.EquipmentPhotoPath,
		}
		response = append(response, resp)
	}
	return response, nil

}
