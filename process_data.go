package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "encoding/hex"
    "flag"
    "fmt"
    "log"
    "os"
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

func AESGCMDecrypt(key, nonce, ciphertext []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    aescm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    plaintext, err := aescm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}

func main() {
    flag.Parse()
    if flag.NArg() < 1 {
        fmt.Println("Usage: process_data <action> <data>")
        fmt.Println("Action: encrypt or decrypt")
        os.Exit(1)
    }

    action := flag.Arg(0)
    data := flag.Arg(1)

    keyHex := "73c6229fcc0f8316af12444411c939a4be864dc53f4912238cae40d787a61b0a"
    nonceHex := "503b965809fefb6fc2ec95cd"

    key, err := hexToBytes(keyHex)
    if err != nil {
        log.Fatalf("Error converting key from hex: %v\n", err)
    }

    nonce, err := hexToBytes(nonceHex)
    if err != nil {
        log.Fatalf("Error converting nonce from hex: %v\n", err)
    }

    if action == "encrypt" {
        plaintext := data
        ciphertext, err := AESGCMEncrypt(key, nonce, []byte(plaintext))
        if err != nil {
            log.Fatalf("Error encrypting data: %v\n", err)
        }

        base64EncodedData := base64.StdEncoding.EncodeToString(ciphertext)
        fmt.Println("Base64 Encoded Data:", base64EncodedData)
    } else if action == "decrypt" {
        ciphertext, err := base64.StdEncoding.DecodeString(data)
        if err != nil {
            log.Fatalf("Error decoding base64 data: %v\n", err)
        }

        decryptedData, err := AESGCMDecrypt(key, nonce, ciphertext)
        if err != nil {
            log.Fatalf("Error decrypting data: %v\n", err)
        }

        fmt.Printf("Decrypted data: %s\n", decryptedData)
    } else {
        fmt.Println("Unknown action. Use 'encrypt' or 'decrypt'.")
        os.Exit(1)
    }
}
