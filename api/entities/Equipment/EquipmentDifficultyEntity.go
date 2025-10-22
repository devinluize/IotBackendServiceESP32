package entities

type EquipmentDifficultyEntities struct {
	EquipmentDifficultyId   int    `gorm:"column:equipment_difficulty_id;primaryKey"`
	EquipmentDifficultyName string `gorm:"column:equipment_difficulty_name;size:100;;not null"`
}

func (*EquipmentDifficultyEntities) TableName() string {
	return "mtr_equipment_difficulty"
}
