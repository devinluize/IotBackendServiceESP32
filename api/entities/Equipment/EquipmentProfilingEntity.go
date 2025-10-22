package entities

const EquipmentProfilingTableName = "mtr_equipment_profile"

type EquipmentProfileEntity struct {
	EquipmentProfileId   int    `gorm:"column:equipment_profile_id;not null;primaryKey"`
	EquipmentProfileName string `gorm:"column:equipment_profile_name;not null"`
}

func (*EquipmentProfileEntity) TableName() string {
	return EquipmentProfilingTableName
}
