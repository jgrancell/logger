package main

import (
	"os"

	"github.com/jgrancell/logger"
	"github.com/jgrancell/logger/formatters"
	"github.com/jgrancell/logger/handlers"
)

func main() {
	// The File Handler has a required option, the log file path
	file_logger := logger.Logger{
		Handler: &handlers.FileHandler{Path: "test.txt"},
	}

	// We can catch any errors the logger has and use the FallbackHandler to ensure they arent lost.
	if err := file_logger.Init(); err != nil {
		// All logging commands give access to the convenience returns ReturnWithInt, ReturnWithError, and Return
		//   to provide int, error, and bool returns respectively.
		os.Exit(file_logger.FallbackHandler.Handle(err.Error()).ReturnWithInt())
	}

	file_logger.Info("example default format file log entry")
	file_logger.WithIdent("aaaa-bbbb-cccc-dddd").Info("example default format file log entry with identifier")

	// Formatters are independent of Handlers, so you can use any Formatter with any Handler.
	// We switch to JSON format below
	file_json := logger.Logger{
		Handler:   &handlers.FileHandler{Path: "./test.txt"},
		Formatter: &formatters.JsonFormatter{},
	}
	_ = file_json.Init()
	file_json.Info("example json formatted file log entry")
	file_json.WithIdent("aaaa-bbbb-cccc-dddd").Info("example json formatted file log entry with identifier")
}
