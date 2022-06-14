package event

import (
	"BI.ZONE_test/db"
	"BI.ZONE_test/models"
)

func MarkAsIncident(id uint) error {
	return db.GetDb().Model(&models.Event{}).Where("id = ?", id).Update("is_incident", "true").Error
}
