package model

import (
	"github.com/google/uuid"
)

type Device struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Host        string    `json:"host"`
	Port        int       `json:"port"`
}
