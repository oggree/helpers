package helpers

import (
	"regexp"
	"strings"
	"testing"
)

func TestInArray(t *testing.T) {
	// Test strings
	strSlice := []string{"apple", "banana", "cherry"}
	
	if ok, idx := InArray("banana", strSlice); !ok || idx != 1 {
		t.Errorf("InArray strings: expected true and 1, got %v and %v", ok, idx)
	}
	
	if ok, idx := InArray("mango", strSlice); ok || idx != -1 {
		t.Errorf("InArray strings: expected false and -1, got %v and %v", ok, idx)
	}

	// Test integers
	intSlice := []int{10, 20, 30, 40}
	
	if ok, idx := InArray(30, intSlice); !ok || idx != 2 {
		t.Errorf("InArray ints: expected true and 2, got %v and %v", ok, idx)
	}
	
	if ok, idx := InArray(50, intSlice); ok || idx != -1 {
		t.Errorf("InArray ints: expected false and -1, got %v and %v", ok, idx)
	}
}

func TestRandomInt(t *testing.T) {
	min, max := 5, 10
	
	for i := 0; i < 100; i++ {
		val := RandomInt(min, max)
		if val < min || val >= max {
			t.Errorf("RandomInt: generated value %d is out of bounds [%d, %d)", val, min, max)
		}
	}
}

func TestRandomIntPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("RandomInt did not panic when min >= max")
		}
	}()
	RandomInt(10, 10)
}

func TestRandomString(t *testing.T) {
	sizes := []int{1, 5, 10, 50, 100}
	
	// Regex that guarantees A-Z only
	isValidAlphabet := regexp.MustCompile(`^[A-Z]*$`).MatchString

	for _, size := range sizes {
		str := RandomString(size)
		
		if len(str) != size {
			t.Errorf("RandomString length mismatch: expected %d, got %d", size, len(str))
		}
		
		if !isValidAlphabet(str) {
			t.Errorf("RandomString contains invalid characters: %q", str)
		}
	}
	
	// Ensure that Z is actually achievable at some point.
	// Since 65-90 is 26 characters, probability of not getting a 'Z' in 1000 characters is almost 0.
	foundZ := false
	largeStr := RandomString(1000)
	if strings.Contains(largeStr, "Z") {
		foundZ = true
	}
	if !foundZ {
		t.Errorf("RandomString generated 1000 chars but did not generate a single 'Z'. Off by one error?")
	}
}

func TestRandomStringEmpty(t *testing.T) {
	str := RandomString(0)
	if str != "" {
		t.Errorf("Expected empty string for length 0, got %q", str)
	}
	
	strNegative := RandomString(-5)
	if strNegative != "" {
		t.Errorf("Expected empty string for length < 0, got %q", strNegative)
	}
}
