package blynkrepositoy

import blynkpayloads "IotBackend/api/payloads/blynk"

type BlynkRepository interface {
	SendDataToBlynk(request blynkpayloads.BlynkDataFromEsp32Request) error
}
