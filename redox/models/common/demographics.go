package common

type Demographics struct {
	FirstName      string
	LastName       string
	MiddleName     string
	DOB            string
	SSN            string
	Sex            string
	Race           string
	IsHispanic     *bool
	MaritalStatus  string
	IsDeceased     *bool
	DeathDateTime  string
	PhoneNumber    Contact
	EmailAddresses []string
	Language       string
	Citizenship    []string
	Address        Address
}
