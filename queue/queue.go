package queue

//Fragment Fragment
type Fragment struct {
	Data       []byte
	EndpointID uint
	DataModel  string
	Event      string
}

//NSQMessage NSQMessage
type NSQMessage struct {
	EndpointIDs []uint
	Fragments   []Fragment
}
