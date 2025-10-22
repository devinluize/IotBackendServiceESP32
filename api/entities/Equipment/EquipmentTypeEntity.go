package entities

const EquipmentTypeTableName = "mtr_equipment_type"

type EquipmentTypeEntity struct {
	EquipmentTypeId   int    `gorm:"column:equipment_type_id;not null;primaryKey"`
	EquipmentTypeName string `gorm:"column:equipment_type_name;size:100;not null"`
}

func (*EquipmentTypeEntity) TableName() string {
	return EquipmentTypeTableName
}
