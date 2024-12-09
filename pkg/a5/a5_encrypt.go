package a5

// Encrypt plaintext using the A5/1 keystream
func Encrypt(plaintext []byte, lfsr1, lfsr2, lfsr3 *LFSR) []byte {
	keystream := GenerateKeystream(lfsr1, lfsr2, lfsr3, len(plaintext)*8) // Generate enough keystream bits
	ciphertext := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		// XOR the plaintext byte with the keystream byte
		ciphertext[i] = plaintext[i] ^ keystream[i]
	}
	return ciphertext
}

// Decrypt ciphertext using the A5/1 keystream
func Decrypt(ciphertext []byte, lfsr1, lfsr2, lfsr3 *LFSR) []byte {
	// Decryption is the same as encryption in XOR-based ciphers
	return Encrypt(ciphertext, lfsr1, lfsr2, lfsr3)
}
