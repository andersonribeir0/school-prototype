package handlers

import (
	"errors"
	"os"

	"github.com/andersonribeir0/school-prototype/internal/models"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

var ErrInvalidBody = errors.New("invalid request body received")

type Response struct {
	Body    any    `json:"body"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

type AddUserRequest struct {
	Username string        `json:"username"`
	Password string        `json:"password"`
	Roles    []models.Role `json:"roles"`
}

type GetUserResponse struct {
	Username string        `json:"username"`
	Roles    []models.Role `json:"roles"`
}
