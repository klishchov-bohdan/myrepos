package login

import "github.com/golang-jwt/jwt"

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	Token *jwt.Token
}
