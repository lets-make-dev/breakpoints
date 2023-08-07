package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getWorkspacePath() string {
	return filepath.Join(".idea", "workspace.xml")
}

func checkWorkspaceFile() bool {
	xmlPath := filepath.Join(".idea", "workspace.xml")
	if _, err := os.Stat(xmlPath); os.IsNotExist(err) {
		fmt.Println("Error: IDEA workspace.xml not found. It doesn't appear you are in a PHPStorm project root folder.")
		return false
	}
	return true
}
