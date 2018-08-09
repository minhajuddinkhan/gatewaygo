package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/minhajuddinkhan/gatewaygo/auth"
	"github.com/minhajuddinkhan/gatewaygo/models"
	"github.com/minhajuddinkhan/todogo/utils"
)

type RedoxRequestMeta struct {
	Meta struct {
		Source struct {
			ID string
		}
	}
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
		db.First(&redoxSource)
		utils.Respond(w, redoxSource)
		return

	}
}
