package handlers

import "fmt"

type AbstractHandler struct {
	Level int
	Name  string
}

func (h *AbstractHandler) Log(message string, err error) error {
	fmt.Println("This is the abstract handler. Message:", message)
	return nil
}
