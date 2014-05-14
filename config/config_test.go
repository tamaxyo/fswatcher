package config

import (
	"testing"
)

var jsonBlob = []byte(`[
  {
    "path" : "./", 
    "pattern" : "*.go",
    "ignore" : ".*", 
    "command" : "go test ./...",
    "recursive" : true
  },
  {
    "path" : "./test", 
    "pattern" : "*.md",
    "command" : "go test ./..."
  }
]`)

var target []Config

func init() {
	target, _ = Parse(jsonBlob)
}

func TestConfigurationSize(t *testing.T) {
	expected := 2
	actual := len(target)
	if actual != expected {
		t.Errorf("incorrect number of parsed configurations - actual: %s, expected: %s", actual, expected)
	}
}

func TestPath(t *testing.T) {
	expected := "./"
	actual := target[0].Path
	if actual != expected {
		t.Errorf("incorrect path has been set - actual: %s, expected: %s", actual, expected)
	}
}

func TestPattern(t *testing.T) {
	expected := "*.md"
	actual := target[1].Pattern
	if actual != expected {
		t.Errorf("incorrect pattern has been set - actual: %s, expected: %s", actual, expected)
	}
}

func TestIgnore(t *testing.T) {
	expected := ".*"
	actual := target[0].Ignore
	if actual != expected {
		t.Errorf("incorrect ignore pattern has been set - actual: %s, expected: %s", actual, expected)
	}
}

func TestCommand(t *testing.T) {
	expected := "go test ./..."
	actual := target[0].Command
	if actual != expected {
		t.Errorf("incorrect command has been set - actual: %s, expected: %s", actual, expected)
	}
}

func TestRecursive(t *testing.T) {
	expected := true
	actual := target[0].Recursive
	if actual != expected {
		t.Errorf("incorrect value of recursive has been set - actual: %s, expected: %s", actual, expected)
	}
}

func TestUnspecifiedString(t *testing.T) {
	expected := ""
	actual := target[1].Ignore
	if actual != expected {
		t.Errorf("incorrect value has been set to unspecified string field - actual: %s, expected: %s", actual, expected)
	}
}

func TestUnspecifiedBool(t *testing.T) {
	expected := false
	actual := target[1].Recursive
	if actual != expected {
		t.Errorf("incorrect value has been set to unspecified bool field - actual: %s, expected: %s", actual, expected)
	}
}
