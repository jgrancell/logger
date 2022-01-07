package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/jgrancell/logger/formats"
	"github.com/jgrancell/logger/handlers"
)

type EntryBatch struct {
	IsFormatted bool
	Format      formats.Format
	BatchId     string
	Messages    []*Message
	IsHandled   bool
	Handler     handlers.Handler
}

type Message struct {
	Formatted  string
	Severity   string
	Raw        string
	Time       time.Time
	Parameters map[string]string
}

func (e *EntryBatch) AddMessage(severity int) {}

func (e *EntryBatch) Handle() error {
	var surfacedErr string
	for _, m := range e.Messages {
		//TODO: Add in better error handling here in the event only one message errors
		var err error

		// Formatting message
		if m.Formatted, err = e.Format.Format(
			m.Time,
			m.Severity,
			m.Raw,
			m.Parameters,
		); err != nil {
			surfacedErr = surfacedErr + " " + err.Error()
			continue
		}

		// Handling message
		if err := e.Handler.Handle(m.Formatted); err != nil {
			surfacedErr = surfacedErr + " " + err.Error()
		}

	}
	return fmt.Errorf("%s", strings.TrimSpace(surfacedErr))
}
