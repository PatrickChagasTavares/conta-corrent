package session

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type SessionJWT struct {
	AccountOriginID int    `json:"account_origin_id"`
	Name            string `json:"name"`
	jwt.StandardClaims
}

type SessionAuth struct {
	Token      string    `json:"token"`
	Expiration time.Time `json:"expiration"`
}
