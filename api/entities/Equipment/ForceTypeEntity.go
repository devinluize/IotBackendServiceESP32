package entities

const ForceTypeTableName = "mtr_force_type"

type ForceTypeEntities struct {
	ForceTypeName string `gorm:"column:force_type_name;size:100;not null"`
	ForceTypeId   int    `gorm:"column:force_type_id;not null;primaryKey"`
}

func (*ForceTypeEntities) TableName() string {
	return ForceTypeTableName
}
