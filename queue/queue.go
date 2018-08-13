package queue

import (
	"github.com/lib/pq"
)

//Fragment Fragment
type Fragment struct {
	Data       []byte
	EndpointID int64
	DataModel  string
	Event      string
}

//NSQMessage NSQMessage
type NSQMessage struct {
	EndpointIDs pq.Int64Array
	Fragments   []Fragment
}
