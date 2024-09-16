package types

import (
    "path/filepath"
    "os"
    "testing"
)

func TestDataType_TempFilename(t *testing.T) {
    tempDir := filepath.Join(os.TempDir(), "Local Storage")
    tempFile := filepath.Join(tempDir, "leveldb_7.temp")

    // Check if the file path contains the correct prefix
    if !containsPrefix(tempFile, "Local Storage/leveldb_7.temp") {
        t.Errorf("Error: %s does not contain %s", tempFile, "Local Storage/leveldb_7.temp")
    }
}

func containsPrefix(filePath, prefix string) bool {
    // Function to check if filePath contains the prefix
    return filepath.HasPrefix(filePath, prefix)
}
