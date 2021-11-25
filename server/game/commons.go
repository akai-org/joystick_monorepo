package game

import "math/rand"

func generateCode(length int) string {
	random := make([]byte, length)
	for i := range random {
		random[i] = byte(rand.Intn(25) + 66)
	}
	return string(random)
}
