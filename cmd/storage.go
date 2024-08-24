package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var servers = map[string]Server{}
var serverFilePath string

type Server struct {
	KeyPath string `json:"key_path"`
	IP      string `json:"ip"`
	Tag     string `json:"tag"`
	User    string `json:"user"`
}

func initStorage() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %v", err)
	}

	configDir := filepath.Join(homeDir, ".ssh_box")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.Mkdir(configDir, 0755)
	}

	serverFilePath = filepath.Join(configDir, "servers.json")

	loadServers()
}

func loadServers() {
	if _, err := os.Stat(serverFilePath); os.IsNotExist(err) {
		servers = make(map[string]Server)
		return
	}

	file, err := os.ReadFile(serverFilePath)
	if err != nil {
		log.Fatalf("Failed to read servers file: %v", err)
	}

	err = json.Unmarshal(file, &servers)
	if err != nil {
		log.Fatalf("Failed to unmarshal servers: %v", err)
	}
}

func saveServers() {
	file, err := json.MarshalIndent(servers, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal servers: %v", err)
	}

	err = os.WriteFile(serverFilePath, file, 0644)
	if err != nil {
		log.Fatalf("Failed to write servers file: %v", err)
	}
}
