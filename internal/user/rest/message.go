package rest

import (
	"go-hex/internal/user/domain"
)

type ErrorResponse struct {
	Cause string `json:"cause"`
}

type CreatedResponse struct {
	User domain.User `json:"user"`
}
