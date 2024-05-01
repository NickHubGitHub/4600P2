package builtins_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/NickHubGitHub/4600P2/Project2/builtins"
)

func TestPrintWorkingDirectory(t *testing.T) {
	//testing with a temporary directory
	tmpDir := t.TempDir()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	var buf bytes.Buffer
	want := tmpDir + "\n"
	if err := builtins.PrintWorkingDirectory(); err != nil {
		t.Fatalf("PrintWorkingDirectory() error = %v, want nil", err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("PrintWorkingDirectory() = %q, want %q", got, want)
	}
}
