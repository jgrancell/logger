package main

import (
	"github.com/jgrancell/logger"
	"github.com/jgrancell/logger/formats"
	"github.com/jgrancell/logger/handlers"
)

func main() {
	// The File Handler has a required option, the log file path
	file := logger.Logger{
		Handler: &handlers.FileHandler{Path: "test.txt"},
	}

	if err := file.Init(); err != nil {
		_ = file.Fallback.Log(err.Error())
	}

	_ = file.Info("example default format file log entry")
	_ = file.Batch("aaaa-bbbb-cccc-dddd").Info("example default format file log entry with batch id")

	// Formatters are independent of Handlers, so you can use any Formatter with any Handler.
	// We switch to JSON format below
	json := logger.Logger{
		Handler: &handlers.FileHandler{Path: "./test.txt"},
		Format:  &formats.LoggerFormat{AsJson: true},
	}
	_ = json.Init()
	_ = json.Info("example json formatted file log entry")
	_ = json.Batch("aaaa-bbbb-cccc-dddd").Info("example json formatted file log entry with batch id")
}
