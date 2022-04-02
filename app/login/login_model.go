package login

import "time"

type auth struct {
	Token      string    `json:"token"`
	Expiration time.Time `json:"expiration"`
}
