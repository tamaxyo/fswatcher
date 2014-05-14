package main

import (
	"testing"
)

func TestMatched(t *testing.T) {
	name := "foo.go"
	pattern := "*.go"
	if !match(name, pattern) {
		t.Errorf("should be match the pattern - name: %s, pattern: %s", name, pattern)
	}
}

func TestMatchedInMultiplePatterns(t *testing.T) {
	name := "foo.go"
	pattern := "*.rb,*.py,*.go"
	if !match(name, pattern) {
		t.Errorf("should be match the pattern - name: %s, pattern: %s", name, pattern)
	}
}

func TestUnmatched(t *testing.T) {
	name := "foo.txt"
	pattern := "*.rb,*.py,*.go"
	if match(name, pattern) {
		t.Errorf("should not be match the pattern - name: %s, pattern: %s", name, pattern)
	}
}

func TestCommandArgs(t *testing.T) {
	command := "go test ./..."
	cmd := setupCommand("/path/to/the/project", command)

	expected := []string{"go", "test", "./..."}
	actual := cmd.Args

	if len(actual) != len(expected) {
		t.Errorf("incorrect number of arguments has been set - actual: %d, expected: %d", len(actual), len(expected))
	}

	for idx, _ := range expected {
		if actual[idx] != expected[idx] {
			t.Errorf("incorrect command argument has been set - actual: %s, expected: %s", actual[idx], expected[idx])
		}

	}
}

func TestCommandDir(t *testing.T) {
	command := "go test ./..."
	cmd := setupCommand("/path/to/the/project", command)

	expected := "/path/to/the/project"
	actual := cmd.Dir

	if actual != expected {
		t.Errorf("incorrect workdir has been set - actual: %s, expected: %s", actual, expected)
	}
}
