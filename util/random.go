package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Correct alphabet (no commas)
const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Initialize the random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)] // Pick a random letter
		sb.WriteByte(c) // Add that letter to the string
	}
	return sb.String()
}

// Generate a random username
func RandomUser() string {
	return RandomString(6)
}


//Generate a random phone
func RandomPhone() string {
	return fmt.Sprintf("+234%d", RandomInt(1000000000, 9999999999))
}

//Generate a random email
func RandomEmail() string {
	username := RandomUser()
	domain :=	"gmail"
	return fmt.Sprintf("%s@%s.com", username, domain) 
}

