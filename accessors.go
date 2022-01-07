package logger

// Fatal Message Accessors
func (l *Logger) Fatal(message string) error {
	return l.CreateEntry(LOG_FATAL, message).Handle()
}

func (e *EntryBatch) Fatal(message string) error {
	return e.WithMessage(LOG_FATAL, message).Handle()
}

// Error Message Accessors
func (l *Logger) Error(err error) error {
	return l.CreateEntry(LOG_ERROR, err.Error()).Handle()
}

func (e *EntryBatch) Error(err error) error {
	return e.WithMessage(LOG_ERROR, err.Error()).Handle()
}

// Warning Message Accessors
func (l *Logger) Warning(message string) error {
	return l.CreateEntry(LOG_WARNING, message).Handle()
}

func (e *EntryBatch) Warning(message string) error {
	return e.WithMessage(LOG_WARNING, message).Handle()
}

// Info Message Accessors
func (l *Logger) Info(message string) error {
	return l.CreateEntry(LOG_INFO, message).Handle()
}

func (e *EntryBatch) Info(message string) error {
	return e.WithMessage(LOG_INFO, message).Handle()
}

// Debug Message Accessors
func (l *Logger) Debug(message string) error {
	return l.CreateEntry(LOG_DEBUG, message).Handle()
}

func (e *EntryBatch) Debug(message string) error {
	return e.WithMessage(LOG_DEBUG, message).Handle()
}

// Trace Message Accessors
func (l *Logger) Trace(message string) error {
	return l.CreateEntry(LOG_TRACE, message).Handle()
}

func (e *EntryBatch) Trace(message string) error {
	return e.WithMessage(LOG_TRACE, message).Handle()
}
