package builtins

import (
	"bufio"
	"fmt"
	"os"
)

func Tail(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: tail <file>")
	}
	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//reads last 15 lines of a file
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) > 15 {
			lines = lines[1:]
		}
	}
	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}
