package web_forum

import (
	"math/rand"
)

// I should probably make sure two posts don't share the same id as that could
// break things
var charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUV1234567890@$^-_+"
var idLength = 6

func randomId() string {
	id := make([]byte, idLength)

	for i := 0; i < idLength; i++ {
		id[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(id)
}
