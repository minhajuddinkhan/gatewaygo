package store

import "github.com/jinzhu/gorm"

//GetSubscriptions GetSubscriptions
func (s *Store) GetSubscriptions(sourceID uint, eventName, dataModelName string) *gorm.DB {

	return s.Database.Preload("Source").
		Preload("Target").
		Preload("Event.DataModel").
		Preload("TargetDestination").
		Preload("Endpoint").
		Joins(`left join events e on "eventId" = e.id`).
		Joins(`left join "dataModels" d on ("dataModelId" = d.id)`).
		Where(`"sourceId" = ? and e."name" = ? and d."name" = ?`, sourceID, eventName, dataModelName)

}
