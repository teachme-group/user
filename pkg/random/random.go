package random

import (
	"fmt"
	"math/rand"
)

func NewConfirmationCode() string {
	code := rand.Intn(900000) + 100000

	return fmt.Sprintf("%06d", code)
}
