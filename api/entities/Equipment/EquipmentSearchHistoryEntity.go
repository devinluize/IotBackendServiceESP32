package entities

import "time"

type EquipmentSearchHistoryEntities struct {
	EquipmentSearchHistoryId int       `gorm:"column:equipment_search_history_id;primary_key;AUTO_INCREMENT" json:"equipment_search_history_id"`
	UserId                   int       `gorm:"column:user_id" json:"user_id"`
	SearchKey                string    `gorm:"column:search_key" json:"search_key"`
	DateSearch               time.Time `gorm:"column:date_search" json:"date_search"`
}

func (*EquipmentSearchHistoryEntities) TableName() string {
	return "trx_equipment_search_history"
}
