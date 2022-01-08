package logger

import (
	"time"

	"github.com/jgrancell/logger/formats"
	"github.com/jgrancell/logger/handlers"
)

const (
	LOG_TRACE = iota
	LOG_DEBUG
	LOG_INFO
	LOG_WARNING
	LOG_ERROR
	LOG_FATAL
)

var (
	severityStrings = []string{
		"trace",
		"debug",
		"info",
		"warning",
		"error",
		"fatal",
	}
)

type Logger struct {
	Name     string
	Level    int
	Handler  handlers.Handler
	Format   formats.Format
	Fallback *handlers.FallbackHandler
}

func (l *Logger) Initialize() error {
	return l.Init()
}

func (l *Logger) Init() error {
	l.Fallback = &handlers.FallbackHandler{}

	if l.Handler == nil {
		l.Handler = &handlers.StdoutHandler{}
	}
	if err := l.Handler.Init(); err != nil {
		return err
	}

	if l.Format == nil {
		l.Format = &formats.LoggerFormat{}
	}
	l.Format.Init()

	return nil
}

func (l *Logger) Batch(id string) *EntryBatch {
	e := &EntryBatch{
		IsFormatted: false,
		Format:      l.Format,
		BatchId:     id,
		Messages:    make([]*Message, 0),
		IsHandled:   false,
		Handler:     l.Handler,
	}
	return e
}

func (l *Logger) CreateEntry(severity int, message string) *EntryBatch {
	if severity < l.Level {
		// The entry doesn't meet our minimum severity level
		return &EntryBatch{}
	}
	e := &EntryBatch{
		IsFormatted: false,
		Format:      l.Format,
		BatchId:     "",
		Messages:    make([]*Message, 0),
		IsHandled:   false,
		Handler:     l.Handler,
	}

	e.Messages = append(e.Messages, &Message{
		Severity: severityStrings[severity],
		Raw:      message,
		Time:     time.Now(),
	})
	return e
}
