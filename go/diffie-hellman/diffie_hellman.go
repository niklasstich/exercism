package diffiehellman

import (
	"math/big"
	"math/rand"
)

func PrivateKey(p *big.Int) *big.Int {
	return big.NewInt(rand.Int63())
}

func PublicKey(b, p, g *big.Int) *big.Int {
	pubkey := g.Exp(g, b, nil)
	pubkey = g.Mod(pubkey, p)
	return pubkey
}
