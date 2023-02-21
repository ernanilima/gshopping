package test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func RootDir() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)

	rootDir := exPath
	for {
		if _, err := os.Stat(filepath.Join(rootDir, "main.go")); !os.IsNotExist(err) {
			break
		}
		rootDir = filepath.Dir(rootDir)
		fmt.Println("ERRR")
	}

	return rootDir
}
