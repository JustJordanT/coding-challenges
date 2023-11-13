package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func CheckErr(e error) {
	if e != nil {
		log.Fatal("Error: ", e)
	}
}

func main() {

	readBytes := flag.Bool("c", false, "reads bytes from file")
	readLines := flag.Bool("l", false, "read lines from file")
	readWords := flag.Bool("w", false, "read words from file")
	readCharacters := flag.Bool("m", false, "read characters from file")
	flag.Parse()

	var arg string

	if !*readBytes && !*readLines && !*readWords && !*readCharacters {
		arg = os.Args[1]
	} else {
		arg = os.Args[2]
	}

	//Make sure at least one argument is provided
	if len(arg) < 2 {
		fmt.Println("No file specified")
		return
	}

	if *readBytes {
		fmt.Println(readBytesFunc(arg), arg)
	}

	if *readLines {
		fmt.Println(readLinesFunc(arg), arg)
	}

	if *readWords {
		fmt.Println(readWordsFunc(arg), arg)
	}

	if *readCharacters {
		fmt.Println(readCharactersFunc(arg), arg)
	}

	if !*readBytes && !*readLines && !*readWords && !*readCharacters {
		fmt.Printf("%v %v %v - %v\n", readLinesFunc(arg), readWordsFunc(arg), readBytesFunc(arg), arg)
	}
}

func readCharactersFunc(arg string) int {
	return len(readFileToString(arg))
}

func readBytesFunc(arg string) int64 {
	file, err := os.Stat(arg)
	CheckErr(err)
	return file.Size()
}

func readLinesFunc(arg string) int {
	file, done := openFile(arg)
	if done {
		return 0
	}

	defer func(file *os.File) {
		err := file.Close()
		CheckErr(err)
	}(file)

	// Create a new Scanner to read the file
	scanner := bufio.NewScanner(file)
	lineCount := 0

	// Read each line and count
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	return lineCount
}

func readWordsFunc(arg string) int {
	fileStrings := readFileToString(arg)
	fields := strings.Fields(fileStrings)

	return len(fields)
}

func readFileToString(filePath string) string {
	file, err := os.ReadFile(filePath)
	CheckErr(err)

	str := string(file)

	return str
}

func openFile(filePath string) (*os.File, bool) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, true
	}

	return file, false
}
