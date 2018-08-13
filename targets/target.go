package targets

var (
	defaultTarget = DefaultTarget{}

	//TargetsMap TargetsMap
	TargetsMap = map[string]func(dataModel, event string) Target{
		"default": func(dataModel, event string) Target {
			return defaultTarget.New(dataModel, event)
		},
	}
)

//Target Target
type Target interface {
	ToFHIR(b []byte) ([]byte, error)
}
