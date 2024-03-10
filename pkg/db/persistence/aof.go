package persistence

import (
	"fmt"
	"log"
	"os"
)

var FileCounter int = 0

func checks(filename string) (bool, bool) {
	fmt.Println("Checking  the AOF")

	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false, false
	}

	maxSIZE := 2 * 1024 * 1024
	return fileInfo.Size() <= int64(maxSIZE), err == nil
}

func WriteAOF(query string) {
	fmt.Println("inside the AOF")
	filename := fmt.Sprintf("%d.txt", FileCounter)

	fileSize, fileExsist := checks(filename)
	if !fileExsist {
		_, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
	}
	if !fileSize {
		FileCounter++
		_, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
	}
	AppendOnlyFile(query, filename)
}
func AppendOnlyFile(query, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer file.Close()

	if _, err := file.WriteString(query + "\n"); err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
}
