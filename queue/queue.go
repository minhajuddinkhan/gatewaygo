package queue

import (
	"github.com/lib/pq"
	"github.com/minhajuddinkhan/gatewaygo/models"
)

//Fragment Fragment
type Fragment struct {
	Data       []byte
	EndpointID int64
	DataModel  string
	Event      string
	Endpoint   models.Endpoints
	TargetKey  string
}

//NSQMessage NSQMessage
type NSQMessage struct {
	EndpointIDs pq.Int64Array
	Fragments   []Fragment
}
