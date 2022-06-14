package event

import (
	"BI.ZONE_test/db"
	"BI.ZONE_test/models"
	"BI.ZONE_test/utils"
	"encoding/base64"
)

func SaveEvent(event models.Event) error {
	event.Message = base64.StdEncoding.EncodeToString(utils.Encode([]byte(event.Message)))
	return db.GetDb().Create(&event).Error
}
