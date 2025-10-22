package MenuPayloads

import "time"

type WeightHistoryPayloads struct {
	//UserId         int       `gorm:"column:user_id" json:"user_id"`
	UserWeight     float64   `gorm:"column:user_weight" json:"user_weight"`
	UserWeightTime time.Time `gorm:"column:user_weight_time" json:"user_weight_time"`
	UserBmi        float64   `json:"user_bmi"`
}

type LastWeightResponse struct {
	UserWeightTime time.Time `json:"user_weight_time"`
	UserId         int       `json:"user_id"`
	UserWeight     float64   `json:"user_weight"`
	UserBmi        float64   `json:"user_bmi"`
}

type WeightHistoryGetAllResponse struct {
	WeightHistoryId int       `gorm:"weight_history_id;primaryKey;not null" json:"weight_history_id"`
	UserId          int       `gorm:"column:user_id" json:"user_id"`
	UserWeight      float64   `gorm:"column:user_weight" json:"user_weight"`
	UserWeightTime  time.Time `gorm:"column:user_weight_time" json:"user_weight_time"`
	UserBmi         float64   `json:"user_bmi"`
}

type GetAllFilterCondition struct {
	DateFrom time.Time `gorm:"column:date_from" json:"date_from"`
	DateTo   time.Time `gorm:"column:date_to" json:"date_to"`
}
