package aoc2020

import "testing"

func TestIfCanGenerateHandshakeEncryptionKey(t *testing.T) {
	ctx := NewRfidContext(5764801, 17807724)
	encKey := ctx.CalculateEncryptionKey()

	if encKey != 14897079 {
		t.Errorf("Expected 14897079, got %v\n", encKey)
	}
}
