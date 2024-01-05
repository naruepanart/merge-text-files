package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io"
)

func mergeTextFiles(outputFileName string) error {
	files, err := filepath.Glob("*.txt")
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	for _, file := range files {
		inputFile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer inputFile.Close()

		_, err = io.Copy(outputFile, inputFile)
		if err != nil {
			return err
		}
	}

	fmt.Println("Merge successful Output file:", outputFileName)
	return nil
}

func main() {
	outputFileName := "merged.txt"

	err := mergeTextFiles(outputFileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
}