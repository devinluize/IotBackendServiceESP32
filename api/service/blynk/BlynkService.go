package blynkservice

import (
	blynkpayloads "IotBackend/api/payloads/blynk"
)

type BlynkService interface {
	SendDataToBlynk(request blynkpayloads.BlynkDataFromEsp32Request) error
}
