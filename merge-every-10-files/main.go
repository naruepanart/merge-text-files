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

func main() {
	outputPrefix := "merged"

	// Retrieve list of text files in the current directory
	files, err := filepath.Glob("*.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Sort files by their numeric part in the filename
	sort.Slice(files, func(i, j int) bool {
		numI, _ := extractNumber(files[i])
		numJ, _ := extractNumber(files[j])
		return numI < numJ
	})

	// Merge files in batches
	batchCount := 0
	for i := 0; i < len(files); i += batchSize {
		batchCount++
		outputFile := fmt.Sprintf("%s-part-%d.txt", outputPrefix, batchCount)

		// Create output file
		outFile, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer outFile.Close()

		fmt.Printf("Merging batch %d into %s\n", batchCount, outputFile)

		// Merge files in the current batch
		for j := i; j < i+batchSize && j < len(files); j++ {
			filePath := files[j]
			fmt.Println("Processing file:", filePath)

			// Open input file
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer file.Close()

			// Copy file content to output file
			_, err = io.Copy(outFile, file)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		fmt.Println("Merge successful. Output file:", outputFile)
	}
}

// extractNumber extracts the numeric part from a filename.
func extractNumber(fileName string) (int, error) {
	match := regexp.MustCompile(`(\d+)`).FindStringSubmatch(filepath.Base(fileName))
	if len(match) < 2 {
		return 0, fmt.Errorf("failed to extract number from filename: %s", fileName)
	}
	return strconv.Atoi(match[1])
}