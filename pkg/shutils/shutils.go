package shutils

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func Copy(fromPathFile, toPathFile string) error {
	sourceFile, err := os.Open(fromPathFile)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(toPathFile)
	if err != nil {
		return err
	}
	defer destFile.Close()

	n, err := io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}
	log.Println(n)
	err = destFile.Sync()
	if err != nil {
		return err
	}
	return nil
}

func HasProjectFile(folderPath string) bool {
	projectPath := filepath.Join(folderPath, "project.godot")
	if _, err := os.Stat(projectPath); err == nil {
		return true
	}
	return false
}
