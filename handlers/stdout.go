// Standard Out Logger
package handlers

import (
	"fmt"
	"os"
)

type StdoutHandler struct{}

func (h *StdoutHandler) Init() error {
	return nil
}

func (h *StdoutHandler) Handle(message string) error {
	if _, err := fmt.Fprintf(os.Stdout, "%s\n", message); err != nil {
		return err
	}
	return nil
}
