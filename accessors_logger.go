package logger

func (l *Logger) LogAndExit(message interface{}) int {
	if _, ok := message.(error); ok {
		l.Error(message.(error))
		return 1
	}
	if _, ok := message.(string); ok {
		text := message.(string)
		if text != "" {
			l.Info(text)
		}
	}
	return 0
}

// Fatal Message Accessors
func (l *Logger) Fatal(message string) error {
	return l.CreateEntry(LOG_FATAL, message).Handle()
}

// Error Message Accessors
func (l *Logger) Error(err error) error {
	return l.CreateEntry(LOG_ERROR, err.Error()).Handle()
}

// Warning Message Accessors
func (l *Logger) Warning(message string) error {
	return l.CreateEntry(LOG_WARNING, message).Handle()
}

// Info Message Accessors
func (l *Logger) Info(message string) error {
	return l.CreateEntry(LOG_INFO, message).Handle()
}

// Debug Message Accessors
func (l *Logger) Debug(message string) error {
	return l.CreateEntry(LOG_DEBUG, message).Handle()
}

// Trace Message Accessors
func (l *Logger) Trace(message string) error {
	return l.CreateEntry(LOG_TRACE, message).Handle()
}
