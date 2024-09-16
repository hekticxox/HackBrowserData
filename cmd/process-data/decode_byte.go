package processdata

import (
    "encoding/hex"
    "fmt"
)

func main() {
    // Existing implementation
}

func hexToBytes(hexStr string) ([]byte, error) {
    return hex.DecodeString(hexStr)
}
