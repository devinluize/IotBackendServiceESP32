package entities

const EquipmentMasterEntityName = "mtr_equipment_master"

type EquipmentMasterEntities struct {
	EquipmentId        int    `gorm:"column:equipment_id;primary_key;AUTO_INCREMENT"`
	EquipmentName      string `gorm:"column:equipment_name"`
	EquipmentCode      string `gorm:"column:equipment_code"`
	EquipmentPhotoPath string `gorm:"equipment_photo_path" json:"equipment_photo_path"`

	//EquipmentTutorialVideoPath string `gorm:"column:equipment_tutorial_video_path"`
	//ForceTypeId                int    `gorm:"column:force_type_id"`
	//ForceType                  ForceTypeEntities
	//MuscleGroupId int `gorm:"column:muscle_group_id"`
	//MuscleGroup                MuscleGroupEntities
	//EquipmentTypeId int `gorm:"column:equipment_type_id"`
	//EquipmentType              EquipmentTypeEntity
	//EquipmentDifficultyId int `gorm:"column:equipment_difficulty_id"`
	//EquipmentDifficulty        EquipmentDifficultyEntities
	//EquipmentDetail []EquipmentDetailEntity `gorm:"foreignKey:EquipmentId;references:EquipmentId"`
}

func (e *EquipmentMasterEntities) TableName() string {
	return EquipmentMasterEntityName
}
