package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetDataDir() string {
	var dataDir string

	switch runtime.GOOS {
	case "linux":
		if xdgDataHome := os.Getenv("XDG_DATA_HOME"); xdgDataHome != "" {
			dataDir = filepath.Join(xdgDataHome, "pScan")
		}
		return filepath.Join(os.Getenv("HOME"), ".local", "share", "pScan")
	case "windows":
		appData := os.Getenv("APPDATA")
		dataDir = filepath.Join(appData, "PScan")
	case "darwin":
		home := os.Getenv("HOME")
		dataDir = filepath.Join(home, "Library", "Application Support", "PScan")
	default:
		panic("system os not supported!")
	}

	if err := os.MkdirAll(dataDir, 0700); err != nil {
		panic("could not initialize app dir!")
	}

	return dataDir
}
