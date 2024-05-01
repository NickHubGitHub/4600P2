package builtins

import (
	"fmt"
	"os"
)

func RemoveFile(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: rm <file>")
	}
	filename := args[0]
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("failed to remove file: %v", err)
	}
	fmt.Printf("File '%s' removed successfully\n", filename)
	return nil
}
