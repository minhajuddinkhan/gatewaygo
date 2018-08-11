package common

type Location struct {
	Type       string
	Facility   string
	Department string
	Room       string
}

type Contact struct {
	Home   string
	Office string
	Mobile string
}

type Address struct {
	StreetAddress string
	City          string
	State         string
	ZIP           string
	County        string
	Country       string
}

type Diagnoses struct {
	Code        string
	Codeset     string
	Description string
	Value       string
}

type Source struct {
	ID   string
	Name string
}

type Destination struct {
	ID   string
	Name string
}

type Message struct {
	ID uint
}
type Transmission struct {
	ID uint
}

type Identifier struct {
	ID     string
	IDType string
}
type Meta struct {
	DataModel     string
	EventType     string
	EventDateTime string
	Test          bool
	Source        Source
	Transmission  Transmission
	FacilityCode  string
}
