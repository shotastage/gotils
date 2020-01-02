package random_test

import (
	"gotils/random"
	"testing"
)

func TestRandomNum(t *testing.T) {

	val := random.RandomInteger(10)

	println(val)
}
