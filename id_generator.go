package main

import (
	"math/rand"
)

func GenerateId(length int) string {
	// DefaultABC is the default URL-friendly alphabet.
	const DefaultABC = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
	var result string = ""
	for i := 0; i < length; i++ {
		result += string(DefaultABC[rand.Intn(len(DefaultABC))])
	}
	return result
}
