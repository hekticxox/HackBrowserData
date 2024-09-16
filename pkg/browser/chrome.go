package browser

import (
    "os"
    "path/filepath"
)

// GetChromeDataDir returns the path to the Chrome User Data directory
func GetChromeDataDir() string {
    // Use the environment variable for LocalAppData
    userProfileDir := filepath.Join(os.Getenv("LOCALAPPDATA"), "Google", "Chrome", "User Data")

    // Return the User Data directory
    return userProfileDir
}
