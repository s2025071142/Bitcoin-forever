package main

import (
	"encoding/binary"
	"errors"

	"github.com/opd-ai/go-randomx"
)

// VerifyRandomXProof checks if a miner's CPU work satisfies the network's current mining difficulty.
func VerifyRandomXProof(blockData []byte, nonce uint64, targetDifficulty uint64, seedHash []byte) (bool, error) {
	// 1. Convert the numeric nonce into raw bytes to append to our block data
	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, nonce)
	inputData := append(blockData, nonceBytes...)

	// 2. Configure and initialize the pure-Go RandomX hashing instance
	config := randomx.Config{
		Mode:     randomx.FastMode, // 2GB FastMode optimization for verification speed
		CacheKey: seedHash,
	}

	hasher, err := randomx.New(config)
	if err != nil {
		return false, errors.New("failed to initialize RandomX proof verification engine")
	}
	defer hasher.Close() // Safely wipe pooled resources and RAM allocations when done

	// 3. Extract the finalized ASIC-resistant cryptographic hash result
	resultHash := hasher.Hash(inputData)

	// 4. Check if the generated hash numeric value meets the required difficulty layer
	hashValue := binary.BigEndian.Uint64(resultHash[:8])
	if hashValue < targetDifficulty {
		return true, nil // The proof of work is fully valid!
	}

	return false, nil // The work submitted wasn't high enough
}
