package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create an array to store the lines
	var lines []string

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Display the lines
	fmt.Println("Contents of test.txt:")
	for _, line := range lines {
		fmt.Println(line)
	}
}
