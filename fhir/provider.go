package fhir

//NewFHIRPractitioner NewFHIRPractitioner
// func NewFHIRPractitioner(RedoxPayload scheduling.New) models.Practitioner {

// 	provider, err := func() (common.Provider, error) {

// 		fmt.Println("WHAT", (RedoxPayload.Visit.AttendingProvider.ID))
// 		if len(RedoxPayload.Visit.AttendingProvider.ID) > 0 {
// 			return RedoxPayload.Visit.AttendingProvider, nil
// 		}
// 		if len(RedoxPayload.Visit.ConsultingProvider.ID) > 0 {
// 			return RedoxPayload.Visit.ConsultingProvider, nil
// 		}
// 		if len(RedoxPayload.Visit.ReferringProvider.ID) > 0 {
// 			return RedoxPayload.Visit.ReferringProvider, nil
// 		}

// 		if len(RedoxPayload.Visit.VisitProvider.ID) > 0 {
// 			return RedoxPayload.Visit.VisitProvider, nil
// 		}

// 		return common.Provider{}, errors.New("No Provider")
// 	}()

// 	if err != nil {
// 		panic("Practitioner not found.")
// 	}
// 	return models.Practitioner{

// 		Identifier: []models.Identifier{
// 			{
// 				Value:  provider.ID,
// 				System: provider.IDType,
// 			},
// 		},
// 		Name: &models.HumanName{
// 			Given:  []string{provider.FirstName},
// 			Family: []string{provider.LastName},
// 		},
// 	}

// }
