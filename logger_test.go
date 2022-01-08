package logger

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jgrancell/logger/formats"
	"github.com/jgrancell/logger/handlers"
)

func TestDefaultLogger(t *testing.T) {

	log := &Logger{}
	log.Init()

	defaultFormat := &formats.LoggerFormat{
		AsJson:         false,
		DateTimeFormat: time.RFC3339,
		PrettyPrint:    false,
	}

	defaultHandler := &handlers.StdoutHandler{}

	if !cmp.Equal(log.Handler, defaultHandler) {
		t.Fatalf("The default log handler is not a default instance of the stdout handler.")
	}

	if !cmp.Equal(log.Format, defaultFormat) {
		fmt.Println(log.Format)
		fmt.Println(defaultFormat)
		t.Fatalf("The default format is not a default instance of the Logger Standard Foramt.")
	}
}
