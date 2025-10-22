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
	"time"
)

type EquipmentCourseRepositoryImpl struct {
}

func NewEquipmentCourseRepositoryImpl() menuRepository.EquipmentCourseRepository {
	return &EquipmentCourseRepositoryImpl{}

}
func (e *EquipmentCourseRepositoryImpl) GetAllEquipmentCourseByEquipment(db *gorm.DB, equipmentId int, cld *cloudinary.Cloudinary) (Equipment.GetAllCourseEquipmentResponse, *responses.ErrorResponses) {
	var mappingEntities []entities.EquipmentCourseDataEntity
	var response Equipment.GetAllCourseEquipmentResponse
	if equipmentId != 0 {
		err := db.Model(&entities.EquipmentCourseDataEntity{}).
			Where(entities.EquipmentCourseDataEntity{EquipmentMasterId: equipmentId}).
			Scan(&mappingEntities).Error
		if err != nil {
			return response, &responses.ErrorResponses{
				Message:    err.Error(),
				Err:        err,
				StatusCode: http.StatusInternalServerError,
			}
		}
	} else {
		return response, nil
	}
	//respons := Equipment.GetAllCourseEquipmentResponse{}

	//get item name first
	equipmentMasterEntities := entities.EquipmentMasterEntities{}
	if equipmentId != 0 {

		err := db.Model(&equipmentMasterEntities).
			Where(entities.EquipmentMasterEntities{EquipmentId: equipmentId}).
			Scan(&equipmentMasterEntities).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response, &responses.ErrorResponses{
					StatusCode: http.StatusNotFound,
					Err:        err,
					Message:    "failed to get item",
				}
			}
			return response, &responses.ErrorResponses{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
				Message:    "failed to get item name please check input",
			}
		}
	}

	//response.EquipmentName = equipmentMasterEntities.EquipmentName
	//response.EquipmentId = equipmentMasterEntities.EquipmentId
	//response.EquipmentPhotoPath =
	for _, mappingEntity := range mappingEntities {
		detail := Equipment.EquipmentMappingDataResponse{
			EquipmentMappingId:   mappingEntity.EquipmentCourseDataId,
			EquipmentMappingName: mappingEntity.EquipmentCourseDataName,
		}
		response.EquipmentMappingData = append(response.EquipmentMappingData, detail)
		response.EquipmentName = equipmentMasterEntities.EquipmentName
		response.EquipmentId = equipmentMasterEntities.EquipmentId
	}
	urls, errImage := cld.Image(equipmentMasterEntities.EquipmentPhotoPath)
	if errImage != nil {
		return response, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        errImage,
			Message:    "error getting image with public id",
		}
	}
	//res.SortOf = url
	response.EquipmentPhotoPath = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
		"dlrd9z1mk",          // Replace with your Cloudinary cloud name
		urls.AssetType,       // e.g., "image"
		urls.DeliveryType,    // e.g., "upload"
		urls.PublicID+".jpg", // Add appropriate file extension
	)

	return response, nil
}
func (e *EquipmentCourseRepositoryImpl) InsertEquipmentCourse(db *gorm.DB, payload Equipment.InsertEquipmentCourseDataPayload) (entities.EquipmentCourseDataEntity, *responses.ErrorResponses) {

	entitiesCourseData := entities.EquipmentCourseDataEntity{
		EquipmentCourseDataName: payload.EquipmentCourseName,
		EquipmentMasterId:       payload.EquipmentMasterId,
		VideoTutorialVideoPath:  payload.VideoTutorialVideoPath,
		EquipmentDifficultyId:   payload.EquipmentDifficultyId,
		EquipmentTypeId:         payload.EquipmentTypeId,
		ForceTypeId:             payload.EquipmentForceTypeId,
		MuscleGroupId:           payload.MuscleGroupId,
		EquipmentProfileId:      payload.EquipmentProfilingId,
	}
	err := db.Create(&entitiesCourseData).First(&entitiesCourseData).Error
	if err != nil {
		return entitiesCourseData, &responses.ErrorResponses{
			Message:    err.Error(),
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}
	paragraphLineNumber := 1
	for _, detail := range payload.InsertEquipmentDetailCourse {

		entitiesDetail := entities.EquipmentDetailEntity{
			//EquipmentDetailId:     0,
			EquipmentCourseDataId: entitiesCourseData.EquipmentCourseDataId,
			TutorialParagraph:     detail.TutorialParagraph,
			TutorialPath:          detail.TutorialPath,
			ParagraphLineNumber:   paragraphLineNumber,
		}
		err = db.Create(&entitiesDetail).First(&entitiesDetail).Error
		if err != nil {
			return entitiesCourseData, &responses.ErrorResponses{
				Message:    err.Error(),
				Err:        err,
				StatusCode: http.StatusInternalServerError,
			}
		}
		paragraphLineNumber += 1
	}
	err = db.Model(&entitiesCourseData).Where(&entities.EquipmentCourseDataEntity{EquipmentCourseDataId: entitiesCourseData.EquipmentCourseDataId}).
		First(&entitiesCourseData).Error
	if err != nil {
		return entitiesCourseData, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to get course data",
		}
	}
	return entitiesCourseData, nil
}
func (e *EquipmentCourseRepositoryImpl) GetEquipmentCourse(db *gorm.DB, courseId int, cld *cloudinary.Cloudinary, userId int) (Equipment.GetCourseByIdResponse, *responses.ErrorResponses) {
	courseEntities := entities.EquipmentCourseDataEntity{}
	response := Equipment.GetCourseByIdResponse{}
	err := db.Model(&courseEntities).Where(entities.EquipmentCourseDataEntity{EquipmentCourseDataId: courseId}).First(&courseEntities).Error
	if err != nil {
		return response, &responses.ErrorResponses{
			Message:    err.Error(),
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}
	urls, errImage := cld.Image(courseEntities.EquipmentMusclePhotoPath)
	if errImage != nil {
		return response, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "error getting image with public id",
		}
	}
	//res.SortOf = url
	courseEntities.EquipmentMusclePhotoPath = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
		"dlrd9z1mk",          // Replace with your Cloudinary cloud name
		urls.AssetType,       // e.g., "image"
		urls.DeliveryType,    // e.g., "upload"
		urls.PublicID+".jpg", // Add appropriate file extension
	)
	urlsVideo, errVideo := cld.Video(courseEntities.VideoTutorialVideoPath)
	if errVideo != nil {
		return response, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "error getting video with public id",
		}
	}
	//res.SortOf = url
	courseEntities.VideoTutorialVideoPath = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
		"dlrd9z1mk",               // Replace with your Cloudinary cloud name
		urlsVideo.AssetType,       // e.g., "image"
		urlsVideo.DeliveryType,    // e.g., "upload"
		urlsVideo.PublicID+".mp4", // Add appropriate file extension
	)

	var courseDetailEntities []entities.EquipmentDetailEntity
	err = db.Model(&entities.EquipmentDetailEntity{}).
		Where(entities.EquipmentDetailEntity{EquipmentCourseDataId: courseId}).
		Scan(&courseDetailEntities).Error
	if err != nil {
		return response, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    err.Error(),
		}
	}

	//get equipmentMasterName
	equipmentMaster := entities.EquipmentMasterEntities{}
	err = db.Model(&equipmentMaster).Where(entities.EquipmentMasterEntities{EquipmentId: courseEntities.EquipmentMasterId}).
		First(&equipmentMaster).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Message:    "equipment with that id is not found",
				Err:        err,
			}
		}
		return response, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    err.Error(),
		}
	}
	//EquipmentDifficultyName
	EquipmentDifficultyEntities := entities.EquipmentDifficultyEntities{}
	err = db.Model(&EquipmentDifficultyEntities).
		Where(entities.EquipmentDifficultyEntities{EquipmentDifficultyId: courseEntities.EquipmentDifficultyId}).
		First(&EquipmentDifficultyEntities).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Message:    "equipment difficulty with that id is not found",
				Err:        err,
			}
		}
		return response, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    err.Error(),
		}
	}
	//EquipmentTypeName
	EquipmentTypeEntities := entities.EquipmentTypeEntity{}
	err = db.Model(&EquipmentTypeEntities).
		Where(entities.EquipmentTypeEntity{EquipmentTypeId: courseEntities.EquipmentTypeId}).
		First(&EquipmentTypeEntities).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Err:        err,
				Message:    "equipment type is not found with that id please check input",
			}
		}
		return response, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to get equipment type please check input",
		}
	}
	//ForceType
	ForceTypeEntities := entities.ForceTypeEntities{}
	err = db.Model(&ForceTypeEntities).
		Where(entities.ForceTypeEntities{ForceTypeId: courseEntities.ForceTypeId}).
		First(&ForceTypeEntities).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Err:        err,
				Message:    "failed to find force type with that id please check input",
			}
		}
	}
	//MuscleGroupName
	MuscleGroup := entities.MuscleGroupEntities{}
	err = db.Model(&MuscleGroup).
		Where(entities.MuscleGroupEntities{MuscleGroupId: courseEntities.MuscleGroupId}).
		First(&MuscleGroup).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Err:        err,
				Message:    "failed to get muscle group",
			}
		}
		return response, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to get musvle group name",
			Err:        err,
		}
	}
	EquipmentProfile := entities.EquipmentProfileEntity{}
	err = db.Model(&EquipmentProfile).
		Where(entities.EquipmentProfileEntity{EquipmentProfileId: courseEntities.EquipmentProfileId}).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, &responses.ErrorResponses{
				StatusCode: http.StatusNotFound,
				Message:    "equipment profile with that id is not found please check input",
				Err:        err,
			}
		}
		return response, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to get equipment profile",
			Err:        err,
		}
	}
	//cek if bookmark
	var isBookmark = false
	err = db.Model(&entities.EquipmentBookmark{}).
		Where(entities.EquipmentBookmark{EquipmentCourseId: courseId,
			UserId: userId,
			//UserId:
		}).
		Select("1").Scan(&isBookmark).Error
	if err != nil {
		return response, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to check bookmark",
		}
	}
	response = Equipment.GetCourseByIdResponse{
		EquipmentCourseDataId:    courseEntities.EquipmentCourseDataId,
		EquipmentCourseDataName:  courseEntities.EquipmentCourseDataName,
		EquipmentMasterId:        courseEntities.EquipmentMasterId,
		EquipmentMasterName:      equipmentMaster.EquipmentName,
		VideoTutorialVideoPath:   courseEntities.VideoTutorialVideoPath,
		EquipmentDifficultyId:    courseEntities.EquipmentDifficultyId,
		EquipmentDifficultyName:  EquipmentDifficultyEntities.EquipmentDifficultyName,
		EquipmentTypeId:          courseEntities.EquipmentTypeId,
		EquipmentTypeName:        EquipmentTypeEntities.EquipmentTypeName,
		ForceTypeId:              courseEntities.ForceTypeId,
		ForceTypeName:            ForceTypeEntities.ForceTypeName,
		MuscleGroupId:            courseEntities.MuscleGroupId,
		MuscleGroupName:          MuscleGroup.MuscleGroupName,
		EquipmentProfileId:       courseEntities.EquipmentProfileId,
		EquipmentProfileName:     EquipmentProfile.EquipmentProfileName,
		EquipmentDetail:          courseDetailEntities,
		IsBookmark:               isBookmark,
		EquipmentMusclePhotoPath: courseEntities.EquipmentMusclePhotoPath,
	}
	return response, nil

}
func (e *EquipmentCourseRepositoryImpl) SearchEquipmentByKey(db *gorm.DB, EquipmentKey string, userId int, cld *cloudinary.Cloudinary) ([]entities.EquipmentMasterEntities, *responses.ErrorResponses) {
	//get entities with equipment key
	var EquipmentResponse []entities.EquipmentMasterEntities
	if EquipmentKey != "" {

		historyLogging := entities.EquipmentSearchHistoryEntities{
			UserId:     userId,
			SearchKey:  EquipmentKey,
			DateSearch: time.Now(),
		}

		// Insert the new history record
		err := db.Create(&historyLogging).Error
		if err != nil {
			return EquipmentResponse, &responses.ErrorResponses{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
				Message:    "failed to log search history",
			}
		}

		// Check the total count of search history records for the user
		//historyCount := 0
		var historyCount int64
		errCount := db.Model(&entities.EquipmentSearchHistoryEntities{}).
			Where("user_id = ?", userId).
			Count(&historyCount).Error
		if errCount != nil {
			return EquipmentResponse, &responses.ErrorResponses{
				StatusCode: http.StatusInternalServerError,
				Err:        errCount,
				Message:    "failed to check search history count",
			}
		}

		if historyCount > 10 {
			var excessCount int
			excessCount = int(historyCount) - 10
			// Delete oldest records if the count exceeds the limit
			var recordsToDelete []entities.EquipmentSearchHistoryEntities
			errSelect := db.Where("user_id = ?", userId).
				Order("date_search ASC").
				Limit(excessCount).
				Find(&recordsToDelete).Error
			if errSelect != nil {
				return EquipmentResponse, &responses.ErrorResponses{
					StatusCode: http.StatusInternalServerError,
					Err:        errSelect,
					Message:    "failed to fetch old search history records",
				}
			}

			// Step 2: Delete the fetched records
			if len(recordsToDelete) > 0 {
				errDelete := db.Delete(&recordsToDelete).Error
				if errDelete != nil {
					return EquipmentResponse, &responses.ErrorResponses{
						StatusCode: http.StatusInternalServerError,
						Err:        errDelete,
						Message:    "failed to clean up old search history records",
					}
				}
			}
		}
	}

	/////////
	err := db.Model(&entities.EquipmentMasterEntities{}).
		Where("equipment_id <> 0 AND equipment_name LIKE ? ", "%"+EquipmentKey+"%").
		Scan(&EquipmentResponse).Error
	if err != nil {
		return EquipmentResponse, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to get equipment by key",
		}
	}
	for i, masterEntities := range EquipmentResponse {
		urls, errImage := cld.Image(masterEntities.EquipmentPhotoPath)
		if errImage != nil {
			return EquipmentResponse, &responses.ErrorResponses{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
				Message:    "error getting image with public id",
			}
		}
		//res.SortOf = url
		EquipmentResponse[i].EquipmentPhotoPath = fmt.Sprintf("https://res.cloudinary.com/%s/%s/%s/%s",
			"dlrd9z1mk",          // Replace with your Cloudinary cloud name
			urls.AssetType,       // e.g., "image"
			urls.DeliveryType,    // e.g., "upload"
			urls.PublicID+".jpg", // Add appropriate file extension
		)
	}

	return EquipmentResponse, nil
}
func (e *EquipmentCourseRepositoryImpl) GetEquipmentSearchHistoryByKey(db *gorm.DB, userId int) ([]entities.EquipmentSearchHistoryEntities, *responses.ErrorResponses) {
	var model []entities.EquipmentSearchHistoryEntities
	err := db.Model(&model).
		Where(entities.EquipmentSearchHistoryEntities{UserId: userId}).
		Order("date_search DESC").
		Limit(10).
		Scan(&model).Error
	if err != nil {
		return model, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to get equipment search history by user id",
		}
	}
	return model, nil
}
func (e *EquipmentCourseRepositoryImpl) DeleteEquipmentSearchHistoryById(db *gorm.DB, equipmentSearchHistoryId int) (bool, *responses.ErrorResponses) {
	err := db.Delete(&entities.EquipmentSearchHistoryEntities{}, entities.EquipmentSearchHistoryEntities{EquipmentSearchHistoryId: equipmentSearchHistoryId}).Error
	if err != nil {
		return false, &responses.ErrorResponses{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
			Message:    "failed to delete equipment search",
		}
	}
	return true, nil
}
