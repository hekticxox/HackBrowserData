package main

import (
    "encoding/base64"
    "encoding/hex"
    "flag"
    "fmt"
    "log"
    "os"

    "github.com/moond4rk/hackbrowserdata/crypto"
)

func hexToBytes(hexStr string) ([]byte, error) {
    return hex.DecodeString(hexStr)
}

func main() {
    flag.Parse()
    if flag.NArg() < 1 {
        fmt.Println("Usage: hack-browser-data <base64-encoded-data>")
        os.Exit(1)
    }

    base64EncodedData := flag.Arg(0)

    // Use the key and nonce you generated
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

    // Decode the base64-encoded data
    encodedData, err := base64.StdEncoding.DecodeString(base64EncodedData)
    if err != nil {
        log.Fatalf("Error decoding base64 data: %v\n", err)
    }

    // Decrypt the data
    decryptedData, err := crypto.AESGCMDecrypt(key, nonce, encodedData)
    if err != nil {
        log.Fatalf("Error decrypting data: %v\n", err)
    }

    // Print the decrypted data
    fmt.Printf("Decrypted data: %s\n", decryptedData)
}
