package common

type Visit struct {
	VisitNumber        string
	AccountNumber      string
	VisitDateTime      string
	PatientClass       string
	Status             string
	Duration           int
	Reason             string
	Type               string
	AttendingProvider  Provider
	ConsultingProvider Provider
	ReferringProvider  Provider
	VisitProvider      Provider
	Location           Location
	Diagnoses          Diagnoses

	//	Instructions []string
}
