// Base logging handler
package handlers

const (
	FATAL_LOG_LEVEL = iota
	ERROR_LOG_LEVEL
	WARNING_LOG_LEVEL
	INFO_LOG_LEVEL
	DEBUG_LOG_LEVEL
	TRACE_LOG_LEVEL
)

type Handler interface {
	Init() error
	Handle(string) error
}
