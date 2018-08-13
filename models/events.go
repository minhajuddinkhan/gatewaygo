package models

type Events struct {
	ID          uint       `gorm:"column:id"`
	Name        string     `gorm:"column:name"`
	DataModelID uint       `gorm:"column:dataModelId"`
	DataModel   DataModels `gorm:"foreignkey:DataModelID"`
}

func (e Events) TableName() string {
	return "public.events"
}
