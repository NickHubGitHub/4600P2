package builtins

import (
	"fmt"
	"os"
)

func PrintWorkingDirectory() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %v", err)
	}
	fmt.Println(wd)
	return nil
}
