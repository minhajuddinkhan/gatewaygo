package models

type Targets struct {
	ID              uint   `gorm:"column:id"`
	Name            string `gorm:"column:name"`
	Key             string `gorm:"column:key"`
	DestinationCode string `gorm:"column:destinationCode"`
	AuthParams      string `gorm:"column:authparams" sql:"json"`
}

func (t Targets) TableName() string {

	return "public.targets"
}
