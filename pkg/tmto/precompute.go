package tmto

import (
	"fmt"

	"github.com/ahenrie/CryptoFinal/pkg/a5"
)

func PrecomputeTable(keyspace uint64, keystreamLength int) map[string]uint64 {
	// Initialize a map to store the results
	finalTable := make(map[string]uint64)

	// Iterate over all keys in the keyspace
	for key := uint64(0); key < keyspace; key++ {
		// Initialize A5/1 for the current key
		lfsr1, lfsr2, lfsr3 := a5.InitializeA5_1(key, 0)

		// Generate the keystream for the current key
		keystream := a5.GenerateKeystream(lfsr1, lfsr2, lfsr3, keystreamLength)

		// Store the mapping of keystream to key in the final table
		finalTable[string(keystream)] = key
	}

	// Optionally print the table for debugging
	fmt.Printf("Generated table with %d entries.\n", len(finalTable))

	return finalTable
}
