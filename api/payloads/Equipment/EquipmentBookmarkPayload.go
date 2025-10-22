package Equipment

type GetBookmarkEquipmentResponse struct {
	UserId              int    `json:"user_id"`
	EquipmentName       string `json:"equipment_name"`
	EquipmentId         int    `json:"equipment_id"`
	EquipmentCourseId   int    `json:"equipment_course_id"`
	EquipmentCourseName string `json:"equipment_course_name"`
	EquipmentPhotoPath  string `json:"equipment_photo_path"`
}
