package targets

import (
	"github.com/minhajuddinkhan/gatewaygo/queue"
	"github.com/minhajuddinkhan/gatewaygo/targets"
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

//GetTarget GetTarget
func GetTarget(targetName, dataModel, event, authParams string) Target {

	var target Target
	if fn, ok := TargetsMap[targetName]; ok {
		target = fn(dataModel, event, authParams)
	} else {
		target = targets.TargetsMap["default"](dataModel, event, authParams)
	}
	return target

}
