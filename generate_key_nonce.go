package main

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "log"
)

func main() {
    key := make([]byte, 32)  // AES-256 requires a 32-byte key
    nonce := make([]byte, 12) // AES-GCM requires a 12-byte nonce

    if _, err := rand.Read(key); err != nil {
        log.Fatalf("Error generating key: %v\n", err)
    }
    if _, err := rand.Read(nonce); err != nil {
        log.Fatalf("Error generating nonce: %v\n", err)
    }

    fmt.Printf("Generated Key: %s\n", hex.EncodeToString(key))
    fmt.Printf("Generated Nonce: %s\n", hex.EncodeToString(nonce))
}
