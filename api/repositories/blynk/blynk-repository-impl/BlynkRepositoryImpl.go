package blynkrepositoryimpl

import (
	configenv "IotBackend/api/config"
	blynkpayloads "IotBackend/api/payloads/blynk"
	blynkrepositoy "IotBackend/api/repositories/blynk"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

type BlynkRepositoryImpl struct {
}

func NewBlynkRepositoryImpl() blynkrepositoy.BlynkRepository {
	return &BlynkRepositoryImpl{}
}

func (b *BlynkRepositoryImpl) SendDataToBlynk(tx *gorm.DB, request blynkpayloads.BlynkDataFromEsp32Request) error {
	_, err := http.Get(fmt.Sprintf("%supdate?token=%s&pin=%s&value=%f",
		configenv.EnvConfigs.BlynkAPIUrl,
		configenv.EnvConfigs.BlynkAPIToken,
		configenv.EnvConfigs.DataStreamTemperature,
		request.Temperature))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	_, err = http.Get(fmt.Sprintf("%supdate?token=%s&pin=%s&value=%f",
		configenv.EnvConfigs.BlynkAPIUrl,
		configenv.EnvConfigs.BlynkAPIToken,
		configenv.EnvConfigs.DataStreamSoilMoisture,
		request.SoilMoisture))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	_, err = http.Get(fmt.Sprintf("%supdate?token=%s&pin=%s&value=%f",
		configenv.EnvConfigs.BlynkAPIUrl,
		configenv.EnvConfigs.BlynkAPIToken,
		configenv.EnvConfigs.DataStreamLightIntensity,
		request.LightIntensity))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	_, err = http.Get(fmt.Sprintf("%supdate?token=%s&pin=%s&value=%f",
		configenv.EnvConfigs.BlynkAPIUrl,
		configenv.EnvConfigs.BlynkAPIToken,
		configenv.EnvConfigs.DataStreamHumidity,
		request.Humidity))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	_, err = http.Get(fmt.Sprintf("%supdate?token=%s&pin=%s&value=%d",
		configenv.EnvConfigs.BlynkAPIUrl,
		configenv.EnvConfigs.BlynkAPIToken,
		configenv.EnvConfigs.DataStreamPollutionLevel,
		request.AirPollution))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
