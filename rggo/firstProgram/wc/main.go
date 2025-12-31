package main

import (
	"bufio" // to read text
	"flag"
	"fmt" // format text
	"io"  // provides input output reader interface
	"os"  // to use os resources
)

func main() {
	countLines := flag.Bool("l", false, "Count lines")
	countBytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	c, err := count(os.Stdin, *countLines, *countBytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
	fmt.Println(c)
}

func count(r io.Reader, countLines bool, countBytes bool) (int, error) {
	scanner := bufio.NewScanner(r)
	if countLines && countBytes {
		fmt.Fprintln(os.Stderr, "cannot give -l -b together.")
		os.Exit(1)
	}

	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	if err := scanner.Err(); err != nil {
		return wc, err
	}
	return wc, nil
}
