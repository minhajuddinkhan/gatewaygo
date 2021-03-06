package common

type Visit struct {
	VisitNumber        string
	AccountNumber      string
	VisitDateTime      string
	PatientClass       string
	Status             string
	Duration           string
	Reason             string
	Type               string
	AttendingProvider  Provider
	ConsultingProvider Provider
	ReferringProvider  Provider
	VisitProvider      Provider
	Location           Location
	Diagnoses          []Code

	//	Instructions []string
}
