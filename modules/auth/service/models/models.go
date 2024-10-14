package models

import "github.com/google/uuid"

type UserCredentials struct {
	Id           uuid.UUID `json:"id" db:"id"`
	Login        string    `json:"login" db:"login"`
	PasswordHash string    `json:"password_hash" db:"password_hash"`
}
