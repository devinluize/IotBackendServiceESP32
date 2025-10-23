package BlynkServiceImpl

import (
	blynkpayloads "IotBackend/api/payloads/blynk"
	blynkrepositoy "IotBackend/api/repositories/blynk"
	blynkservice "IotBackend/api/service/blynk"
	"gorm.io/gorm"
)

type blynkServiceImpl struct {
	db              *gorm.DB
	blynkRepository blynkrepositoy.BlynkRepository
}

func NewBlynkServiceImpl(db *gorm.DB, blynkRepository blynkrepositoy.BlynkRepository) blynkservice.BlynkService {
	return &blynkServiceImpl{
		db:              db,
		blynkRepository: blynkRepository,
	}
}

func (b *blynkServiceImpl) SendDataToBlynk(request blynkpayloads.BlynkDataFromEsp32Request) error {
	//tx := b.db.Begin()
	//defer helper.CommitOrRollback(tx)
	err := b.blynkRepository.SendDataToBlynk(b.db, request)
	if err != nil {
		return err
	}
	return nil
}
