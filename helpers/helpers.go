package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func Min(x, y int) int {
	if x <= y {
		return x
	}

	return y
}

func GenerateNumberOfLength(num_len int) string {
	rand.Seed(time.Now().UnixNano())

	min := 1
	for i := 0; i < num_len-1; i++ {
		min *= 10
	}

	max := 1
	for i := 0; i < num_len; i++ {
		max *= 10
	}

	return fmt.Sprintf("%d", rand.Intn(max-min)+min)
}
