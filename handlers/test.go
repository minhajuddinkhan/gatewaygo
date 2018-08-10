package handlers

import (
	"net/http"

	"github.com/minhajuddinkhan/todogo/utils"

	"github.com/lib/pq"

	"github.com/jinzhu/gorm"
)

//Endpoint Endpoint
type Endpoint struct {
	ID            uint          `gorm:"column:id"`
	DependentURLs pq.Int64Array `gorm:"column:dependenturls"`
}

//TestHandler TestHandler
func TestHandler(db *gorm.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ids := []int64{2}
		e := Endpoint{
			ID:            3,
			DependentURLs: ids,
		}

		db.Find(&e)
		utils.Respond(w, e)
		return
	}
}
