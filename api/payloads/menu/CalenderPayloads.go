package MenuPayloads

import "time"

type CalendarInsertPayload struct {
	EventName     string    `gorm:"column:event_name" json:"event_name"`
	EventDate     time.Time `gorm:"column:event_date" json:"event_date"`
	UserId        int       `gorm:"column:user_id" json:"user_id"`
	EventTimeFrom time.Time `gorm:"column:event_time_from" json:"event_time_from"`
	EventTimeTo   time.Time `gorm:"column:event_time_to" json:"event_time_to"`
}
type CalendarUpdatePayload struct {
	EventId       int       `gorm:"column:event_id;primaryKey;not null" json:"event_id"`
	EventName     string    `gorm:"column:event_name" json:"event_name"`
	EventDate     time.Time `gorm:"column:event_date" json:"event_date"`
	UserId        int       `gorm:"column:user_id" json:"user_id"`
	EventTimeFrom time.Time `gorm:"column:event_time_from" json:"event_time_from"`
	EventTimeTo   time.Time `gorm:"column:event_time_to" json:"event_time_to"`
}
type CalendarGetByIdResponse struct {
	EventName     string    `gorm:"column:event_name" json:"event_name"`
	EventDate     time.Time `gorm:"column:event_date" json:"event_date"`
	UserId        int       `gorm:"column:user_id" json:"user_id"`
	EventTimeFrom time.Time `gorm:"column:event_time_from" json:"event_time_from"`
	EventTimeTo   time.Time `gorm:"column:event_time_to" json:"event_time_to"`
	EventId       int       `gorm:"column:event_id;primaryKey;not null" json:"event_id"`
}
