package logger

// Fatal Message Accessors
func (e *EntryBatch) Fatal(message string) error {
	return e.WithMessage(LOG_FATAL, message).Handle()
}

// Error Message Accessors
func (e *EntryBatch) Error(err error) error {
	return e.WithMessage(LOG_ERROR, err.Error()).Handle()
}

// Warning Message Accessors
func (e *EntryBatch) Warning(message string) error {
	return e.WithMessage(LOG_WARNING, message).Handle()
}

// Info Message Accessors
func (e *EntryBatch) Info(message string) error {
	return e.WithMessage(LOG_INFO, message).Handle()
}

// Debug Message Accessors
func (e *EntryBatch) Debug(message string) error {
	return e.WithMessage(LOG_DEBUG, message).Handle()
}

// Trace Message Accessors
func (e *EntryBatch) Trace(message string) error {
	return e.WithMessage(LOG_TRACE, message).Handle()
}
