package models

import (
	"github.com/google/uuid"
)


type FetchUserModel struct {
	Id       uuid.UUID 
	Username string    
	Email    string    
	Password string
}

type RegisterUserModel struct {
	Username string    `json:"username" validate:"required,min=5,max=20"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8,max=24"`
}

type LoginReqBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}

type UpdateUserModel struct {
	Username string `json:"username" validate:"omitempty,min=5,max=20"`
}
