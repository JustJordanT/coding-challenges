### Code Documentation: Word Count Program in Go

#### Package
- **main**: This is the primary package for the application.

#### Imports
- **bufio**: Provides buffered I/O utilities.
- **flag**: Helps in parsing command-line flags.
- **fmt**: Implements formatted I/O with functions analogous to C's printf and scanf.
- **io**: Provides basic interfaces to I/O primitives.
- **os**: Supplies a platform-independent interface to operating system functionality.
- **strings**: Contains functions for manipulating strings.
- **unicode/utf8**: Supplies Unicode UTF-8 encoding and decoding.

#### Types
- **Count**: A struct used to hold the counts of lines, words, characters, and bytes.
  - `lines`: Integer, counts the number of lines.
  - `words`: Integer, counts the number of words.
  - `chars`: Integer, counts the number of characters.
  - `bytes`: Integer, counts the number of bytes.

#### Functions

- **main()**:
  - Initializes command-line flags for counting lines, words, characters, and bytes.
  - Parses the command-line flags.
  - Processes files passed as command-line arguments, or standard input if no files are provided.
  - For each input source, it calls `wc()` to perform the counting, and `printCount()` to display the results.

- **wc(reader io.Reader) Count**:
  - Takes an `io.Reader` as input.
  - Counts lines, words, characters, and bytes in the provided input.
  - Uses a scanner to read input line by line.
  - Words are counted by splitting lines into fields.
  - Characters are counted using UTF-8 encoding.
  - Returns a `Count` struct with the totals.

- **printCount(count Count, countLines, countWords, countChars, countBytes bool)**:
  - Takes a `Count` struct and boolean flags indicating what to print.
  - Conditionally prints the counts of lines, words, characters, and bytes based on the provided flags.

#### Error Handling
- In `main()`, if an error occurs while opening a file, it prints an error message to the standard error stream and continues processing the next file.

#### Command-Line Flags
- `-l`: Count lines.
- `-w`: Count words.
- `-m`: Count characters.
- `-c`: Count bytes.

Each flag is a boolean type, and when set to true, the corresponding count (lines, words, characters, or bytes) is performed and displayed.

## Running the program

To run this program and see its output, you can use various command-line flags to specify what counts you want (lines, words, characters, bytes) and provide either file names or use standard input for the text source. Here are some example usages:

1. **Counting Lines, Words, Characters, and Bytes from a File**
   - Command: `go run main.go -l -w -m -c file.txt`
   - This command will process the contents of `file.txt` and print the counts of lines, words, characters, and bytes.

2. **Counting Only Words and Characters from Multiple Files**
   - Command: `go run main.go -w -m file1.txt file2.txt`
   - This command will process `file1.txt` and `file2.txt`, printing the counts of words and characters for each file.

3. **Counting Lines from Standard Input**
   - Command: `echo "Hello World" | go run main.go -l`
   - This uses `echo` to send a string to the standard input of the program. The program will then count and print the number of lines (which, in this case, is 1).

4. **Counting Bytes from a File**
   - Command: `go run main.go -c file.txt`
   - This will count and display the number of bytes in `file.txt`.

5. **Interactive Mode with Standard Input**
   - Command: `go run main.go -w`
   - Run the command and then type some text directly into the terminal. Press `Ctrl+D` (on Unix-like systems) or `Ctrl+Z` then `Enter` (on Windows) to signal the end of input. The program will then count and print the number of words in the input text.

Remember, these examples assume you have Go installed and your environment is set up to compile and run Go programs. Also, replace `file.txt`, `file1.txt`, and `file2.txt` with actual file paths on your system.
