package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func mergeTextFiles(outputFile string) error {
	files, err := filepath.Glob("*.txt")
	if err != nil {
		return fmt.Errorf("failed to retrieve files: %v", err)
	}

	sortFilesByNumber(files)

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	for _, file := range files {
		fmt.Println("Processing:", file)
		if err := mergeFiles(outFile, file); err != nil {
			return err
		}
	}

	fmt.Println("Merge successful. Output file:", outputFile)
	return nil
}

func sortFilesByNumber(files []string) {
	sort.Slice(files, func(i, j int) bool {
		numI, _ := strconv.Atoi(filepath.Base(files[i][:len(files[i])-4]))
		numJ, _ := strconv.Atoi(filepath.Base(files[j][:len(files[j])-4]))
		return numI < numJ
	})
}

func mergeFiles(output io.Writer, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	_, err = io.Copy(output, file)
	return err
}

func main() {
	outputFile := "merged.txt"

	if err := mergeTextFiles(outputFile); err != nil {
		fmt.Println("Error:", err)
	}
}