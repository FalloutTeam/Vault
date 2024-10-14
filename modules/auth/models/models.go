package models

import "github.com/google/uuid"

type UserCredentials struct {
	Id           uuid.UUID `json:"id" db:"id"`
	Login        string    `json:"login" db:"login"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	MfaEnabled   bool      `json:"mfa_enabled" db:"mfa_enabled"`
}

type UserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Totp     string `json:"totp"`
}
