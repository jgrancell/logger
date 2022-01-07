// Standard Out Logger
package handlers

import (
	"fmt"
)

type FallbackHandler struct{}

func (h *FallbackHandler) Handle(message string) error {
	fmt.Println("The fallback handler has been triggered due to a logging failure by the configured handler.")
	fmt.Println("Handler Error:", message)
	return nil
}

func (h *FallbackHandler) Log(message string) error {
	return h.Handle(message)
}
