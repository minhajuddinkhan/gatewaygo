package mappers

var (
	//DataModelMapper DataModelMapper
	DataModelMapper = map[string]string{
		"Scheduling": "appointment",
	}

	//EventMapper EventMapper
	EventMapper = map[string]map[string]string{
		"appointment": map[string]string{
			"New": "Schedule",
		},
	}
)
