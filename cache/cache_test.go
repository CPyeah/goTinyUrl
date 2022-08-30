package cache

import (
	"testing"
)

func init() {
	Init()
}

func TestAdd(t *testing.T) {
	Add("aba")
}

func TestExists(t *testing.T) {
	println(Exists("aba"))
	println(Exists("abc"))
}
