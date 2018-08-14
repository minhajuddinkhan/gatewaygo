package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/sirupsen/logrus"

	"github.com/minhajuddinkhan/gatewaygo/constants"
	"github.com/minhajuddinkhan/gatewaygo/queue"
	"github.com/minhajuddinkhan/gatewaygo/store"
	nsq "github.com/nsqio/go-nsq"

	"github.com/darahayes/go-boom"

	_ "github.com/lib/pq"
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

//ListenerHandler ListenerHandler
func ListenerHandler(store *store.Store, producer *nsq.Producer) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if len(r.Header.Get("verification-token")) == 0 {
			boom.BadRequest(w, "no verification token")
			return
		}
		var rd models.RedoxDestination
		err := store.GetToken(r.Header.Get("verification-token")).First(&rd).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				boom.NotFound(w, err.Error())
			}
			boom.BadRequest(w, err.Error())
			return
		}

		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			boom.BadRequest(w, err.Error())
			return
		}

		var x RedoxRequestMeta
		err = json.Unmarshal(b, &x)
		if err != nil {
			boom.BadRequest(w, err.Error())
			return
		}

		var redoxSource models.RedoxSources
		if store.GetRedoxSourceBySourceID(x.Meta.Source.ID).Find(&redoxSource).RowsAffected == 0 {
			boom.NotFound(w, "Cannot find redox source")
			return
		}

		gatewayDataModel, err := mappers.GetDataModel(x.Meta.DataModel)
		if err != nil {
			boom.NotFound(w, err.Error())
			return
		}
		gatewayEvent, err := mappers.GetEvent(gatewayDataModel, x.Meta.EventType)
		if err != nil {
			boom.NotFound(w, err.Error())
			return

		}
		subscription := models.Subscriptions{
			SourceID: redoxSource.TargetID,
			Event: models.Events{
				Name: gatewayEvent,
				DataModel: models.DataModels{
					Name: gatewayDataModel,
				},
			},
		}
		err = store.GetSubscriptions(subscription.SourceID, subscription.Event.Name, subscription.Event.DataModel.Name).
			Find(&subscription).Error

		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				boom.NotFound(w, err)
				return
			}
		}
		store.GetPopulatedEndpoints(&subscription.Endpoint)

		nsqMessage := queue.NSQMessage{
			EndpointIDs: subscription.Endpoint.DependentURLIDs,
			Fragments:   []queue.Fragment{},
			Target:      subscription.Target,
			Source:      subscription.Source,
		}

		fragments, err := mappers.GetMappedFragments(subscription, b)
		if err != nil {
			boom.BadRequest(w, err.Error())
			return
		}

		orderedFragments := []queue.Fragment{}
		for _, endpointID := range nsqMessage.EndpointIDs {
			for _, nestedF := range fragments {
				if endpointID == nestedF.EndpointID {
					orderedFragments = append(orderedFragments, nestedF)
				}
			}
		}
		nsqMessage.Fragments = orderedFragments

		b, _ = json.Marshal(nsqMessage)
		err = producer.Publish(constants.TOPIC, b)
		if err != nil {
			boom.ResourceGone(w, err)
			logrus.Error("cudnt publish", err.Error())
			return
		}
		logrus.Info("MSG PUBLISHED!")
		utils.Respond(w, struct {
			Done bool
		}{
			true,
		})

	}
}
