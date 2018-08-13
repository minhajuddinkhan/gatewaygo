package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

//DepententURLIDs DepententURLIDs
type DepententURLIDs pq.Int64Array
type PostgresJson struct {
	json.RawMessage
}

func (j PostgresJson) Value() (driver.Value, error) {
	return j.MarshalJSON()
}

func (j *PostgresJson) Scan(src interface{}) error {
	if bytes, ok := src.([]byte); ok {
		return json.Unmarshal(bytes, j)

	}
	return errors.New(fmt.Sprint("Failed to unmarshal JSON from DB", src))
}

//Endpoints Endpoints
type Endpoints struct {
	ID                int64         `gorm:"column:id"`
	URL               string        `gorm:"column:url"`
	TargetID          uint          `gorm:"column:targetId"`
	EventID           uint          `gorm:"column:eventId"`
	DataModelID       uint          `gorm:"column:dataModelId"`
	DependentURLIDs   pq.Int64Array `gorm:"column:dependentUrls"`
	VerificationToken string        `gorm:"column:verificationToken"`
	IsAuth            bool          `gorm:"column:isAuth"`
	Target            Targets       `gorm:"foreignkey:TargetID"`
	Event             Events        `gorm:"foreignkey:EventID"`
	DataModel         DataModels    `gorm:"foreignKey:DataModelID"`
	Params            string        `gorm:"column:params" sql:"json"`
	DependentURLs     []Endpoints
}

//GetPopulatedEndpoints GetPopulatedEndpoints
func (e *Endpoints) GetPopulatedEndpoints(db *gorm.DB) {

	for _, ep := range e.DependentURLIDs {
		endpoint := Endpoints{ID: ep}
		err := db.Preload("Event.DataModel").Find(&endpoint).Error
		if err != nil {
			fmt.Println("ERRRRR", err)
		}
		e.DependentURLs = append(e.DependentURLs, endpoint)

	}

	//	return e.DependentURLs
}
