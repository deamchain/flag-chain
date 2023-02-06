package utils

import "math/big"

// ToUnit number of to Wei
func ToUnit(u uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(u), big.NewInt(1e18))
}
