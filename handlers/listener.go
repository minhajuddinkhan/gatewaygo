package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/darahayes/go-boom"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/minhajuddinkhan/gatewaygo/auth"
	"github.com/minhajuddinkhan/gatewaygo/mappers"
	"github.com/minhajuddinkhan/gatewaygo/models"

	"github.com/minhajuddinkhan/todogo/utils"
)

type RedoxRequestMeta struct {
	Meta struct {
		DataModel string
		EventType string
		Source    struct {
			ID string
		}
	}
}

type Payload struct {
	Source    uint `json:"source"`
	DataModel string
	Event     string
}

//ListenerHandler ListenerHandler
func ListenerHandler(db *gorm.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		err := auth.VerifyToken(db, r.Header.Get("verification-token"))
		if err != nil {
			utils.Respond(w, err)
		}

		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		var x RedoxRequestMeta
		json.Unmarshal(b, &x)

		redoxSource := models.RedoxSources{
			RedoxID: x.Meta.Source.ID,
		}
		if db.First(&redoxSource).RowsAffected == 0 {
			boom.NotFound(w, "Cannot find redox source")
			return
		}
		payload := Payload{
			Source:    redoxSource.TargetID,
			DataModel: mappers.DataModelMapper[x.Meta.DataModel],
		}
		payload.Event = mappers.EventMapper[payload.DataModel][x.Meta.EventType]
		//	utils.Respond(w, redoxSource)

		subscription := models.Subscriptions{
			Source: models.Targets{
				ID: payload.Source,
			},
			Event: models.Events{
				Name: payload.Event,
			},
		}

		err = db.Preload("Source").
			Preload("Target").
			Preload("Event.DataModel").
			Preload("TargetDestination").
			Preload("Endpoint").
			Find(&subscription).Error
		if err != nil {
			boom.NotFound(w, err)
			return
		}

		dependentUrls := []uint{}
		for _, str := range strings.Split(subscription.Endpoint.DependentURLs, ",") {
			str = strings.Trim(str, "{")
			str = strings.Trim(str, "}")
			endpointID, _ := (strconv.ParseUint(str, 10, 32))
			dependentUrls = append(dependentUrls, uint(endpointID))
		}

		utils.Respond(w, subscription)
		return

		var depUrls []int
		db.Find(&depUrls)
	}
}
