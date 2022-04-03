package password

import (
	"crypto/sha512"
	"fmt"
	"math/rand"
	"time"
)

type Password interface {
	Encode(password, salt string) string
	Verify(decoded, encoded, salt string) bool
	Salt() string
}

type passwordImpl struct {
}

func NewPassword() Password {
	password := &passwordImpl{}

	return password
}

func (p *passwordImpl) Encode(password, salt string) string {
	return p.hash(password + salt)
}

func (p *passwordImpl) Verify(decoded, encoded, salt string) bool {
	return encoded == p.Encode(decoded, salt)
}

func (p *passwordImpl) Salt() string {
	b := make([]byte, 16)
	rand.Seed(time.Now().UnixNano())
	rand.Read(b)
	r := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return p.hash(r)
}

func (p *passwordImpl) hash(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
