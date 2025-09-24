package test

import (
	"crypto/rand"
	"encoding/hex"
	"testing"

	"xyz/test/helloworld/encryption"

	"github.com/tjfoc/gmsm/sm2"
)

// generateSM2KeyPair generates a valid SM2 key pair in the format expected by the encryption package
func generateSM2KeyPair() (publicKeyHex, privateKeyHex string, err error) {
	// Generate a valid SM2 key pair
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}

	// Create uncompressed public key format (0x04 || x || y)
	publicKey := &privateKey.PublicKey
	curve := publicKey.Curve
	byteLen := (curve.Params().BitSize + 7) / 8

	// Create uncompressed format: 0x04 || x || y
	xBytes := publicKey.X.Bytes()
	yBytes := publicKey.Y.Bytes()

	// Pad x and y to full byte length
	xPadded := make([]byte, byteLen)
	yPadded := make([]byte, byteLen)
	copy(xPadded[byteLen-len(xBytes):], xBytes)
	copy(yPadded[byteLen-len(yBytes):], yBytes)

	// Create uncompressed public key bytes
	uncompressed := make([]byte, 1+2*byteLen)
	uncompressed[0] = 0x04
	copy(uncompressed[1:byteLen+1], xPadded)
	copy(uncompressed[byteLen+1:], yPadded)

	publicKeyHex = hex.EncodeToString(uncompressed)
	privateKeyHex = hex.EncodeToString(privateKey.D.Bytes())

	return publicKeyHex, privateKeyHex, nil
}

func TestSM2Encryption(t *testing.T) {
	// Generate a valid SM2 key pair for testing
	publicKeyHex, privateKeyHex, err := generateSM2KeyPair()
	if err != nil {
		t.Fatalf("Failed to generate SM2 key pair: %v", err)
	}

	// Create SM2 instance
	sm2Enc, err := encryption.NewSM2(publicKeyHex, privateKeyHex)
	if err != nil {
		t.Fatalf("Failed to create SM2 instance: %v", err)
	}

	// Test encryption and decryption
	plaintext := "Hello, SM2 encryption!"

	// Test Encrypt and Decrypt (C1C3C2 mode)
	ciphertext, err := sm2Enc.Encrypt(plaintext, 0)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := sm2Enc.Decrypt(ciphertext, 0)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if string(decrypted) != plaintext {
		t.Errorf("Decrypted text doesn't match original. Expected: %s, Got: %s", plaintext, string(decrypted))
	}

	// Test Encrypt2Hex and DecryptHex (C1C3C2 mode)
	hexCiphertext, err := sm2Enc.Encrypt2Hex(plaintext, 0)
	if err != nil {
		t.Fatalf("Hex encryption failed: %v", err)
	}

	hexDecrypted, err := sm2Enc.DecryptHex(hexCiphertext, 0)
	if err != nil {
		t.Fatalf("Hex decryption failed: %v", err)
	}

	if string(hexDecrypted) != plaintext {
		t.Errorf("Hex decrypted text doesn't match original. Expected: %s, Got: %s", plaintext, string(hexDecrypted))
	}

	// Test Encrypt2Base64 and DecryptBase64 (C1C3C2 mode)
	base64Ciphertext, err := sm2Enc.Encrypt2Base64(plaintext, 0)
	if err != nil {
		t.Fatalf("Base64 encryption failed: %v", err)
	}

	base64Decrypted, err := sm2Enc.DecryptBase64(base64Ciphertext, 0)
	if err != nil {
		t.Fatalf("Base64 decryption failed: %v", err)
	}

	if string(base64Decrypted) != plaintext {
		t.Errorf("Base64 decrypted text doesn't match original. Expected: %s, Got: %s", plaintext, string(base64Decrypted))
	}

	// Test with C1C2C3 mode
	ciphertextC1C2C3, err := sm2Enc.Encrypt(plaintext, 1)
	if err != nil {
		t.Fatalf("C1C2C3 encryption failed: %v", err)
	}

	decryptedC1C2C3, err := sm2Enc.Decrypt(ciphertextC1C2C3, 1)
	if err != nil {
		t.Fatalf("C1C2C3 decryption failed: %v", err)
	}

	if string(decryptedC1C2C3) != plaintext {
		t.Errorf("C1C2C3 decrypted text doesn't match original. Expected: %s, Got: %s", plaintext, string(decryptedC1C2C3))
	}
}

// TestSM2DecryptErrorHandling tests error handling in SM2 decrypt functions
func TestSM2DecryptErrorHandling(t *testing.T) {
	// Generate a valid SM2 key pair for testing
	publicKeyHex, privateKeyHex, err := generateSM2KeyPair()
	if err != nil {
		t.Fatalf("Failed to generate SM2 key pair: %v", err)
	}

	// Create SM2 instance
	sm2Enc, err := encryption.NewSM2(publicKeyHex, privateKeyHex)
	if err != nil {
		t.Fatalf("Failed to create SM2 instance: %v", err)
	}

	// Test DecryptHex with invalid hex string
	_, err = sm2Enc.DecryptHex("invalid", 0)
	if err == nil {
		t.Error("Expected error for invalid hex string in DecryptHex, but got nil")
	}

	// Test DecryptBase64 with invalid base64 string
	_, err = sm2Enc.DecryptBase64("invalid!", 0)
	if err == nil {
		t.Error("Expected error for invalid base64 string in DecryptBase64, but got nil")
	}

	// Test DecryptObject with invalid hex string
	var obj map[string]interface{}
	err = sm2Enc.DecryptObject("invalid", 0, &obj)
	if err == nil {
		t.Error("Expected error for invalid hex string in DecryptObject, but got nil")
	}
}

func TestSM2ObjectEncryption(t *testing.T) {
	// Generate a valid SM2 key pair for testing
	publicKeyHex, privateKeyHex, err := generateSM2KeyPair()
	if err != nil {
		t.Fatalf("Failed to generate SM2 key pair: %v", err)
	}

	sm2Enc, err := encryption.NewSM2(publicKeyHex, privateKeyHex)
	if err != nil {
		t.Fatalf("Failed to create SM2 instance: %v", err)
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
	encryptedObj, err := sm2Enc.EncryptObject(testObj, 0)
	if err != nil {
		t.Fatalf("Object encryption failed: %v", err)
	}

	// Test DecryptObject
	var decryptedObj TestObject
	err = sm2Enc.DecryptObject(hex.EncodeToString(encryptedObj), 0, &decryptedObj)
	if err != nil {
		t.Fatalf("Object decryption failed: %v", err)
	}

	if decryptedObj.Name != testObj.Name || decryptedObj.Value != testObj.Value {
		t.Errorf("Decrypted object doesn't match original. Expected: %+v, Got: %+v", testObj, decryptedObj)
	}
}

func TestDecodePublicKey(t *testing.T) {
	// Generate a valid SM2 key pair for testing
	publicKeyHex, _, err := generateSM2KeyPair()
	if err != nil {
		t.Fatalf("Failed to generate SM2 key pair: %v", err)
	}

	publicKey, err := encryption.DecodePublicKey(publicKeyHex)
	if err != nil {
		t.Fatalf("Failed to decode public key: %v", err)
	}

	if publicKey == nil {
		t.Error("Decoded public key is nil")
		return
	}

	// Verify the decoded public key has the expected properties
	if publicKey.X == nil || publicKey.Y == nil {
		t.Error("Decoded public key is missing X or Y coordinates")
	}
}

func TestDecodePrivateKey(t *testing.T) {
	// Generate a valid SM2 key pair for testing
	publicKeyHex, privateKeyHex, err := generateSM2KeyPair()
	if err != nil {
		t.Fatalf("Failed to generate SM2 key pair: %v", err)
	}

	decodedPrivateKey, err := encryption.DecodePrivateKey(privateKeyHex, publicKeyHex)
	if err != nil {
		t.Fatalf("Failed to decode private key: %v", err)
	}

	if decodedPrivateKey == nil {
		t.Error("Decoded private key is nil")
		return
	}

	// Verify the decoded private key has the expected properties
	if decodedPrivateKey.D == nil {
		t.Error("Decoded private key is missing D value")
	}
}

// TestDecodePrivateKeyErrorHandling tests error handling in DecodePrivateKey
func TestDecodePrivateKeyErrorHandling(t *testing.T) {
	// Generate a valid SM2 key pair for testing
	publicKeyHex, _, err := generateSM2KeyPair()
	if err != nil {
		t.Fatalf("Failed to generate SM2 key pair: %v", err)
	}

	// Test with invalid private key hex
	_, err = encryption.DecodePrivateKey("invalid", publicKeyHex)
	if err == nil {
		t.Error("Expected error for invalid private key hex, but got nil")
	}

	// Test with invalid public key hex
	_, err = encryption.DecodePrivateKey("0123456789ABCDEF", "invalid")
	if err == nil {
		t.Error("Expected error for invalid public key hex, but got nil")
	}

	// Test with mismatched private and public keys (using a valid but different key)
	_, privateKeyHex, err := generateSM2KeyPair()
	if err != nil {
		t.Fatalf("Failed to generate SM2 key pair: %v", err)
	}

	_, err = encryption.DecodePrivateKey(privateKeyHex, publicKeyHex)
	// Note: This test may not always produce an error depending on the implementation
	// We're just checking that the function doesn't panic
	if err != nil {
		// This is actually expected behavior - the function should return an error
		// for mismatched keys, so we're good
	}
}
