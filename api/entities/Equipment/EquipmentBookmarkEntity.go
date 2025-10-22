package entities

import "IotBackend/api/entities"

type EquipmentBookmark struct {
	EquipmentBookmarkId int `gorm:"column:equipment_bookmark_id;not null;primaryKey" json:"equipment_bookmark_id"`
	EquipmentCourseId   int `gorm:"column:equipment_course_id" json:"equipment_course_id"`
	EquipmentCourse     EquipmentCourseDataEntity
	UserId              int `gorm:"column:user_id" json:"user_id"`
	User                entities.Users
}

func (*EquipmentBookmark) TableName() string {
	return "mtr_equipment_bookmark"
}
