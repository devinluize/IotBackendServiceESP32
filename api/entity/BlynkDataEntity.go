package entity

import "time"

type BlynkData struct {
	SoilMoisture   float64   `json:"soil_moisture"`
	LightIntensity float64   `json:"light_intensity"`
	Temperature    float64   `json:"temperature"`
	AirPollution   int       `json:"air_pollution"`
	Humidity       float64   `json:"humidity"`
	CreateAt       time.Time `json:"create_at"`
}
