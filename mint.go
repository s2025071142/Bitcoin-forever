package app

import (
	"math"
	sdk "://github.com"
)

// CalculateBlockSubsidy determines the miner reward for the current block height.
// It mirrors Bitcoin's halving mechanics but enforces a strict Monero-style tail emission floor.
func CalculateBlockSubsidy(height int64) sdk.Coin {
	initialSubsidy := float64(50 * 1000000)     // 50 coins starting block reward
	halvingInterval := int64(210000)            // Reward halves every 210,000 blocks
	tailEmissionFloor := float64(0.5 * 1000000) // 0.5 coins permanent block reward floor

	halvings := height / halvingInterval

	if halvings > 60 { 
		return sdk.NewCoin("native", sdk.NewInt(int64(tailEmissionFloor)))
	}

	currentSubsidy := initialSubsidy / math.Pow(2, float64(halvings))

	if currentSubsidy < tailEmissionFloor {
		currentSubsidy = tailEmissionFloor
	}

	return sdk.NewCoin("native", sdk.NewInt(int64(currentSubsidy)))
}
