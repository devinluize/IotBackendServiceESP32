package entities

import "time"

type SearchHistoryEntities struct {
	SearchHistoryId int       `gorm:"column:search_history_id;primary_key;AUTO_INCREMENT" json:"search_history_id"`
	UserId          int       `gorm:"column:user_id" json:"user_id"`
	SearchKey       string    `gorm:"column:search_key" json:"search_key"`
	DateSearch      time.Time `gorm:"column:date_search" json:"date_search"`
}

func (*SearchHistoryEntities) TableName() string {
	return "trx_search_history"
}
