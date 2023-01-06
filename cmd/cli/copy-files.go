package main

import (
	"embed"
	"errors"
	"io/ioutil"
	"os"
)

//go:embed templates
var templateFS embed.FS

func copyFileFromTemplate(templatePath, destPath string) error {
	if fileExists(destPath) {
		return errors.New(destPath + " already exists")
	}
	data, err := templateFS.ReadFile(templatePath)

	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile(data, destPath)
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

func copyDataToFile(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0644)

	if err != nil {
		return err
	}
	return nil
}

func fileExists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}

	return true
}
