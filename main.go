package main

import (
	"fmt"
	"log"

	"github.com/ahenrie/CryptoFinal/pkg/a5"
	"github.com/ahenrie/CryptoFinal/pkg/tmto"
)

func encrypt(plaintext []byte, key uint64, keystreamLength int) []byte {
	// Initialize A5/1 with the key and frame number (0 for simplicity)
	lfsr1, lfsr2, lfsr3 := a5.InitializeA5_1(key, 0)

	// Encrypt the plaintext using A5/1's Encrypt function
	ciphertext := a5.Encrypt(plaintext, lfsr1, lfsr2, lfsr3)
	return ciphertext
}

func decrypt(ciphertext []byte, precomputedTable map[string]uint64, keystreamLength int) ([]byte, uint64) {
	// Generate the keystream from the ciphertext
	// Assuming ciphertext has been XORed with the keystream
	keystream := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i++ {
		keystream[i] = ciphertext[i] // Reverse the XOR operation for decryption
	}

	// Look up the key by the keystream in the precomputed table
	key, found := precomputedTable[string(keystream)]
	if !found {
		log.Fatal("Key not found in precomputed table!")
	}

	// Decrypt the ciphertext using the guessed key
	lfsr1, lfsr2, lfsr3 := a5.InitializeA5_1(key, 0)
	plaintext := a5.Decrypt(ciphertext, lfsr1, lfsr2, lfsr3)

	return plaintext, key
}

func main() {
	// Sample plaintext to test
	plaintext := []byte("Hello, A5/1 encryption!")

	// Choose a key (this is for encryption only, in practice it would be unknown for decryption)
	key := uint64(0x1234567890ABCDEF)

	// Define keystream length (e.g., 16 bytes)
	keystreamLength := 16

	// Encrypt the plaintext
	ciphertext := encrypt(plaintext, key, keystreamLength)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// Precompute the table (for simplicity, limit the keyspace size)
	keyspace := uint64(256) // Smaller keyspace for testing purposes
	//workers := 4           // Number of workers for parallel processing
	keystreamTable := tmto.PrecomputeTable(keyspace, keystreamLength)

	// Decrypt the ciphertext by guessing the key using the precomputed table
	decrypted, guessedKey := decrypt(ciphertext, keystreamTable, keystreamLength)

	fmt.Printf("Decrypted text: %s\n", decrypted)
	fmt.Printf("Guessed key: %x\n", guessedKey)
}
