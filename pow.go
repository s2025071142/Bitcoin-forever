package app

import (
	"encoding/binary"
	"errors"
	
	"github.com/opd-ai/go-randomx"
)

// VerifyRandomXProof checks if a miner's CPU work satisfies the network's current mining difficulty.
func VerifyRandomXProof(blockData []byte, nonce uint64, targetDifficulty uint64) (bool, error) {
	// 1. Convert the numeric nonce into raw bytes to append to our block data
	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, nonce)
	inputData := append(blockData, nonceBytes...)

	// 2. Initialize a secure, ephemeral RandomX Cache
	cache, err := randomx.NewCache()
	if err != nil {
		return false, errors.New("failed to initialize RandomX cache structure")
	}
	defer cache.Free() // Safely wipe memory allocations when done

	// 3. Execute the randomized CPU-intensive mathematical hashing process
	vm, err := randomx.NewVM(cache)
	if err != nil {
		return false, errors.New("failed to boot RandomX virtual machine instance")
	}
	defer vm.Free()

	// 4. Extract the finalized ASIC-resistant cryptographic hash result
	resultHash := vm.Hash(inputData)

	// 5. Check if the generated hash numeric value meets the required difficulty layer
	hashValue := binary.BigEndian.Uint64(resultHash[:8])
	if hashValue < targetDifficulty {
		return true, nil // The proof of work is fully valid!
	}

	return false, nil // The work submitted wasn't high enough
}
