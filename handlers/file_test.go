package handlers

import (
	"os"
	"testing"
)

func createTestDir(t *testing.T) string {
	tempDir, err := os.MkdirTemp(".", "testtemp")
	if err != nil {
		t.Errorf("failed creating tempdir: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll(tempDir) })
	return tempDir
}

func TestFileHandlerInit(t *testing.T) {
	tempDir := createTestDir(t)

	handler := &FileHandler{
		Path: tempDir + "/test.file",
		Mode: 0600,
	}

	err := handler.Init()
	if err != nil {
		t.Errorf("expected to create file %s, got err: %v", handler.Path, err)
	}

	err = handler.Handle("foobar fizzbuzz")
	if err != nil {
		t.Errorf("got error %v while writing log line", err)
	}
	s, err := os.ReadFile(handler.Path)
	if err != nil {
		t.Errorf("expected to read file %s, got err: %v", handler.Path, err)
	}

	if string(s) != "foobar fizzbuzz\n" {
		t.Errorf("got unexpected file content value: %s", string(s))
	}
}
