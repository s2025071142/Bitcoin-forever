package main

import (
	"://github.com"
)

// AdjustDynamicGasLimit calculates a moving target for the EVM execution engine.
// It dynamically increases capacity during congestion and prevents network spam.
func AdjustDynamicGasLimit(historicalAverageGas uint64) uint64 {
	const TargetGasLimit uint64 = 30000000 
	const MaxGrowthFactor float64 = 2.0 

	if historicalAverageGas <= TargetGasLimit {
		return TargetGasLimit
	}

	demandRatio := float64(historicalAverageGas) / float64(TargetGasLimit)
	
	if demandRatio > MaxGrowthFactor {
		demandRatio = MaxGrowthFactor
	}

	dynamicGasLimit := uint64(float64(TargetGasLimit) * demandRatio)

	return dynamicGasLimit
}
