package helper

import (
	"math/rand"
	"strings"

	"github.com/shopspring/decimal"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomFloat(min, max float64) float64 {
	val := (min + rand.Float64()*(max-min)) * 100
	val, _ = decimal.NewFromFloat(val).Truncate(2).Float64()
	return val
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
