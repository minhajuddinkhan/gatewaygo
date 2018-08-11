package common

type Provider struct {
	ID          string
	IDType      string
	FirstName   string
	LastName    string
	Credentials []string
	Address     Address
	PhoneNumber Contact
	Location    Location
}
