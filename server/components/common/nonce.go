package common

import (
	"math/rand"
	"time"
)

func RandomNonce() int {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000)
	return code
}
