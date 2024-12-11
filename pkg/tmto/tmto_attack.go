package tmto

// TMTOAttack performs the Time-Memory Trade-Off attack
// It takes the precomputed table and a known keystream and returns the corresponding key if found
func TMTOAttack(precomputedTable map[string]uint64, knownKeystream []byte) (uint64, bool) {
	// Convert the keystream into a string format to check in the precomputed table
	keystreamStr := string(knownKeystream)

	// Check if the keystream exists in the precomputed table
	key, found := precomputedTable[keystreamStr]

	return key, found
}
