package main

import (
	"strings"
	"testing"
)

func TestGoVersion(t *testing.T) {
	v, err := GoVersion()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	acceptable := []string{
		"devel", "go1.0", "go1.1", "go1.2", "go1.3",
		"go1.4", "go1.4.1", "go1.4.2", "go1.4.3",
		"go1.5", "go1.5.1", "go1.5.2", "go1.5.3", "go1.5.4",
		"go1.6", "go1.6.1", "go1.6.2", "go1.6.3", "go1.6.4",
		"go1.7", "go1.7.1", "go1.7.2", "go1.7.3", "go1.7.4", "go1.7.5", "go1.7.6",
		"go1.8", "go1.8.1", "go1.8.2", "go1.8.3", "go1.8.4", "go1.8.5", "go1.8.6", "go1.8.7",
		"go1.9", "go1.9.1", "go1.9.2", "go1.9.3", "go1.9.4", "go1.9.5", "go1.9.6", "go1.9.7",
		"go1.10", "go1.10.1", "go1.10.2", "go1.10.3", "go1.10.4", "go1.10.5", "go1.10.6", "go1.10.7", "go1.10.8",
		"go1.11", "go1.11.1", "go1.11.2", "go1.11.3", "go1.11.4", "go1.11.5",
		"go1.12",
	}
	found := false
	for _, expected := range acceptable {
		if strings.HasPrefix(v, expected) {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("bad: %#v", v)
	}
}
