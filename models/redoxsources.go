package models

//RedoxSources Redox Source Table
type RedoxSources struct {
	ID       uint    `gorm:"column:id"`
	RedoxID  string  `gorm:"column:redoxId"`
	Key      string  `gorm:"column:key"`
	Secret   string  `gorm:"column:secret"`
	TargetID uint    `gorm:"column:targetId"`
	Target   Targets `gorm:"foreignkey:TargetID"`
}

func (rs RedoxSources) TableName() string {

	return "public.redoxSources"
}
