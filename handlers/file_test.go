package handlers

import (
	"io/ioutil"
	"os"
	"testing"
)

func createTestDir(t *testing.T) string {
	tempDir, err := ioutil.TempDir(".", "testtemp")
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

	handler.Handle("foobar fizzbuzz")
	s, err := ioutil.ReadFile(handler.Path)
	if err != nil {
		t.Errorf("expected to read file %s, got err: %v", handler.Path, err)
	}

	if string(s) != "foobar fizzbuzz\n" {
		t.Errorf("got unexpected file content value: %s", string(s))
	}
}
