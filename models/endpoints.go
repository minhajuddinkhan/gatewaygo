package models

type Endpoints struct {
	ID            uint   `gorm:"column:id"`
	URL           string `gorm:"column:url"`
	TargetID      uint   `gorm:"column:targetId"`
	EventID       uint   `gorm:"column:eventId"`
	DependentURLs string `gorm:"column:dependentUrls;type:integer[]"`
	//	Params            JSONB  `gorm:"column:params" sql:"type:jsonb"`
	VerificationToken string `gorm:"column:verificationToken"`
	IsAuth            bool   `gorm:"column:isAuth"`

	Target Targets `gorm:"foreignkey:TargetID"`
	Event  Events  `gorm:"foreignkey:EventID"`
}
