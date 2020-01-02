package random_test

import (
	"gotils/random"
	"testing"
)

func TestHex(t *testing.T) {
	val := random.Hex(10)

	t.Log(val)
}
func TestRandomInteger(t *testing.T) {

	val := random.RandomInteger(10)

	t.Log(val)
}
