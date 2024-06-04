package helperfunctions

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	
}
