package utils

import (
	"github.com/Nelwhix/pScan/entity"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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

func ParseProcessOutput(output string) entity.Process {
	input := strings.Fields(output)
	// removing the headers from lsof
	trimmedInput := input[9:]
	var process entity.Process

	process.Command = trimmedInput[0]
	process.PID, _ = strconv.Atoi(trimmedInput[1])
	process.User = trimmedInput[2]
	process.FD = trimmedInput[3]
	process.Type = trimmedInput[4]
	process.Device = trimmedInput[5]
	process.Size = trimmedInput[6]
	process.Node = trimmedInput[7]
	process.Name = trimmedInput[8]

	return process
}
