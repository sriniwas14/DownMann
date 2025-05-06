package configloader

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var DestFolder string
var TempDir string
var MaxThreads int64 = 4

type Config struct {
	MaxThreads int64 `json:"maxThreads"`
}

var AppConfig = Config{
	MaxThreads: 4, // default value
}

func getUserCacheDir() string {
	home := os.Getenv("HOME")
	switch runtime.GOOS {
	case "windows":
		if localAppData := os.Getenv("LOCALAPPDATA"); localAppData != "" {
			return filepath.Join(localAppData, "downmann", "Cache")
		}
	case "darwin":
		return filepath.Join(home, "Library", "Caches", "downmann")
	default:
		if xdg := os.Getenv("XDG_CACHE_HOME"); xdg != "" {
			return filepath.Join(xdg, "downmann")
		}
		return filepath.Join(home, ".cache", "downmann")
	}
	return filepath.Join(home, ".cache", "downmann")
}

func getUserConfigDir() string {
	home := os.Getenv("HOME")
	switch runtime.GOOS {
	case "windows":
		if appData := os.Getenv("APPDATA"); appData != "" {
			return filepath.Join(appData, "downmann")
		}
	case "darwin":
		return filepath.Join(home, "Library", "Application Support", "downmann")
	default:
		if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
			return filepath.Join(xdg, "downmann")
		}
		return filepath.Join(home, ".config", "downmann")
	}
	return filepath.Join(home, ".config", "downmann")
}

func ensureConfigExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Creating default config at %s\n", path)
		data, err := json.MarshalIndent(AppConfig, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal default config: %v", err)
		}
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			log.Fatalf("Failed to create config dir: %v", err)
		}
		if err := os.WriteFile(path, data, 0644); err != nil {
			log.Fatalf("Failed to write default config: %v", err)
		}
	}
}

func loadConfig(path string) {
	ensureConfigExists(path)
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Could not read config file: %v\n", err)
		return
	}
	err = json.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Printf("Error parsing config file: %v\n", err)
	}
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	TempDir = getUserCacheDir()
	DestFolder = filepath.Join(homeDir, "Downloads")

	if err := os.MkdirAll(TempDir, 0755); err != nil {
		log.Fatalf("Failed to create temp dir: %v", err)
	}

	configPath := filepath.Join(getUserConfigDir(), "downmann.json")
	loadConfig(configPath)

	log.Println("Home Directory:", homeDir)
	log.Println("Cache Directory:", TempDir)
	log.Println("Destination Folder:", DestFolder)
	log.Printf("Loaded Config: %+v\n", AppConfig)
}
