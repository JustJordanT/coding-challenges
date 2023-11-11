package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func CheckErr(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		//return
	}
}

func main() {

	readBytes := flag.Bool("c", false, "reads bytes from file")
	readLines := flag.Bool("l", false, "read lines from file")
	readWords := flag.Bool("w", false, "read words from file")
	readCharacters := flag.Bool("m", false, "read characters from file")
	flag.Parse()

	var arg string

	if !*readBytes && !*readLines && !*readWords {
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
		readBytesFunc(arg)
	}

	if *readLines {
		readLinesFunc(arg)
	}

	if *readWords {
		readWordsFunc(arg)
	}

	if *readCharacters {
		_, err := readCharactersFunc(arg)
		CheckErr(err)
	}

	if !*readBytes && !*readLines && !*readWords {
		readLinesFunc(arg)
		readWordsFunc(arg)
		readBytesFunc(arg)
	}
}

func readCharactersFunc(arg string) (int, error) {
	return fmt.Println(len(readFileToString(arg)), arg)
}

func readBytesFunc(arg string) {
	file, err := os.Stat(arg)
	CheckErr(err)
	fmt.Println(file.Size(), arg)
}

func readLinesFunc(arg string) bool {
	file, done := openFile(arg)
	if done {
		return true
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
		return true
	}

	fmt.Println(lineCount, arg)
	return false
}

func readWordsFunc(arg string) {
	fileStrings := readFileToString(arg)
	fields := strings.Fields(fileStrings)

	fmt.Println(len(fields), arg)
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
