package builtins_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/NickHubGitHub/4600P2/Project2/builtins"
)

func TestHead(t *testing.T) {
	// temporary file to test
	tmpFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	content := []byte("testing this\nfile. woo woo\n")
	if _, err := tmpFile.Write(content); err != nil {
		t.Fatalf("Failed to write : %v", err)
	}
	var buf bytes.Buffer
	if err := builtins.Cat(tmpFile.Name()); err != nil {
		t.Fatalf("Cat() error = %v, want nil", err)
	}
	got := buf.String()
	want := "Hello, world!\nThis is a test file.\n"
	if got != want {
		t.Errorf("Cat() = %q, want %q", got, want)
	}
}
