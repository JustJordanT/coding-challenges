package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

type Count struct {
	lines, words, chars, bytes int
}

func main() {
	countLines := flag.Bool("l", false, "Count lines")
	countWords := flag.Bool("w", false, "Count words")
	countChars := flag.Bool("m", false, "Count characters")
	countBytes := flag.Bool("c", false, "Count bytes")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		count := wc(os.Stdin)
		printCount(count, *countLines, *countWords, *countChars, *countBytes)
	} else {
		for _, filename := range files {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error opening file %s: %v\n", filename, err)
				continue
			}
			count := wc(file)
			file.Close()
			printCount(count, *countLines, *countWords, *countChars, *countBytes)
		}
	}
}

func wc(reader io.Reader) Count {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var count Count
	for scanner.Scan() {
		count.lines++
		line := scanner.Text()
		count.words += len(strings.Fields(line))
		count.bytes += len(line) // Count bytes
		count.chars += utf8.RuneCountInString(line)
	}

	return count
}

func printCount(count Count, countLines, countWords, countChars, countBytes bool) {
	if countLines {
		fmt.Printf("Lines: %d\n", count.lines)
	}
	if countWords {
		fmt.Printf("Words: %d\n", count.words)
	}
	if countChars {
		fmt.Printf("Characters: %d\n", count.chars)
	}
	if countBytes {
		fmt.Printf("Bytes: %d\n", count.bytes)
	}
}
