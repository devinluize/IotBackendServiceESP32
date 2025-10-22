package entities

type PaymentMethod struct {
	PaymentMethodId   int    `gorm:"column:payment_method_id;primaryKey;not null" json:"payment_method_id"`
	PaymentMethodName string `gorm:"payment_method_name" json:"payment_method_name"`
}

func (*PaymentMethod) TableName() string {
	return "mtr_payment_method"
}
