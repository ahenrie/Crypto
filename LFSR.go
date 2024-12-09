package main

type LFSR struct {
	state    uint32
	mask     uint32
	size     int
	majority int
}

// Clock the LFSR
func (l *LFSR) Clock(clockBit bool) {
	if clockBit {
		feedback := parity(l.state & l.mask)
		l.state = ((l.state << 1) | feedback) & ((1 << l.size) - 1)
	}
}

func (l *LFSR) ClockingBit() bool {
	return (l.state>>(l.size-l.majority))&1 == 1
}

func parity(val uint32) uint32 {
	var result uint32
	for val > 0 {
		result ^= val & 1
		val >>= 1
	}
	return result
}
