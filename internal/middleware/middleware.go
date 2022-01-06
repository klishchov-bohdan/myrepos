package middleware

import (
	"github.com/klishchov-bohdan/logger"
	"github.com/klishchov-bohdan/myrepos/pkg/token"
)

func CheckAccessToken(access, secret string) (*token.JWTCustomClaims, error) {
	claims, err := token.ValidateToken(access, secret)
	if err != nil {
		fileLogger, _ := logger.NewFileLogger("logs/logs.txt")
		defer fileLogger.CloseFile()
		fileLogger.Error("Access token is not valid")
		return nil, err
	}
	return claims, nil
}
