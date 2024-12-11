package a5

import "math/bits"

type LFSR struct {
	state    uint32
	mask     uint32
	size     int
	clockBit int // Position of the clocking bit
}

// Clock the LFSR with irregular clocking
func (l *LFSR) Clock(clockBit bool) {
	if clockBit {
		feedback := parity(l.state & l.mask)
		// Shift the state and insert the feedback bit at the LSB
		l.state = ((l.state << 1) | feedback) & ((1 << l.size) - 1)
	}
}

// Get the clocking bit
func (l *LFSR) ClockingBit() bool {
	return (l.state>>(l.size-l.clockBit))&1 == 1
}

func parity(value uint32) uint32 {
	return uint32(bits.OnesCount32(value) % 2)
}

func InitializeA5_1(key uint64, frameNumber uint32) (*LFSR, *LFSR, *LFSR) {
	lfsr1 := &LFSR{state: 0, mask: 0x072000, size: 19, clockBit: 8}
	lfsr2 := &LFSR{state: 0, mask: 0x300000, size: 22, clockBit: 10}
	lfsr3 := &LFSR{state: 0, mask: 0x700080, size: 23, clockBit: 10}

	// Load key into LFSRs
	for i := 0; i < 64; i++ {
		bit := (key >> (63 - i)) & 1
		lfsr1.state ^= uint32(bit) << (i % lfsr1.size)
		lfsr2.state ^= uint32(bit) << (i % lfsr2.size)
		lfsr3.state ^= uint32(bit) << (i % lfsr3.size)
	}

	// Load frame number into LFSRs
	for i := 0; i < 22; i++ {
		bit := (frameNumber >> (21 - i)) & 1
		lfsr1.state ^= uint32(bit) << (i % lfsr1.size)
		lfsr2.state ^= uint32(bit) << (i % lfsr2.size)
		lfsr3.state ^= uint32(bit) << (i % lfsr3.size)
	}

	// Clock LFSRs 100 times after loading key and frame number (standard A5/1 initialization step)
	for i := 0; i < 100; i++ {
		lfsr1.Clock(true)
		lfsr2.Clock(true)
		lfsr3.Clock(true)
	}

	return lfsr1, lfsr2, lfsr3
}

func GenerateKeystream(lfsr1, lfsr2, lfsr3 *LFSR, length int) []uint8 {
	keystream := make([]uint8, length)
	for i := 0; i < length; i++ {
		m := majorityVote(lfsr1.ClockingBit(), lfsr2.ClockingBit(), lfsr3.ClockingBit())
		lfsr1.Clock(lfsr1.ClockingBit() == m)
		lfsr2.Clock(lfsr2.ClockingBit() == m)
		lfsr3.Clock(lfsr3.ClockingBit() == m)
		keystream[i] = uint8((lfsr1.state & 1) ^ (lfsr2.state & 1) ^ (lfsr3.state & 1))
	}
	return keystream
}

func majorityVote(a, b, c bool) bool {
	return (a && b) || (b && c) || (a && c)
}
