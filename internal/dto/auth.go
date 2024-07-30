package dto

import "github.com/google/uuid"

type AuthBody struct {
	Username     string `json:"username" binding:"required"`
	PasswordOne  string `json:"passwordOne" binding:"required"`
	PasswordTwo  string `json:"passwordTwo"`
	LanguageCode string `json:"languageCode"`
}

type AuthResponse struct {
	IsValid          bool      `json:"isValid,omitempty"`
	IsEmployee       bool      `json:"isEmployee,omitempty"`
	IsFirstLogined   bool      `json:"isFirstLogined,omitempty"`
	HasManyCompanies bool      `json:"hasManyCompanies,omitempty"`
	Message          string    `json:"message,omitempty"`
	AccessToken      string    `json:"accessToken,omitempty"`
	RefreshToken     string    `json:"refreshToken,omitempty"`
	UserID           uuid.UUID `json:"userId,omitempty"`
}
