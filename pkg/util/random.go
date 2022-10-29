package util

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano())
}

func RandomAmount() int64 {
	return int64(gofakeit.IntRange(200, 10000))
}

func RandomCurrency() string {
	currencies := []string{"RWF", "KSH", "UGX"}
	return currencies[gofakeit.IntRange(0, len(currencies)-1)]
}

func RandomID() int64 {
	return int64(gofakeit.IntRange(1, 1000))
}
