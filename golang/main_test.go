package main

import (
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile-*.tmp")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	if !Exists(tempFile.Name()) {
		t.Errorf("Exists should return true for an existing file")
	}

	if Exists(tempFile.Name() + "-notfound") {
		t.Errorf("Exists should return false for a non-existent file")
	}
}
