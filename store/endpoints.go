package store

import (
	"fmt"

	"github.com/minhajuddinkhan/gatewaygo/models"
)

//GetPopulatedEndpoints GetPopulatedEndpoints
func (s *Store) GetPopulatedEndpoints(e *models.Endpoints) {

	for _, ep := range e.DependentURLIDs {
		endpoint := models.Endpoints{ID: ep}
		err := s.Database.Preload("Event.DataModel").Find(&endpoint).Error
		if err != nil {
			fmt.Println("ERRRRR", err)
		}
		e.DependentURLs = append(e.DependentURLs, endpoint)

	}

}
