package helper

import "math/rand"

const letter = "asdfghjklqwertyuiopzxcvbnm"

func GenRandomString(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
