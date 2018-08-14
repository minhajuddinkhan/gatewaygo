package targets

import (
	"github.com/minhajuddinkhan/gatewaygo/queue"
)

var (
	defaultTarget   = DefaultTarget{}
	blackwellTarget = BlackwellTarget{}

	//TargetsMap TargetsMap
	TargetsMap = map[string]func(dataModel, event, authParams string) Target{
		"default": func(dataModel, event, authParams string) Target {
			return defaultTarget.New(dataModel, event, authParams)
		},
		"blackwell": func(dataModel, event, params string) Target {
			return blackwellTarget.New(dataModel, event, params)
		},
	}
)

//Target Target
type Target interface {
	ToFHIR(b []byte, destinationCode string) ([]byte, error)
	Execute(payload *queue.NSQMessage)
	GetAttribute(key string) (string, error)
}
