package main



import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)



func main() {

	if len(os.Args) < 2 {
		panic(fmt.Sprintln("usage: serialClean <filename>"))
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("could not open file %s : %s ", os.Args[1], err))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)

		if err == nil && token != nil {
			n, ierr := strconv.ParseInt(string(token), 10, 32)
			if ierr != nil {
				fmt.Fprintln(os.Stderr, "Invalid input: %s", ierr)
				token = nil
			}
			if n > 1024 {
				fmt.Fprintln(os.Stderr, "Number to big: %s", n)
				token = nil
			}
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			fmt.Printf("%s\n", scanner.Text())
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid input: %s", err)
	}

}