package auth

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserID    int  `json:"user_id"`
	IsPremium bool `json:"is_premium"`
	jwt.StandardClaims
}
