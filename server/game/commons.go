package game

import (
	"math/rand"
	"time"
)

var randGen *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	randGen = rand.New(source)
}

func generateCode(length int) string {
	random := make([]byte, length)
	for i := range random {
		random[i] = byte(randGen.Intn(25) + 66)
	}
	return string(random)
}

func deleteIndex[DataType interface{}](slice []DataType, index int) []DataType {
	return append(slice[:index], slice[index+1:]...)
}
