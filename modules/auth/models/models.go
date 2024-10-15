package models

import "github.com/google/uuid"

type UserCredentials struct {
	Id           uuid.UUID `json:"id" db:"id"`
	Login        string    `json:"login" db:"login"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
	MfaEnabled   bool      `json:"mfa_enabled" db:"mfa_enabled"`
}

type UserPassLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Totp     string `json:"totp"`
}

type AppRoleLogin struct {
	RoleId   string `json:"role_id" db:"role_id"`
	SecretId string `json:"secret_id" db:"secret_id"`
}

type AppRoleCredentials struct {
	RoleId     string `json:"role_id" db:"role_id"`
	SecretHash string `json:"secret_hash" db:"secret_hash"`
}
