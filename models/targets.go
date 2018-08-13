package models

import (
	"database/sql/driver"
	"encoding/json"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

type Targets struct {
	ID              uint   `gorm:"column:id"`
	Name            string `gorm:"column:name"`
	Key             string `gorm:"column:key"`
	DestinationCode string `gorm:"column:destinationCode"`
	AuthParams      JSONB  `gorm:"column:authparams" sql:"type:jsonb"`
}

func (t Targets) TableName() string {

	return "public.targets"
}
