package payloads

type RegisterPayloads struct {
	Username        string  `gorm:"UserName" json:"user_name"`
	Useremail       string  `gorm:"UserEmail" json:"user_email"`
	Userpasword     string  `gorm:"UserPassword" json:"user_password"`
	UserPhoneNumber string  `gorm:"UserPhoneNumber" json:"user_phone_number"`
	UserGender      string  `gorm:"UserGender" json:"user_gender"`
	UserHeight      float64 `gorm:"UserHeight" json:"user_height"`
	UserWeight      float64 `gorm:"UserWeight" json:"user_weight"`
}

type LoginPaylods struct {
	Useremail   string `gorm:"UserEmail" json:"user_email"`
	Userpasword string `gorm:"UserPassword" json:"user_password"`
}
type LoginRespons struct {
	UserName  string `json:"username"`
	UserEmail string `json:"userEmail"`
	IsVIP     bool   `json:"isVIP"`
	Token     string `json:"token"`
}

type UserBmiResponse struct {
	UserId     int     `json:"user_id"`
	Bmi        float64 `json:"bmi"`
	UserWeight float64 `json:"user_weight"`
	UserHeight float64 `json:"user_height"`
}

type GetUserDetailById struct {
	UserDetailId           int     `gorm:"column:user_detail_id;primaryKey;not null" json:"user_detail_id"`
	UserId                 int     `gorm:"column:user_id" json:"user_id"`
	UserName               string  `json:"user_name"`
	UserWeight             float64 `gorm:"column:user_weight" json:"user_weight"`
	UserHeight             float64 `gorm:"column:user_height" json:"user_height"`
	UserGender             string  `gorm:"column:user_gender;size:1" json:"user_gender"`
	UserProfileDescription string  `gorm:"column:user_profile_description" json:"user_profile_description"`
	UserProfileImage       string  `gorm:"column:user_profile_image" json:"user_profile_image"`
	UserPhoneNumber        string  `gorm:"column:user_phone_number" json:"user_phone_number"`
	//UserName string `json:"user_name"`
	UserEmail string `json:"user_email"`
}
