package main

import (
	"fmt"

	"github.com/ahenrie/CryptoFinal/pkg/tmto"
)

func main() {
	// Define the keyspace, keystream length, and number of workers for precomputation
	keyspace := uint64(1000000000) // Adjust based on your keyspace size
	keystreamLength := 16          // Length of the keystream to be generated

	// Precompute the table of keystreams and corresponding keys
	precomputedTable := tmto.PrecomputeTable(keyspace, keystreamLength)

	// Example: assume we have a known keystream (from encryption or attack)
	// In a real scenario, you would extract this from intercepted ciphertext
	knownKeystream := []byte{
		0x1A, 0x2B, 0x3C, 0x4D, 0x5E, 0x6F, 0x7A, 0x8B,
		0x9C, 0xAD, 0xBE, 0xCF, 0xD0, 0xE1, 0xF2, 0x03,
	}

	// Perform the TMTO attack
	key, found := tmto.TMTOAttack(precomputedTable, knownKeystream)

	if found {
		fmt.Printf("Key found: %d\n", key)
	} else {
		fmt.Println("Keystream not found in precomputed table.")
	}
}
