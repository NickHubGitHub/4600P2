package builtins

import (
	"bufio"
	"fmt"
	"os"
)

func Cat(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: cat <file>")
	}

	filename := args[0]

	// Opens the file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	//reads file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}
