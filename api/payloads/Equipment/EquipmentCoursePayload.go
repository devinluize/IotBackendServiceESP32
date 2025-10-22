package Equipment

import entities "IotBackend/api/entities/Equipment"

type EquipmentMappingDataResponse struct {
	EquipmentMappingId   int    `json:"equipment_mapping_id"`
	EquipmentMappingName string `json:"equipment_mapping_name"`
}
type GetAllCourseEquipmentResponse struct {
	EquipmentId          int                            `json:"equipment_id"`
	EquipmentName        string                         `json:"equipment_name"`
	EquipmentPhotoPath   string                         `json:"equipment_photo_path"`
	EquipmentMappingData []EquipmentMappingDataResponse `json:"equipment_mapping_data"`
}
type InsertEquipmentDetailCoursePayload struct {
	EquipmentCourseDataId int    `gorm:"column:equipment_course_data_id" json:"equipment_course_data_id"`
	TutorialParagraph     string `gorm:"column:tutorial_paragraph" json:"tutorial_paragraph"`
	TutorialPath          string `gorm:"column:tutorial_path" json:"tutorial_path"`
}
type InsertEquipmentCourseDataPayload struct {
	EquipmentMasterId           int                                  `json:"equipment_master_id"`
	EquipmentCourseName         string                               `json:"equipment_course_name"`
	MuscleGroupId               int                                  `json:"muscle_group_id"`
	EquipmentTypeId             int                                  `json:"equipment_type_id"`
	EquipmentForceTypeId        int                                  `json:"equipment_force_type_id"`
	EquipmentProfilingId        int                                  `json:"equipment_profiling_id"`
	EquipmentDifficultyId       int                                  `json:"equipment_difficulty_id"`
	VideoTutorialVideoPath      string                               `gorm:"column:video_tutorial_video_path" json:"video_tutorial_video_path"`
	InsertEquipmentDetailCourse []InsertEquipmentDetailCoursePayload `json:"detail_course"`
}

type GetCourseByIdResponse struct {
	EquipmentCourseDataId    int                              `gorm:"column:equipment_mapping_data_id;primaryKey" json:"equipment_mapping_data_id"`
	EquipmentCourseDataName  string                           `gorm:"column:equipment_mapping_data_name" json:"equipment_mapping_data_entity_name"`
	EquipmentMasterId        int                              `gorm:"column:equipment_master_id" json:"equipment_master_id"`
	EquipmentMasterName      string                           `json:"equipment_master_name"`
	VideoTutorialVideoPath   string                           `gorm:"column:video_tutorial_video_path" json:"video_tutorial_video_path"`
	EquipmentDifficultyId    int                              `gorm:"column:equipment_difficulty_id" json:"equipment_difficulty_id"`
	EquipmentDifficultyName  string                           `json:"equipment_difficulty_name"`
	EquipmentTypeId          int                              `gorm:"column:equipment_type_id" json:"equipment_type_id"`
	EquipmentTypeName        string                           `json:"equipment_type_name"`
	ForceTypeId              int                              `gorm:"column:force_type_id" json:"force_type_id"`
	ForceTypeName            string                           `json:"force_type_name"`
	MuscleGroupId            int                              `gorm:"column:muscle_group_id" json:"muscle_group_id"`
	MuscleGroupName          string                           `json:"muscle_group_name"`
	EquipmentProfileId       int                              `json:"equipment_profile_id"`
	EquipmentProfileName     string                           `json:"equipment_profile_name"`
	EquipmentDetail          []entities.EquipmentDetailEntity `json:"equipment_detail"`
	IsBookmark               bool                             `json:"is_bookmark"`
	EquipmentMusclePhotoPath string                           `gorm:"equipment_muscle_photo_path" json:"equipment_muscle_photo_path"`
}
type AiLensPayload struct {
	//CloudinaryPublicId string `json:"cloudinary_public_id"`
	ImageUrl string `json:"image_url"`
	UserId   int    `json:"user_id"`
}
type AiLensEquipmentMasterResponse struct {
	EquipmentMasterId   int    `json:"equipment_master_id"`
	EquipmentMasterCode string `json:"equipment_master_code"`
}
type AiLensResponse struct {
	ApiSuccess  bool                           `json:"api_success"`
	ApiMessage  string                         `json:"api_message"`
	ApiResponse *AiLensEquipmentMasterResponse `json:"api_response"`
}
