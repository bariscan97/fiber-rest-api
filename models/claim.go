package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claim struct {	
	User  FetchUserModel
	jwt.RegisteredClaims
}
