package models

//RedoxSubscriptions RedoxSubscriptions
type RedoxSubscriptions struct {
	RedoxSourceID      uint `gorm:"column:redoxSource"`
	RedoxDestinationID uint `gorm:"column:redoxDestination"`
	DataModelID        uint `gorm:"column:dataModelId"`
}
