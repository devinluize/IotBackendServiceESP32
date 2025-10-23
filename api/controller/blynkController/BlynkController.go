package blynkController

import (
	"IotBackend/api/helper"
	encrypthelper "IotBackend/api/helper/encrypt"
	blynkpayloads "IotBackend/api/payloads/blynk"
	"IotBackend/api/payloads/responses"
	blynkservice "IotBackend/api/service/blynk"
	"encoding/json"
	"fmt"
	"net/http"
)

type BlynkController interface {
	SendDataToBlynk(writer http.ResponseWriter, request *http.Request)
}

type BlynkControllerImpl struct {
	blynkService blynkservice.BlynkService
}

func NewBlynkControllerImpl(blynkService blynkservice.BlynkService) BlynkController {
	return &BlynkControllerImpl{
		blynkService: blynkService,
	}
}
func (b *BlynkControllerImpl) SendDataToBlynk(writer http.ResponseWriter, request *http.Request) {
	var BlynkEsp32Request blynkpayloads.BlynkEsp32Request
	var BlynkData blynkpayloads.BlynkDataFromEsp32Request
	helper.ReadFromRequestBody(request, &BlynkEsp32Request)
	decryptedData, err := encrypthelper.DecryptAESCTR(BlynkEsp32Request.BlynkEsp32Request)
	if err != nil {
		helper.ReturnError(writer, &responses.ErrorResponses{
			Message:    "Failed to decrypt data",
			Err:        err,
			StatusCode: http.StatusInternalServerError,
			Success:    false,
		})
		return
	}

	fmt.Println("decryptedData : ", decryptedData)
	err = json.Unmarshal([]byte(decryptedData), &BlynkData)
	if err != nil {
		helper.ReturnError(writer, &responses.ErrorResponses{
			Message:    "Failed to unmarshal data",
			Err:        err,
			StatusCode: http.StatusInternalServerError,
			Success:    false,
		})
		return
	}
	err = b.blynkService.SendDataToBlynk(BlynkData)
	if err != nil {
		helper.ReturnError(writer, &responses.ErrorResponses{
			Message:    "Failed to send data to Blynk",
			Err:        err,
			StatusCode: http.StatusInternalServerError,
			Success:    false,
		})
		return
	}
	helper.HandleSuccess(writer, "success", "success to send data to blynk", http.StatusOK)

}
