package handlers

import (
	"net/http"

	"github.com/lib/pq"
	"github.com/minhajuddinkhan/todogo/utils"

	"github.com/jinzhu/gorm"
)

//Endpoint Endpoint
type Endpoint struct {
	ID              int64         `gorm:"column:id"`
	DependentURLIDs pq.Int64Array `gorm:"column:dependenturl"`
	URL             string        `gorm:"column:url"`
	DependentUrls   []Endpoint
}

func (e Endpoint) TableName() string {
	return "public.endpoints"
}

func (e *Endpoint) PopulateDependentUrls(db *gorm.DB) {

	for _, endpointID := range e.DependentURLIDs {
		populatedURL := Endpoint{ID: endpointID}
		db.Find(&populatedURL)
		e.DependentUrls = append(e.DependentUrls, populatedURL)
	}

}

//TestHandler TestHandler
func TestHandler(db *gorm.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		e := Endpoint{ID: 3}
		db.Find(&e)
		e.PopulateDependentUrls(db)
		utils.Respond(w, e)
		return
	}
}
