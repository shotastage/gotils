package random

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomText(letter int) string {
	return ""
}

func RandomInteger(digit int) int {
	rand.Seed(time.Now().Unix())

	for i := 1; i <= 10; i++ {
		fmt.Println(rand.Intn(10))
	}

	return 10
}
