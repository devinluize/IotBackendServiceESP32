package MenuPayloads

type TimerInsertPayload struct {
	TimerId   int    `gorm:"column:timer_id" json:"timer_id"`
	TimerName string `json:"timer_name"`
	//UserId           int    `gorm:"column:user_id;" json:"user_id"`
	//RemindingHours   int    `gorm:"column:reminding_hours;" json:"reminding_hours"`
	//RemindingMinutes int    `gorm:"column:reminding_minutes;" json:"reminding_minutes"`
	TimerDescription string `json:"timer_description"`
}

type TimerQueueInsertResponse struct {
	TimerId                    int    `gorm:"column:timer_id" json:"timer_id"`
	TimerQueueName             string `gorm:"column:timer_queue_name;not null" json:"timer_queue_name"`
	TimerQueueRemindingHour    int    `gorm:"column:timer_queue_reminding_hour" json:"timer_queue_reminding_hour"`
	TimerQueueRemindingMinutes int    `gorm:"column:timer_queue_reminding_minutes" json:"timer_queue_reminding_minutes"`
	TimerQueueRemindingSecond  int    `gorm:"column:timer_queue_reminding_second" json:"timer_queue_reminding_second"`
}

type TimerQueueUpdatePayload struct {
	TimerQueueId               int    `gorm:"column:timer_queue_id" json:"timer_queue_id"`
	TimerQueueName             string `gorm:"column:timer_queue_name;not null" json:"timer_queue_name"`
	TimerQueueRemindingHour    int    `gorm:"column:timer_queue_reminding_hour" json:"timer_queue_reminding_hour"`
	TimerQueueRemindingMinutes int    `gorm:"column:timer_queue_reminding_minutes" json:"timer_queue_reminding_minutes"`
	TimerQueueRemindingSecond  int    `json:"timer_queue_reminding_second"`
}

type GetTimerResponse struct {
	TimerId   int    `gorm:"column:timer_id;primary_key;auto_increment;not null" json:"timer_id"`
	UserId    int    `gorm:"column:user_id;" json:"user_id"`
	TimerName string `gorm:"column:timer_name" json:"timer_name"`
}

type GetAllTimerByUserIdResponse struct {
	TimerId          int    `json:"timer_id"`
	UserId           int    `json:"user_id"`
	TimerName        string `gorm:"column:timer_name" json:"timer_name"`
	TimerDescription string `json:"timer_description"`
}
