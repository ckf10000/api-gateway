// Package jwt
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     jwt_service.go
* Description:  TODO
* Author:       mfkif
* CreateDate:   2024-10-19 11:22:10
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func generateExp() int64 {
	return time.Now().Add(time.Hour * TokenExp).Unix()
}

func GenerateToken(userID string) (int64, string, error) {
	exp := generateExp()
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)
	return exp, tokenStr, err
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user_id"].(string), nil
	}
	return "", err
}
