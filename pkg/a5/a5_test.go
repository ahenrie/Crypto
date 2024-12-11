package a5

import (
	"fmt"
	"testing"
)

// TestLFSRClock tests the LFSR's clocking mechanism.
func TestLFSRClock(t *testing.T) {
	lfsr := &LFSR{
		state:    0b1001010101010101010, // Example state
		mask:     0b1100000000000011001,
		size:     19,
		clockBit: 8,
	}

	initialState := lfsr.state
	lfsr.Clock(true) // Clock with the clockBit set
	if lfsr.state == initialState {
		t.Errorf("LFSR did not clock correctly: got %b, want something different", lfsr.state)
	}
}

// TestClockingBit tests the correct extraction of the clocking bit.
func TestClockingBit(t *testing.T) {
	lfsr := &LFSR{
		state:    0b1000000000000000001, // MSB is 1
		mask:     0b1100000000000011001,
		size:     19,
		clockBit: 1, // Clocking bit at position 1
	}

	if !lfsr.ClockingBit() {
		t.Errorf("Expected clocking bit to be true, got false")
	}

	lfsr.state = 0b0100000000000000000 // MSB is now 0
	if lfsr.ClockingBit() {
		t.Errorf("Expected clocking bit to be false, got true")
	}
}

// TestParity tests the parity function.
func TestParity(t *testing.T) {
	tests := []struct {
		value    uint32
		expected uint32
	}{
		{0b0, 0},
		{0b1, 1},
		{0b11, 0},
		{0b111, 1},
		{0b101, 0},   // Odd number of 1s
		{0b110, 0},   // Even number of 1s
		{0b1111, 0},  // Even number of 1s
		{0b10001, 0}, // Odd number of 1s
	}

	for _, test := range tests {
		result := parity(test.value)
		if result != test.expected {
			t.Errorf("parity(%b) = %d, want %d", test.value, result, test.expected)
		}
	}
}

// TestInitializeA5_1 tests the initialization of the A5/1 registers.
func TestInitializeA5_1(t *testing.T) {
	key := uint64(0b1010101010101010101010101010101010101010101010101010101010101010)
	frameNumber := uint32(0b110110)

	lfsr1, lfsr2, lfsr3 := InitializeA5_1(key, frameNumber)

	if lfsr1.state == 0 || lfsr2.state == 0 || lfsr3.state == 0 {
		t.Errorf("One or more LFSRs were not initialized correctly")
	}
}

// TestMajorityVote tests the majority vote function.
func TestMajorityVote(t *testing.T) {
	tests := []struct {
		a, b, c bool
		want    bool
	}{
		{true, true, true, true},
		{true, true, false, true},
		{true, false, false, false},
		{false, false, false, false},
		{true, false, true, true},
	}

	for _, test := range tests {
		result := majorityVote(test.a, test.b, test.c)
		if result != test.want {
			t.Errorf("majorityVote(%v, %v, %v) = %v, want %v", test.a, test.b, test.c, result, test.want)
		}
	}
}

// TestGenerateKeystream tests the keystream generation.
func TestGenerateKeystream(t *testing.T) {
	key := uint64(0b1010101010101010101010101010101010101010101010101010101010101010)
	frameNumber := uint32(0b110110)

	lfsr1, lfsr2, lfsr3 := InitializeA5_1(key, frameNumber)

	keystream := GenerateKeystream(lfsr1, lfsr2, lfsr3, 10)
	if len(keystream) != 10 {
		t.Errorf("Expected keystream length of 10, got %d", len(keystream))
	}

	// Check that keystream bits are valid (0 or 1)
	for i, bit := range keystream {
		if bit != 0 && bit != 1 {
			t.Errorf("Invalid keystream bit at position %d: got %d, want 0 or 1", i, bit)
		}
	}
}

func TestEncryptionDecryption(t *testing.T) {
	key := uint64(0x1234567890ABCDEF)
	frameNumber := uint32(0x0001)
	plaintext := []byte("Hello, A5/1 encryption!")

	// Encrypt
	lfsr1, lfsr2, lfsr3 := InitializeA5_1(key, frameNumber)
	ciphertext := Encrypt(plaintext, lfsr1, lfsr2, lfsr3)

	// Re-initialize LFSRs for decryption
	lfsr1, lfsr2, lfsr3 = InitializeA5_1(key, frameNumber)
	decrypted := Decrypt(ciphertext, lfsr1, lfsr2, lfsr3)

	if string(decrypted) == string(plaintext) {
		fmt.Printf("Congrats the plaintext: %s was encrypted: %x then decrypted: %s\n", plaintext, ciphertext, decrypted)
	}

	if string(decrypted) != string(plaintext) {
		t.Errorf("Decrypted text does not match original. Got: %s, Want: %s", decrypted, plaintext)
	}
}
