package password

import (
	"crypto/sha512"
	"fmt"
	"math/rand"
	"time"
)

func Encode(password, salt string) string {
	return hash(password + salt)
}

func Verify(decoded, encoded, salt string) bool {
	return encoded == Encode(decoded, salt)
}

func Salt() string {
	b := make([]byte, 16)
	rand.Seed(time.Now().UnixNano())
	rand.Read(b)
	r := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return hash(r)
}

func hash(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
