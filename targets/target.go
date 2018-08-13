package targets

//Target Target
type Target interface {
	toFHIR(interface{}) interface{}
}
