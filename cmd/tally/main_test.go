package main

import (
	"bytes"
	"errors"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "tally")
	if out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}

	const input = "pear apple pear banana apple cherry pear"
	tests := []struct {
		name     string
		args     []string
		stdin    string
		stdout   string
		exitCode int
	}{
		{"default alphabetical", nil, input, "apple 2\nbanana 1\ncherry 1\npear 3\n", 0},
		{"top 3", []string{"-top", "3"}, input, "pear 3\napple 2\nbanana 1\n", 0},
		{"top equals form", []string{"-top=1"}, input, "pear 3\n", 0},
		{"top exceeds distinct", []string{"-top", "10"}, input, "pear 3\napple 2\nbanana 1\ncherry 1\n", 0},
		{"top empty input", []string{"-top", "3"}, "", "", 0},
		{"default empty input", nil, "", "", 0},
		{"zero", []string{"-top", "0"}, input, "", 2},
		{"negative", []string{"-top", "-1"}, input, "", 2},
		{"noninteger", []string{"-top", "abc"}, input, "", 2},
		{"fraction", []string{"-top", "1.5"}, input, "", 2},
		{"overflow", []string{"-top", strings.Repeat("9", 100)}, input, "", 2},
		{"missing argument", []string{"-top"}, input, "", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(bin, tt.args...)
			cmd.Stdin = strings.NewReader(tt.stdin)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			err := cmd.Run()
			code := 0
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				code = exitErr.ExitCode()
			} else if err != nil {
				t.Fatalf("run: %v", err)
			}
			if code != tt.exitCode {
				t.Fatalf("exit code = %d, want %d (stderr %q)", code, tt.exitCode, stderr.String())
			}
			if stdout.String() != tt.stdout {
				t.Fatalf("stdout = %q, want %q", stdout.String(), tt.stdout)
			}
			if tt.exitCode == 2 {
				if !strings.Contains(stderr.String(), "-top") || !strings.Contains(stderr.String(), "Usage") {
					t.Fatalf("stderr = %q, want usage diagnostic", stderr.String())
				}
			} else if stderr.Len() != 0 {
				t.Fatalf("stderr = %q, want empty", stderr.String())
			}
		})
	}
}
