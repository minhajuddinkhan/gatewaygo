package targets

import (
	"github.com/minhajuddinkhan/gatewaygo/queue"
)

var (
	defaultTarget   = DefaultTarget{}
	blackwellTarget = BlackwellTarget{}

	//TargetsMap TargetsMap
	targetsMap = map[string]func(dataModel, event, authParams string) Target{
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
func GetTarget(targetName, dataModel, event, authParams string) (Target, bool) {

	var target Target
	isDefault := false
	if fn, ok := targetsMap[targetName]; ok {
		target = fn(dataModel, event, authParams)
	} else {
		target = targetsMap["default"](dataModel, event, authParams)
		isDefault = true
	}
	return target, isDefault

}
