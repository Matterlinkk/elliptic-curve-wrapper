package test_main

import (
	"awesomeProject/wrappers"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"testing"
)

func SetRandom(bits int) *big.Int {
	randomInt, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), uint(bits)))
	if err != nil {
		panic(err)
	}
	return randomInt
}

func TestCorrectness(t *testing.T) {

	curve := elliptic.P256()

	ecWrapper := wrappers.NewECWrapper(curve)

	G := ecWrapper.GetPointG()

	k := SetRandom(256)
	d := SetRandom(256)

	H1 := ecWrapper.ScalarMult(d, G)

	H2 := ecWrapper.ScalarMult(k, H1)

	H3 := ecWrapper.ScalarMult(k, G)

	H4 := ecWrapper.ScalarMult(d, H3)

	h2x, h2y := H2.Params()
	h4x, h4y := H4.Params()

	if !H2.IsEqualTo(H4) {
		t.Errorf("Invalid result: x-values: %s, %s, y-values: %s, %s", h2x, h2y, h4x, h4y)
	}
}
