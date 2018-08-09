package models

type TargetDestinations struct {
	ID         uint   `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	PracticeID string `gorm:"column:practiceId"`
}

func (targetDest TargetDestinations) TableName() string {
	return "public.targetDestinations"
}
