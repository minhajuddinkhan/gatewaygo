package mappers

import "errors"
import "fmt"

var (
	dataModelMapper = map[string]string{
		"Scheduling": "appointment",
	}

	eventMapper = map[string]map[string]string{
		"appointment": map[string]string{
			"New": "Schedule",
		},
	}
)

//GetDataModel GetDataModel
func GetDataModel(dataModel string) (string, error) {

	if d, ok := dataModelMapper[dataModel]; ok {
		return d, nil
	}
	return "", errors.New("DataModel not found")
}

//GetEvent GetEvent
func GetEvent(dataModel, event string) (string, error) {
	fmt.Println("HERE??", dataModel)
	if _, ok := eventMapper[dataModel]; ok {
		if e, ok := eventMapper[dataModel][event]; ok {
			return e, nil
		}
		return "", errors.New("Event not found")
	}
	return "", errors.New("DataModel not found")
}
