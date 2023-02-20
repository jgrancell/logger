package logger

import (
	"reflect"
	"testing"
	"time"

	"github.com/jgrancell/logger/formats"
	"github.com/jgrancell/logger/handlers"
)

func TestDefaultLogger(t *testing.T) {

	log := &Logger{}
	_ = log.Init()

	defaultFormat := &formats.LoggerFormat{
		AsJson:         false,
		DateTimeFormat: time.RFC3339,
		PrettyPrint:    false,
	}

	defaultHandler := &handlers.StdoutHandler{}

	if reflect.TypeOf(log.Handler) != reflect.TypeOf(defaultHandler) {
		t.Fatal("The default log handler is not an instance of the stdout handler.")
	}

	if reflect.TypeOf(log.Format) != reflect.TypeOf(defaultFormat) {
		t.Fatal("The default format is not an instance of the LoggerFormat Formatter")
	}
}
