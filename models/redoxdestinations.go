package models

//RedoxDestination RedoxDestination
type RedoxDestination struct {
	VerificationToken   string             `gorm:"column:verificationToken"`
	RedoxID             string             `gorm:"column:redoxId"`
	ID                  uint               `gorm:"column:id"`
	TargetDestinationID uint               `gorm:"column:targetDestId"`
	TargetDestination   TargetDestinations `gorm:"foreignkey:id"`
}

func (redoxDest RedoxDestination) TableName() string {
	return "public.redoxDestinations"
}
