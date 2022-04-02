package session

import "github.com/golang-jwt/jwt"

type SessionJWT struct {
	AccountOriginID int    `json:"account_origin_id"`
	Name            string `json:"name"`
	jwt.StandardClaims
}
