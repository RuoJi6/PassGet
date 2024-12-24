package navicat

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/blowfish"
)

func xor(a, b []byte) []byte {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = a[i] ^ b[i]
	}
	return result
}

func decrypt(encrypted string) (string, error) {
	// Compute SHA1 hash of the key
	key := sha1.Sum([]byte("3DC5CA39"))

	// Initialize Blowfish cipher with the hashed key
	cipher, err := blowfish.NewCipher(key[:])
	if err != nil {
		return "", fmt.Errorf("failed to create Blowfish cipher: %w", err)
	}

	// Generate the IV by encrypting a block of all 0xFF bytes
	iv := make([]byte, blowfish.BlockSize)
	for i := range iv {
		iv[i] = 0xFF
	}
	cipher.Encrypt(iv, iv)

	// Decode the ciphertext from hex
	ciphertext, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex string: %w", err)
	}

	// Decrypt the ciphertext
	cv := iv
	plaintext := []byte{}
	fullRounds := len(ciphertext) / blowfish.BlockSize
	leftLength := len(ciphertext) % blowfish.BlockSize

	// Process full blocks
	for i := 0; i < fullRounds; i++ {
		block := ciphertext[i*blowfish.BlockSize : (i+1)*blowfish.BlockSize]
		decrypted := make([]byte, blowfish.BlockSize)
		cipher.Decrypt(decrypted, block)
		decrypted = xor(decrypted, cv)
		plaintext = append(plaintext, decrypted...)
		cv = xor(cv, block)
	}

	// Process any remaining bytes
	if leftLength > 0 {
		cvEncrypted := make([]byte, blowfish.BlockSize)
		cipher.Encrypt(cvEncrypted, cv)
		plaintext = append(plaintext, xor(ciphertext[fullRounds*blowfish.BlockSize:], cvEncrypted[:leftLength])...)
	}

	// Return the plaintext as a string
	return string(plaintext), nil
}
