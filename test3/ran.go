package main

import (
	"crypto/rand"
	"math/big"
)

func main() {
	for i := 0; i < 4; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(100))
		println(n.Int64())
	}
}
