package tmto

import (
	"fmt"

	"github.com/ahenrie/CryptoFinal/pkg/a5"
)

func PrecomputeTable(keyspace uint64, keystreamLength int, knownKey uint64, insertionPoint uint64) map[string]uint64 {
	finalTable := make(map[string]uint64)

	for key := uint64(0); key < keyspace; key++ {
		lfsr1, lfsr2, lfsr3 := a5.InitializeA5_1(key, 0)
		keystream := a5.GenerateKeystream(lfsr1, lfsr2, lfsr3, keystreamLength)

		// Store the keystream as raw bytes (binary)
		keystreamBytes := string(keystream) // or `keystreamBytes := keystream` if using []byte
		finalTable[keystreamBytes] = key

		// Progress tracking
		if key%100000 == 0 {
			fmt.Printf("Progress: Generated %d entries\n", key)
		}

		// Plant the known key after reaching the insertion point
		if key == insertionPoint {
			fmt.Println("This is taking too long, lets implant your key!")
			fmt.Printf("Planted key at %d: %x\n", insertionPoint, knownKey)
			lfsr1, lfsr2, lfsr3 = a5.InitializeA5_1(knownKey, 0)
			knownKeystream := a5.GenerateKeystream(lfsr1, lfsr2, lfsr3, keystreamLength)
			knownKeystreamBytes := string(knownKeystream) // or `knownKeystreamBytes := knownKeystream` for []byte
			finalTable[knownKeystreamBytes] = knownKey
		}
	}

	fmt.Printf("Generated table with %d unique entries.\n", len(finalTable))
	return finalTable
}

func PrintTable(table map[string]uint64) {
	for keystream, key := range table {
		// Convert the keystream back to a readable format, if needed (e.g., hex encoding)
		fmt.Printf("Keystream: %x, Key: %x\n", []byte(keystream), key)
	}
}

func SearchTableByKeystream(table map[string]uint64, keystream []byte) (bool, uint64, []byte) {
	if table == nil {
		panic("table is nil")
	}
	if len(keystream) == 0 {
		panic("keystream is empty")
	}
	keystreamStr := string(keystream)
	if guessedKey, exists := table[keystreamStr]; exists {
		return true, guessedKey, keystream
	}
	return false, 0, nil
}
