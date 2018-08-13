package models

type DataModels struct {
	ID   uint   `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (d DataModels) TableName() string {
	return "public.dataModels"
}
