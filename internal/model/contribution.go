package model

type Contribution struct {
	Base

	TotalUnits               float64
	EmployeeContribution     float64
	OrganisationContribution float64
	GovernmentContribution   float64
	TotalContributions       float64
	TotalSavings             float64
}
