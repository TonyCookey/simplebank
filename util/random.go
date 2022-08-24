package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {

	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomAccountName generates a rnadom account name with first and last names
func RandomAccountName() string {
	return strings.Title(strings.ToLower(RandomString(int(RandomInt(5, 7))) + " " + RandomString(int(RandomInt(5, 7)))))
}

// RandomAmount generates a random amount
func RandomAmount() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency
func RandomCurrency() string {
	currencies := []string{"EUR", "NGN", "USD", "GBP", "CAD"}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}
