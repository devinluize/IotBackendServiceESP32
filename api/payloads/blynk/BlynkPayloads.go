package blynkpayloads

type BlynkDataFromEsp32Request struct {
	SoilMoisture   float64 `json:"soil_moisture"`
	LightIntensity float64 `json:"light_intensity"`
	Temperature    float64 `json:"temperature"`
	AirPollution   int     `json:"air_pollution"`
	Humidity       float64 `json:"humidity"`
}
