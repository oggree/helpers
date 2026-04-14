package helpers

import (
	"math/rand/v2"
)

// InArray checks if a specific value exists within a slice of comparable items.
// It returns true and the first index if found, or false and -1 if not found.
// The function uses Go Generics, providing type safety and better performance than reflection.
func InArray[T comparable](val T, array []T) (exists bool, index int) {
	for i, v := range array {
		if v == val {
			return true, i
		}
	}
	return false, -1
}

// RandomInt generates a pseudo-random integer in the half-open interval [min, max).
// It panics if min >= max, consistent with the standard library behavior.
func RandomInt(min, max int) int {
	if min >= max {
		panic("helpers.RandomInt: min must be strictly less than max")
	}
	return min + rand.IntN(max-min)
}

// RandomString generates a random string of the given length consisting only of uppercase A-Z characters.
func RandomString(length int) string {
	if length <= 0 {
		return ""
	}
	
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		// ASCII 65 is 'A', 90 is 'Z'.
		// We pass max as 91 because RandomInt output bounds are [min, max).
		bytes[i] = byte(RandomInt(65, 91))
	}
	return string(bytes)
}