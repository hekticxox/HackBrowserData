package main

import (
    "crypto/rand"
    "fmt"
)

func generateKey() ([]byte, error) {
    key := make([]byte, 32) // AES-256
    _, err := rand.Read(key)
    if err != nil {
        return nil, err
    }
    return key, nil
}

func generateNonce() ([]byte, error) {
    nonce := make([]byte, 12) // GCM standard nonce size
    _, err := rand.Read(nonce)
    if err != nil {
        return nil, err
    }
    return nonce, nil
}

func main() {
    // Existing implementation
}
