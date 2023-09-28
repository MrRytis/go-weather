package model

import (
	_ "github.com/gookit/validate"
	"time"
)

type AuthRequest struct {
	Email    string `json:"email" validate:"required|email"`
	Password string `json:"password" validate:"required|minLen:8"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
