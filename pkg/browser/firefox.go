package browser

import (
    "os"
    "path/filepath"
)

// GetFirefoxDataDir returns the path to the Firefox Profiles directory
func GetFirefoxDataDir() string {
    // Use the environment variable for AppData
    userProfileDir := filepath.Join(os.Getenv("APPDATA"), "Mozilla", "Firefox", "Profiles")

    // Return the Profiles directory
    return userProfileDir
}
