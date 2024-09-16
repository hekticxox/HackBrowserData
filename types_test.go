package types

import (
    "testing"
    "strings"
)

func TestDataType_TempFilename(t *testing.T) {
    // Example test case
    got := TempFilename()
    want := "Local Storage/leveldb_7.temp"

    if !strings.Contains(got, want) {
        t.Errorf("TempFilename() = %v, want to contain %v", got, want)
    }
}
