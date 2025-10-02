package main

import (
	"fmt"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
)

func main() {
	// Example: Initialize a G1 point
	bls12381.Generators()

	fmt.Println("Project initialized with BLS12-381 G1 generator")
}
