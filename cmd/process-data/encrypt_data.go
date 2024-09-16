package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "io"
)

// Convert hex string to bytes
func hexToBytes(hexStr string) ([]byte, error) {
    return hex.DecodeString(hexStr)
}

// Encrypt data using AES-GCM
func AESGCMEncrypt(key, nonce, plaintext []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    aead, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    ciphertext := aead.Seal(nil, nonce, plaintext, nil)
    return ciphertext, nil
}

func main() {
    // Existing implementation
}
