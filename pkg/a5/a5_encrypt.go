package a5

func Encrypt(plaintext []byte, lfsr1, lfsr2, lfsr3 *LFSR) []byte {
	keystream := GenerateKeystream(lfsr1, lfsr2, lfsr3, len(plaintext))
	ciphertext := make([]byte, len(plaintext))
	for i := range plaintext {
		ciphertext[i] = plaintext[i] ^ keystream[i]
	}
	return ciphertext
}

func Decrypt(ciphertext []byte, lfsr1, lfsr2, lfsr3 *LFSR) []byte {
	return Encrypt(ciphertext, lfsr1, lfsr2, lfsr3)
}
