package diffiehellman

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func randBigInt(max *big.Int) (*big.Int, error) {
	if max.Sign() <= 0 {
		return nil, errors.New("max must be positive")
	}

	return rand.Int(rand.Reader, max)
}

func PrivateKey(p *big.Int) *big.Int {
	t := new(big.Int)
	r, err := randBigInt(t.Sub(p, big.NewInt(2)))
	if err != nil {
		panic(err)
	}
	return r.Add(r, big.NewInt(2))
}

func PublicKey(a, p *big.Int, g int64) *big.Int {
	//convert g to bigint
	res := new(big.Int)
	res.Exp(big.NewInt(g), a, p)

	return res
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	priv := PrivateKey(p)
	pub := PublicKey(priv, p, g)
	return priv, pub
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	r := new(big.Int)
	r.Exp(public2, private1, p)
	return r
}
