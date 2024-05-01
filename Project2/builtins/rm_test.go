package builtins_test

import (
	"os"
	"testing"

	"github.com/NickHubGitHub/4600P2/Project2/builtins"
)

func TestRemoveFile(t *testing.T) {
	//temporary file for testing
	tmpFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	err = builtins.RemoveFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("RemoveFile() error = %v, want nil", err)
	}
	// Check if file got removed
	_, err = os.Stat(tmpFile.Name())
	if !os.IsNotExist(err) {
		t.Errorf("RemoveFile() did not remove the file: %v", tmpFile.Name())
	}
}
