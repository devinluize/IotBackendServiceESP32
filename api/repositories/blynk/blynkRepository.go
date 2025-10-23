package blynkrepositoy

import (
	blynkpayloads "IotBackend/api/payloads/blynk"
	"gorm.io/gorm"
)

type BlynkRepository interface {
	SendDataToBlynk(tx *gorm.DB, request blynkpayloads.BlynkDataFromEsp32Request) error
}
