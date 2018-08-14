package targets

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/minhajuddinkhan/gatewaygo/queue"
)

type BlackwellTarget struct {
	DataModel  string
	Event      string
	name       string
	Params     string
	ParamsJSON struct {
		Secret string
		APIKey string
	}
	DestinationCode string
}

type endpointParameters struct {
	Method string
}

type blackwellApiRequest struct {
}

func (bt *BlackwellTarget) ToFHIR(b []byte, destinationCode string) ([]byte, error) {

	return b, nil
}

func (bt *BlackwellTarget) Execute(nsqMessage *queue.NSQMessage) {

	client := http.Client{}
	for _, f := range nsqMessage.Fragments {

		var endpointParams endpointParameters
		timeStamp := strconv.Itoa(int(time.Now().UnixNano()))
		json.Unmarshal([]byte(f.Endpoint.Params), &endpointParams)
		reqData := fmt.Sprintf("%s-%s-%s", timeStamp, f.Endpoint.URL, string(f.Data))
		fmt.Println("REQD", reqData)
		h := hmac.New(sha256.New, []byte(bt.ParamsJSON.Secret))
		h.Write([]byte(reqData))
		signature := hex.EncodeToString(h.Sum(nil))

		method := "POST"
		req, err := http.NewRequest(method, f.Endpoint.URL, bytes.NewBuffer(f.Data))
		if err != nil {
			panic(err)
		}

		authorizationHeader := fmt.Sprintf("Credential=%s,Window=%s,Es_code=%s,Signature=%s", bt.ParamsJSON.APIKey, timeStamp, nsqMessage.Source.DestinationCode, signature)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", authorizationHeader)
		req.Header.Set("external-system-code", nsqMessage.Source.DestinationCode)
		fmt.Println(authorizationHeader)
		resp, _ := client.Do(req)

		b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			panic(err)
		}
		var x interface{}
		if resp.StatusCode >= 400 {
			what := json.Unmarshal(b, &x)
			if what != nil {
				logrus.Error("cannot parse response")
				return
			}
			logrus.Error(x)

			return
		}

	}

}

func (bt *BlackwellTarget) GetAttribute(key string) (string, error) {

	return "", nil
}

func (bt *BlackwellTarget) New(dataModel, event, params string) *BlackwellTarget {

	if len(params) != 0 {
		err := json.Unmarshal([]byte(params), &bt.ParamsJSON)
		if err != nil {
			panic(err.Error())
		}
	}
	bt.name = "default"
	bt.DataModel = dataModel
	bt.Event = event
	return bt

}
