package crypto

import (
    "crypto/aes"
    "crypto/cipher"
    "errors"
)

// AESGCMEncrypt encrypts the plaintext using AES-GCM with the provided key and nonce.
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

// AESGCMDecrypt decrypts the given ciphertext using AES-GCM with the provided key and nonce.
func AESGCMDecrypt(key, nonce, ciphertext []byte) ([]byte, error) {
    if len(nonce) != 12 {
        return nil, errors.New("invalid nonce size, must be 12 bytes")
    }

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
