package test

import (
	"encoding/hex"
	"testing"

	"xyz/test/helloworld/encryption"
)

func TestSM3Hashing(t *testing.T) {
	// Test data
	testData := "Hello, SM3 hashing!"
	// Correct hash value for "Hello, SM3 hashing!"
	expectedHash := "6206b115368bbe42bd69f40e1a76120576f7c2232c8d61b2d6ebba6bf79c7d73"

	// Test EncodeToSM3
	hash := encryption.EncodeToSM3(testData)

	// Convert to hex string for comparison
	hashHex := hex.EncodeToString(hash)

	if hashHex != expectedHash {
		t.Errorf("SM3 hash doesn't match expected value. Expected: %s, Got: %s", expectedHash, hashHex)
	}
}

func TestSM3HashingEmptyString(t *testing.T) {
	// Test with empty string
	testData := ""
	expectedHash := "1ab21d8355cfa17f8e61194831e81a8f22bec8c728fefb747ed035eb5082aa2b"

	// Test EncodeToSM3
	hash := encryption.EncodeToSM3(testData)

	// Convert to hex string for comparison
	hashHex := hex.EncodeToString(hash)

	if hashHex != expectedHash {
		t.Errorf("SM3 hash of empty string doesn't match expected value. Expected: %s, Got: %s", expectedHash, hashHex)
	}
}

func TestSM3HashingDifferentInputs(t *testing.T) {
	// Test that different inputs produce different hashes
	input1 := "test1"
	input2 := "test2"

	hash1 := encryption.EncodeToSM3(input1)
	hash2 := encryption.EncodeToSM3(input2)

	// Convert to hex strings for comparison
	hash1Hex := hex.EncodeToString(hash1)
	hash2Hex := hex.EncodeToString(hash2)

	if hash1Hex == hash2Hex {
		t.Error("SM3 hashes of different inputs should not be equal")
	}
}
