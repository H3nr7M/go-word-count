package main
// bufio is a package in the Go standard library that implements buffered I/O. It exposes an API similar to the one in the io package, but with the added benefit of buffering.
// log is a package in the Go standard library that implements a simple logging package. It exposes a simple API for logging messages to a variety of output destinations.
import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    // Get the file name from the command-line arguments
	// because the comand is ./WordCount doc.txt
    if len(os.Args) < 2 {
        log.Fatal("Please provide a file name")
    }

    fileName := os.Args[1]

    // Open the file
	// Open the file for reading. If there is an error, log it and exit.
    file, err := os.Open(fileName)
	//nil is a special value in Go that represents the absence of a value.
    if err != nil {
        log.Fatal(err)
    }
	//defer is a keyword in Go that schedules a function call (the deferred function) to be run immediately before the function executing the defer returns. When a function is deferred, it's not executed immediately but instead pushed onto a stack, and it will be executed right before the surrounding function returns or panics.
    defer file.Close()

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)

    // Initialize a variable to count the number of words
    wordCount := 0

    // Iterate over each line of the file
    for scanner.Scan() {
        // Split the line into words using whitespace as the delimiter
        words := strings.Fields(scanner.Text())

        // Increment the word count by the number of words in the line
        wordCount += len(words)
    }

    // Check if there was an error while reading the file
	//short variable declaration
	//The semicolon ; is used to separate multiple statements written in a single line in Go.
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // Print the word count
    fmt.Printf("Word count for file %s: %d\n", fileName, wordCount)
}
