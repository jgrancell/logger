package handlers

import (
	"fmt"
	"io/fs"
	"os"

	"golang.org/x/sys/unix"
)

type FileHandler struct {
	Path string
	Mode fs.FileMode
}

func (h *FileHandler) Init() error {
	if h.Mode == 0 {
		h.Mode = 0600
	}

	if h.Path == "" {
		return fmt.Errorf("the FileHandler requires a Path attribute")
	}

	file, err := os.OpenFile(h.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, h.Mode)
	if err != nil {
		return err
	}
	defer file.Close()

	return unix.Access(h.Path, unix.W_OK)
}

func (h *FileHandler) Handle(message string) error {
	file, err := os.OpenFile(h.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, h.Mode)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(message + "\n"); err != nil {
		return err
	}
	return nil
}
