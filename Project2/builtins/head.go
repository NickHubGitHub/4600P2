package builtins

import (
	"bufio"
	"fmt"
	"os"
)

func Head(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: head <file>")
	}

	filename := args[0]

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//reads first 15 lines of a file
	lineCount := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lineCount++
		if lineCount >= 15 {
			break
		}
	}
	return nil
}
