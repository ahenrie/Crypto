package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ahenrie/CryptoFinal/pkg/a5"
	"github.com/ahenrie/CryptoFinal/pkg/tmto"
)

func main() {
	// Get user's name
	fmt.Println("*********************************")
	fmt.Println("* Welcome to the A5/1 TMTO *")
	fmt.Println("*********************************")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Printf("Hello, %s. We will be encrypting your name today then breaking the encryption.\n", input)

	// Perform byte conversion and insertion point for the key
	plaintext := []byte(" Hi! " + input + ". We inserted your key into the lookup table to save time.")
	insertionPoint := uint64(999999)

	// Define known parameters
	frameNumber := uint32(0x0001)
	keystreamLength := 64 // 64 bits

	// Key handling
	var key uint64
	for {
		fmt.Println("*********************************")
		fmt.Println("* Please enter a 64-bit hexadecimal key *")
		fmt.Println("* (e.g. 1234567890abcdef) *")
		fmt.Println("*********************************")
		reader2 := bufio.NewReader(os.Stdin)
		keyInput, _ := reader2.ReadString('\n')
		keyInput = strings.TrimSpace(keyInput)

		// Validate the input
		if len(keyInput) != 16 {
			fmt.Println("Invalid input. Please enter a 16-character hexadecimal string.")
			continue
		}

		// Check if the input is a valid hexadecimal string
		if !regexp.MustCompile(`^[0-9a-fA-F]+$`).MatchString(keyInput) {
			fmt.Println("Invalid input. Please enter a hexadecimal string.")
			continue
		}

		// If the input is valid, parse it to a uint64
		key, _ = strconv.ParseUint(keyInput, 16, 64)
		break
	}

	// Initialize A5/1 with the key and frame number
	lfsr1, lfsr2, lfsr3 := a5.InitializeA5_1(key, frameNumber)

	// Encrypt the plaintext
	ciphertext := a5.Encrypt(plaintext, lfsr1, lfsr2, lfsr3)
	fmt.Println("*********************************")
	fmt.Println("* Encryption Complete! *")
	fmt.Println("*********************************")
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// Precompute the table
	fmt.Println("*********************************")
	fmt.Println("* Generating Table... *")
	fmt.Println("*********************************")
	table := tmto.PrecomputeTable(1000000, keystreamLength, key, insertionPoint)

	lfsr11, lfsr21, lfsr31 := a5.InitializeA5_1(key, 0)
	keystream := a5.GenerateKeystream(lfsr11, lfsr21, lfsr31, 64)

	// Search for the key in the table
	fmt.Println("*********************************")
	fmt.Println("* Attempting to Decrypt... *")
	fmt.Println("*********************************")
	found, foundKey, foundKeystream := tmto.SearchTableByKeystream(table, keystream)
	if found {
		fmt.Println("*********************************")
		fmt.Println("* Decryption Successful! *")
		fmt.Println("*********************************")
		fmt.Printf("Found key: %x\n", foundKey)
		fmt.Printf("Found keystream: %x\n", foundKeystream)

		// Decrypt the ciphertext using the found key
		lfsr1, lfsr2, lfsr3 = a5.InitializeA5_1(foundKey, frameNumber)
		decrypted := a5.Decrypt(ciphertext, lfsr1, lfsr2, lfsr3)
		fmt.Printf("Decrypted plaintext: %s\n", decrypted)
	} else {
		fmt.Println("*********************************")
		fmt.Println("* Decryption Failed! *")
		fmt.Println("*********************************")
		fmt.Println("Keystream not found in table")
	}
}
