package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

const batchSize = 10

func mergeTextFiles(outputPrefix string) error {
	files, err := filepath.Glob("*.txt")
	if err != nil {
		return fmt.Errorf("failed to retrieve files: %v", err)
	}

	sortFilesByNumber(files)

	batchCount := 0

	for i := 0; i < len(files); i += batchSize {
		batchCount++
		outputFile := fmt.Sprintf("%s-part-%d.txt", outputPrefix, batchCount)

		outFile, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file: %v", err)
		}
		defer outFile.Close()

		fmt.Printf("Merging batch %d into %s\n", batchCount, outputFile)

		for j := i; j < i+batchSize && j < len(files); j++ {
			file := files[j]
			if err := mergeFiles(outFile, file); err != nil {
				return err
			}
		}

		fmt.Println("Merge successful. Output file:", outputFile)
	}

	return nil
}

func sortFilesByNumber(files []string) {
	sort.Slice(files, func(i, j int) bool {
		numI, _ := extractNumber(files[i])
		numJ, _ := extractNumber(files[j])
		return numI < numJ
	})
}

func extractNumber(fileName string) (int, error) {
	re := regexp.MustCompile(`(\d+)`)
	match := re.FindStringSubmatch(filepath.Base(fileName))
	if len(match) < 2 {
		return 0, fmt.Errorf("failed to extract number from filename: %s", fileName)
	}
	return strconv.Atoi(match[1])
}

func mergeFiles(output io.Writer, filePath string) error {
	fmt.Println("Processing file:", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	_, err = io.Copy(output, file)
	return err
}

func main() {
	outputPrefix := "merged"

	if err := mergeTextFiles(outputPrefix); err != nil {
		fmt.Println("Error:", err)
	}
}