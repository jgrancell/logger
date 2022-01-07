package handlers

type Handled struct {
	HandlerError error
}

func (h *Handled) ReturnWithInt() int {
	if h.HandlerError == nil {
		return 0
	}
	return 1
}
