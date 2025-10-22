package entities

const EquipmentCourseDataEntityTableName = "mtr_equipment_course_data"

type EquipmentCourseDataEntity struct {
	EquipmentCourseDataId    int    `gorm:"column:equipment_mapping_data_id;primaryKey" json:"equipment_mapping_data_id"`
	EquipmentCourseDataName  string `gorm:"column:equipment_mapping_data_name;size:50" json:"equipment_mapping_data_entity_name"`
	EquipmentMasterId        int    `gorm:"column:equipment_master_id" json:"equipment_master_id"`
	EquipmentMusclePhotoPath string `gorm:"equipment_muscle_photo_path;size:100" json:"equipment_muscle_photo_path"`
	EquipmentMaster          EquipmentMasterEntities
	VideoTutorialVideoPath   string `gorm:"column:video_tutorial_video_path;size:100" json:"video_tutorial_video_path"`
	EquipmentDifficultyId    int    `gorm:"column:equipment_difficulty_id" json:"equipment_difficulty_id"`
	EquipmentDifficulty      EquipmentDifficultyEntities
	EquipmentTypeId          int `gorm:"column:equipment_type_id" json:"equipment_type_id"`
	EquipmentType            EquipmentTypeEntity
	ForceTypeId              int `gorm:"column:force_type_id" json:"force_type_id"`
	ForceType                ForceTypeEntities
	MuscleGroupId            int `gorm:"column:muscle_group_id" json:"muscle_group_id"`
	MuscleGroup              MuscleGroupEntities
	EquipmentProfileId       int `json:"equipment_profile_id"`
	EquipmentProfile         EquipmentProfileEntity
	EquipmentDetail          []EquipmentDetailEntity `gorm:"foreignKey:EquipmentCourseDataId;references:EquipmentCourseDataId" json:"equipment_detail"`
}

func (*EquipmentCourseDataEntity) TableName() string {
	return EquipmentCourseDataEntityTableName

}
