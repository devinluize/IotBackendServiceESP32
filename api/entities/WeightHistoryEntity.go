package entities

import "time"

type WeightHistoryEntities struct {
	WeightHistoryId int `gorm:"weight_history_id;primaryKey;not null" json:"weight_history_id"`
	UserId          int `gorm:"column:user_id" json:"user_id"`
	User            Users
	UserWeight      float64   `gorm:"column:user_weight" json:"user_weight"`
	UserWeightTime  time.Time `gorm:"column:user_weight_time" json:"user_weight_time"`
	UserBmi         float64   `gorm:"column:user_bmi" json:"user_bmi"`
}

type WeightHistoryPayloads struct {
	WeightHistoryId int `gorm:"weight_history_id;primaryKey;not null" json:"weight_history_id"`
	UserId          int `gorm:"column:user_id" json:"user_id"`
	//User            Users
	UserWeight     float64   `gorm:"column:user_weight" json:"user_weight"`
	UserWeightTime time.Time `gorm:"column:user_weight_time" json:"user_weight_time"`
}

func (*WeightHistoryEntities) TableName() string {
	return "trx_weight_history"
}
