package crypto

import "errors"

// ErrCiphertextLengthIsInvalid is an error indicating an invalid length of ciphertext.
var ErrCiphertextLengthIsInvalid = errors.New("ciphertext length is invalid")
