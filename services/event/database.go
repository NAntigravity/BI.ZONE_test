package event

import (
	"BI.ZONE_test/db"
	"BI.ZONE_test/models"
	"BI.ZONE_test/utils"
	"encoding/base64"
	"fmt"
)

func GetDecodedData(id uint) (models.Event, error) {
	eventFromDB := models.Event{ID: id}
	err := db.GetDb().First(&eventFromDB).Error
	if err != nil {
		return models.Event{}, err
	}
	msg, err := base64.StdEncoding.DecodeString(eventFromDB.Message)
	if err != nil {
		return models.Event{}, err
	}
	eventFromDB.Message = string(utils.Decode(msg))
	return eventFromDB, nil
}

func GetEventsList(sortingParam *string, orderParam *string, start int, end int, filterParam *string, filterBy *string) ([]models.EventWithoutMessage, error) {
	var allEvents []models.Event
	err := db.GetDb().Where(utils.PrepareWhereExpression(filterParam, filterBy)).Order(fmt.Sprintf("%s %s", *sortingParam, *orderParam)).Offset(start).Limit(end - start).Find(&allEvents).Error
	if err != nil {
		return []models.EventWithoutMessage{}, err
	}
	var result []models.EventWithoutMessage
	for _, event := range allEvents {
		result = append(result, event.RemoveMessage())
	}
	return result, err
}

func GetEventsAmount(filterParam *string, filterBy *string) (int64, error) {
	var count int64
	err := db.GetDb().Model(&models.Event{}).Where(utils.PrepareWhereExpression(filterParam, filterBy)).Count(&count).Error
	return count, err
}
