package polymarket

import (
	"math/big"
)

// FormatWei converts wei to MATIC (1 MATIC = 10^18 wei)
func FormatWei(wei *big.Int) string {
	matic := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return matic.Text('f', 6)
}

// FormatGwei converts wei to Gwei (1 Gwei = 10^9 wei)
func FormatGwei(wei *big.Int) string {
	gwei := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e9))
	return gwei.Text('f', 2)
}

