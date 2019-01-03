package diffiehellman

import (
	"crypto/rand"
	"io"
	"math/big"
)

func PrivateKey(prime *big.Int) *big.Int {
	var Reader io.Reader

	prime.Add(prime, big.NewInt(-2)) // we subtract two from the prime number to make range go from 0 to prime -2, later we add one to change range to 1 to prime -1

	key, err := rand.Int(Reader, prime)
	// read up on check error handling in go
	if err != nil {
		panic("Private key creation failed !")
	}
	// we add one to shift range to correct area
	key.Add(key, big.NewInt(1))
	return key
}

func PublicKey(privateKey *big.Int, p *big.Int, g int64) *big.Int {
	mod := new(big.Int)
	bigG := big.NewInt(g)
	return mod.Mod(bigG.Exp(bigG, privateKey, nil), p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	pr := PrivateKey(p)
	pu := PublicKey(private, p, g)
	return pr, pu
}

func SecretKey(private, public, p *big.Int) *big.Int {
	mod := new(big.Int)
	return mod.Mod(public.Exp(public, private, nil), p)

}
