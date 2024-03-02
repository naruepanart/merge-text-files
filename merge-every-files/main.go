package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

const bufferSize = 4096 // Adjust buffer size as needed

func main() {
	outputFile := "merged.txt"

	files, err := filepath.Glob("*.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sort.Slice(files, func(i, j int) bool {
		numI, _ := extractNumber(files[i])
		numJ, _ := extractNumber(files[j])
		return numI < numJ
	})

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outFile.Close()

	for _, file := range files {
		fmt.Println("Processing:", file)
		srcFile, err := os.Open(file)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer srcFile.Close()

		reader := bufio.NewReaderSize(srcFile, bufferSize)
		writer := bufio.NewWriterSize(outFile, bufferSize)

		if _, err := io.Copy(writer, reader); err != nil {
			fmt.Println("Error:", err)
			return
		}

		if err := writer.Flush(); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	fmt.Println("Merge successful. Output file:", outputFile)
}

func extractNumber(fileName string) (int, error) {
	match := regexp.MustCompile(`(\d+)`).FindStringSubmatch(filepath.Base(fileName))
	if len(match) < 2 {
		return 0, fmt.Errorf("failed to extract number from filename: %s", fileName)
	}
	return strconv.Atoi(match[1])
}