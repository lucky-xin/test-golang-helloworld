package test

import (
	"bytes"
	"encoding/hex"
	"testing"
	"xyz/test/helloworld/encryption"
)

func TestSM4Encryption(t *testing.T) {
	// SM4 requires 16-byte key and 16-byte IV
	keyHex := "0123456789ABCDEFFEDCBA9876543210"
	ivHex := "00000000000000000000000000000000"

	// Create SM4 instance from hex
	sm4, err := encryption.FromHex(keyHex, ivHex)
	if err != nil {
		t.Fatalf("Failed to create SM4 instance: %v", err)
	}

	// Test data
	plaintext := "Hello, SM4 encryption!"

	// Test Encrypt and Decrypt
	ciphertext, err := sm4.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := sm4.Decrypt(ciphertext)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if string(decrypted) != plaintext {
		t.Errorf("Decrypted text doesn't match original. Expected: %s, Got: %s", plaintext, string(decrypted))
	}
}

func TestSM4HexEncryption(t *testing.T) {
	keyHex := "0123456789ABCDEFFEDCBA9876543210"
	ivHex := "00000000000000000000000000000000"

	sm4, err := encryption.FromHex(keyHex, ivHex)
	if err != nil {
		t.Fatalf("Failed to create SM4 instance: %v", err)
	}

	plaintext := "Hello, SM4 hex encryption!"

	// Test Encrypt2Hex and DecryptHex
	hexCiphertext, err := sm4.Encrypt2Hex(plaintext)
	if err != nil {
		t.Fatalf("Hex encryption failed: %v", err)
	}

	hexDecrypted, err := sm4.DecryptHex(hexCiphertext)
	if err != nil {
		t.Fatalf("Hex decryption failed: %v", err)
	}

	if string(hexDecrypted) != plaintext {
		t.Errorf("Hex decrypted text doesn't match original. Expected: %s, Got: %s", plaintext, string(hexDecrypted))
	}
}

func TestSM4Base64Encryption(t *testing.T) {
	keyHex := "0123456789ABCDEFFEDCBA9876543210"
	ivHex := "00000000000000000000000000000000"

	sm4, err := encryption.FromHex(keyHex, ivHex)
	if err != nil {
		t.Fatalf("Failed to create SM4 instance: %v", err)
	}

	plaintext := "Hello, SM4 base64 encryption!"

	// Test Encrypt2Base64 and DecryptBase64
	base64Ciphertext, err := sm4.Encrypt2Base64(plaintext)
	if err != nil {
		t.Fatalf("Base64 encryption failed: %v", err)
	}

	base64Decrypted, err := sm4.DecryptBase64(base64Ciphertext)
	if err != nil {
		t.Fatalf("Base64 decryption failed: %v", err)
	}

	if string(base64Decrypted) != plaintext {
		t.Errorf("Base64 decrypted text doesn't match original. Expected: %s, Got: %s", plaintext, string(base64Decrypted))
	}
}

func TestSM4ObjectEncryption(t *testing.T) {
	keyHex := "0123456789ABCDEFFEDCBA9876543210"
	ivHex := "00000000000000000000000000000000"

	sm4, err := encryption.FromHex(keyHex, ivHex)
	if err != nil {
		t.Fatalf("Failed to create SM4 instance: %v", err)
	}

	// Test object encryption/decryption
	type TestObject struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	testObj := TestObject{
		Name:  "test",
		Value: 123,
	}

	// Test EncryptObject
	encryptedObj, err := sm4.EncryptObject(testObj)
	if err != nil {
		t.Fatalf("Object encryption failed: %v", err)
	}

	// Test DecryptObject
	var decryptedObj TestObject
	err = sm4.DecryptObject(hex.EncodeToString(encryptedObj), &decryptedObj)
	if err != nil {
		t.Fatalf("Object decryption failed: %v", err)
	}

	if decryptedObj.Name != testObj.Name || decryptedObj.Value != testObj.Value {
		t.Errorf("Decrypted object doesn't match original. Expected: %+v, Got: %+v", testObj, decryptedObj)
	}
}

func TestSM4FromBase64(t *testing.T) {
	// Test creating SM4 instance from base64 encoded key and IV
	keyBase64 := "ASNFZ4mrze/93LqYdlQyEA=="
	ivBase64 := "AAAAAAAAAAAAAAAAAAAAAA=="

	sm4, err := encryption.FromBase64(keyBase64, ivBase64)
	if err != nil {
		t.Fatalf("Failed to create SM4 instance from base64: %v", err)
	}

	if sm4 == nil {
		t.Error("SM4 instance from base64 is nil")
	}
}

func TestSM4NewSM4(t *testing.T) {
	// Test creating SM4 instance directly from byte arrays
	key, _ := hex.DecodeString("0123456789ABCDEFFEDCBA9876543210")
	iv, _ := hex.DecodeString("00000000000000000000000000000000")

	sm4, err := encryption.NewSM4(key, iv)
	if err != nil {
		t.Fatalf("Failed to create SM4 instance: %v", err)
	}

	if sm4 == nil {
		t.Error("SM4 instance is nil")
	}
}

// Test padding and unpadding functions
func TestPaddingUnpadding(t *testing.T) {
	// Test data
	plaintext := []byte("Hello, SM4 encryption!")
	blockSize := 16

	// Test padding
	paddedData := encryption.PaddingLastGroup(plaintext, blockSize)
	if len(paddedData)%blockSize != 0 {
		t.Errorf("Padded data length is not multiple of block size. Expected: %d, Got: %d",
			len(paddedData)%blockSize, 0)
	}

	// Test unpadding
	unpaddedData := encryption.UnpaddingLastGroup(paddedData)
	if !bytes.Equal(unpaddedData, plaintext) {
		t.Errorf("Unpadded data doesn't match original. Expected: %s, Got: %s",
			string(plaintext), string(unpaddedData))
	}
}

// TestUnpaddingLastGroupEdgeCases tests edge cases for UnpaddingLastGroup function
func TestUnpaddingLastGroupEdgeCases(t *testing.T) {
	// Test with empty byte slice
	emptyData := []byte{}
	result := encryption.UnpaddingLastGroup(emptyData)
	if !bytes.Equal(result, emptyData) {
		t.Errorf("UnpaddingLastGroup should return empty slice for empty input. Expected: %v, Got: %v",
			emptyData, result)
	}

	// Test with data that has valid padding
	// Create data with 5 bytes of padding (value 0x05)
	dataWithPadding := []byte("test data")
	for i := 0; i < 5; i++ {
		dataWithPadding = append(dataWithPadding, 0x05)
	}
	expected := []byte("test data")
	result = encryption.UnpaddingLastGroup(dataWithPadding)
	if !bytes.Equal(result, expected) {
		t.Errorf("UnpaddingLastGroup failed for valid padding. Expected: %v, Got: %v",
			expected, result)
	}

	// Test with data that has invalid padding (padding value too large)
	dataWithInvalidPadding := []byte("test data")
	for i := 0; i < 20; i++ { // 20 is larger than the data length
		dataWithInvalidPadding = append(dataWithInvalidPadding, 0x14) // 0x14 = 20
	}
	result = encryption.UnpaddingLastGroup(dataWithInvalidPadding)
	if !bytes.Equal(result, dataWithInvalidPadding) {
		t.Errorf("UnpaddingLastGroup should return original data for invalid padding. Expected: %v, Got: %v",
			dataWithInvalidPadding, result)
	}

	// Test with data that has invalid padding (padding value is zero)
	dataWithZeroPadding := []byte("test data")
	dataWithZeroPadding = append(dataWithZeroPadding, 0x00)
	result = encryption.UnpaddingLastGroup(dataWithZeroPadding)
	if !bytes.Equal(result, dataWithZeroPadding) {
		t.Errorf("UnpaddingLastGroup should return original data for zero padding. Expected: %v, Got: %v",
			dataWithZeroPadding, result)
	}
}

// Test error handling
func TestSM4ErrorHandling(t *testing.T) {
	// Test with invalid hex key
	_, err := encryption.FromHex("invalid", "00000000000000000000000000000000")
	if err == nil {
		t.Error("Expected error for invalid hex key, but got nil")
	}

	// Test with invalid hex IV
	_, err = encryption.FromHex("0123456789ABCDEFFEDCBA9876543210", "invalid")
	if err == nil {
		t.Error("Expected error for invalid hex IV, but got nil")
	}

	// Test with invalid base64 key
	_, err = encryption.FromBase64("invalid!", "AAAAAAAAAAAAAAAAAAAAAA==")
	if err == nil {
		t.Error("Expected error for invalid base64 key, but got nil")
	}

	// Test with invalid base64 IV
	_, err = encryption.FromBase64("ASNFZ4mrze/93LqYdlQyEA==", "invalid!")
	if err == nil {
		t.Error("Expected error for invalid base64 IV, but got nil")
	}
}

// TestSM4DecryptErrorHandling tests error handling in SM4 decrypt functions
func TestSM4DecryptErrorHandling(t *testing.T) {
	keyHex := "0123456789ABCDEFFEDCBA9876543210"
	ivHex := "00000000000000000000000000000000"

	sm4, err := encryption.FromHex(keyHex, ivHex)
	if err != nil {
		t.Fatalf("Failed to create SM4 instance: %v", err)
	}

	// Test DecryptHex with invalid hex string
	_, err = sm4.DecryptHex("invalid")
	if err == nil {
		t.Error("Expected error for invalid hex string in DecryptHex, but got nil")
	}

	// Test DecryptBase64 with invalid base64 string
	_, err = sm4.DecryptBase64("invalid!")
	if err == nil {
		t.Error("Expected error for invalid base64 string in DecryptBase64, but got nil")
	}
}

// Test edge cases and boundary conditions
func TestEdgeCases(t *testing.T) {
	keyHex := "0123456789ABCDEFFEDCBA9876543210"
	ivHex := "00000000000000000000000000000000"

	sm4, err := encryption.FromHex(keyHex, ivHex)
	if err != nil {
		t.Fatalf("Failed to create SM4 instance: %v", err)
	}

	// Test with empty string
	plaintext := ""
	ciphertext, err := sm4.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Encryption failed for empty string: %v", err)
	}

	decrypted, err := sm4.Decrypt(ciphertext)
	if err != nil {
		t.Fatalf("Decryption failed for empty string: %v", err)
	}

	if string(decrypted) != plaintext {
		t.Errorf("Decrypted text doesn't match original for empty string. Expected: %s, Got: %s", plaintext, string(decrypted))
	}

	// Test with exact block size multiple
	plaintext = "0123456789ABCDEF" // 16 bytes
	ciphertext, err = sm4.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Encryption failed for exact block size: %v", err)
	}

	decrypted, err = sm4.Decrypt(ciphertext)
	if err != nil {
		t.Fatalf("Decryption failed for exact block size: %v", err)
	}

	if string(decrypted) != plaintext {
		t.Errorf("Decrypted text doesn't match original for exact block size. Expected: %s, Got: %s", plaintext, string(decrypted))
	}
}

// Test DecryptObject error handling
func TestDecryptObjectErrorHandling(t *testing.T) {
	keyHex := "0123456789ABCDEFFEDCBA9876543210"
	ivHex := "00000000000000000000000000000000"

	sm4, err := encryption.FromHex(keyHex, ivHex)
	if err != nil {
		t.Fatalf("Failed to create SM4 instance: %v", err)
	}

	// Test with invalid hex string
	var obj map[string]interface{}
	err = sm4.DecryptObject("invalid", &obj)
	if err == nil {
		t.Error("Expected error for invalid hex string, but got nil")
	}

	// Test with invalid JSON - need to encrypt valid hex first
	// Create a SM4 instance for encryption
	key, _ := hex.DecodeString(keyHex)
	iv, _ := hex.DecodeString(ivHex)
	sm4Enc, _ := encryption.NewSM4(key, iv)

	// Encrypt some invalid JSON data
	invalidJSON := []byte("{ invalid json }")
	encryptedInvalidJSON, err := sm4Enc.Encrypt(string(invalidJSON))
	if err != nil {
		t.Fatalf("Failed to encrypt invalid JSON: %v", err)
	}

	// Now test decrypting it
	err = sm4.DecryptObject(hex.EncodeToString(encryptedInvalidJSON), &obj)
	if err == nil {
		t.Error("Expected error for invalid JSON, but got nil")
	}
}
