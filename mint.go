package app

import (
	"math"
	sdk "://github.com"
)

// CalculateBlockSubsidy determines the miner reward for the current block height.
// It decreases the reward gradually with every single block until hitting the tail floor.
func CalculateBlockSubsidy(height int64) sdk.Coin {
	initialSubsidy := float64(50 * 1000000)       // 50 coins starting block reward (in base units)
	tailEmissionFloor := float64(0.005 * 1000000) // 0.005 coins permanent floor

	// This decay factor controls how fast the reward shrinks. 
	// 0.000005 means the reward drops by a tiny fraction every single block.
	decayFactor := 0.000005 

	// Calculate the smooth gradual drop based on the current block height
	currentSubsidy := initialSubsidy * math.Exp(-decayFactor * float64(height))

	// If the smooth calculation drops below your floor, lock it at 0.005 forever!
	if currentSubsidy < tailEmissionFloor {
		currentSubsidy = tailEmissionFloor
	}

	return sdk.NewCoin("native", sdk.NewInt(int64(currentSubsidy)))
}
