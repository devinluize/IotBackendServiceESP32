package entities

type EquipmentDetailEntity struct {
	EquipmentDetailId     int    `gorm:"column:equipment_detail_id;primaryKey"`
	EquipmentCourseDataId int    `gorm:"column:equipment_course_data_id" json:"equipment_course_data_id"`
	TutorialParagraph     string `gorm:"column:tutorial_paragraph"`
	TutorialPath          string `gorm:"column:tutorial_path"`
	ParagraphLineNumber   int    `gorm:"column:paragraph_line_number"`
}

func (*EquipmentDetailEntity) TableName() string {
	return "mtr_equipment_detail"
}
