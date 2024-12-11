package tmto

import "testing"

func TestPrecomputeTable(t *testing.T) {
	keyspace := uint64(256)  // Adjust keyspace for thorough testing
	keystreamLength := 16    // Keystream length
	expectedSize := keyspace // Expect this many entries

	// Call the PrecomputeTable function
	table := PrecomputeTable(keyspace, keystreamLength)

	// Check if the table has the expected number of entries
	if len(table)+1 != int(expectedSize) {
		t.Fatalf("Expected %d entries in the table, got %d", expectedSize, len(table))
	}

	// Verify keystream uniqueness
	seenKeystreams := make(map[string]struct{})
	for keystream := range table {
		if _, exists := seenKeystreams[keystream]; exists {
			t.Fatalf("Duplicate keystream found: %s", keystream)
		}
		seenKeystreams[keystream] = struct{}{}
	}
}
