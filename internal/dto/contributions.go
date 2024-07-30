package dto

import (
	"github.com/google/uuid"
)

type Contributions struct {
	UserID                   uuid.UUID `json:"applicationUserId" binding:"required"`
	PersonalID               string    `json:"personalId" binding:"required"`
	FirstName                string    `json:"firstName" binding:"required"`
	LastName                 string    `json:"lastName" binding:"required"`
	TotalUnits               float64   `json:"totalUnits" binding:"required"`
	EmployeeContribution     float64   `json:"empContr" binding:"required"`
	OrganisationContribution float64   `json:"orgContr" binding:"required"`
	GovernmentContribution   float64   `json:"govtContr" binding:"required"`
	TotalContributions       float64   `json:"totalContr" binding:"required"`
	TotalSavings             float64   `json:"cummSavings" binding:"required"`
}
