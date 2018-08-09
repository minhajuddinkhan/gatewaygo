package models

type RedoxSources struct {
	ID      uint   `gorm:"column:id"`
	RedoxID string `gorm:"column:redoxId"`
}

func (rs RedoxSources) TableName() string {

	return "public.redoxSources"
}
