package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Specify the path to the MP3 file
	mp3FilePath := os.Args[1]

	outputFilePath := "output.txt"

	if len(os.Args) > 2 {
		outputFilePath = os.Args[2]
	}

	notify := fmt.Sprintf("Locating %s ", mp3FilePath)
	fmt.Println(notify)

	// Open the MP3 file
	file, err := os.Open(mp3FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Define the chunk size (in bytes)
	chunkSize := 1024 // Change this value as per your requirements

	// Create a buffer to hold the chunk of data
	buffer := make([]byte, chunkSize)

	// Create or open the output file for writing
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Read the file in chunks until the end
	for {
		// Read a chunk of data
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		// Check if we have reached the end of the file
		if bytesRead == 0 {
			break
		}

		// Process the chunk of data and write it to the output file
		processChunk(buffer[:bytesRead], outputFile)
	}
}

func processChunk(chunk []byte, outputFile *os.File) {
	// Write the chunk of data to the output file
	_, err := outputFile.Write(chunk)
	if err != nil {
		log.Fatal(err)
	}
}
