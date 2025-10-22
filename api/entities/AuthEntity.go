package entities

type Users struct {
	UserId       int    `gorm:"column:user_id;primaryKey;not null" json:"user_id"`
	UserName     string `gorm:"column:user_name;size:255" json:"user_name"`
	UserEmail    string `gorm:"column:user_email;size:50" json:"user_email"`
	UserPassword string `gorm:"column:user_password;size:255" json:"user_password"`
	//IsVIP        bool       `gorm:"column:is_vip" json:"is_vip"`
	UserDetail UserDetail `gorm:"foreignKey:UserId;references:UserId" json:"user_detail"`
}

func (*Users) TableName() string {
	return "mtr_user"
}
