package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "encoding/hex"
    "fmt"
    "log"
)

func hexToBytes(hexStr string) ([]byte, error) {
    return hex.DecodeString(hexStr)
}

func AESGCMEncrypt(key, nonce, plaintext []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    aescm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    ciphertext := aescm.Seal(nil, nonce, plaintext, nil)
    return ciphertext, nil
}

func main() {
    // Use the key and nonce you generated
    keyHex := "73c6229fcc0f8316af12444411c939a4be864dc53f4912238cae40d787a61b0a"
    nonceHex := "503b965809fefb6fc2ec95cd"
    plaintext := "Hello, World!"

    key, err := hexToBytes(keyHex)
    if err != nil {
        log.Fatalf("Error converting key from hex: %v\n", err)
    }

    nonce, err := hexToBytes(nonceHex)
    if err != nil {
        log.Fatalf("Error converting nonce from hex: %v\n", err)
    }

    ciphertext, err := AESGCMEncrypt(key, nonce, []byte(plaintext))
    if err != nil {
        log.Fatalf("Error encrypting data: %v\n", err)
    }

    base64EncodedData := base64.StdEncoding.EncodeToString(ciphertext)
    fmt.Println("Base64 Encoded Data:", base64EncodedData)
}
