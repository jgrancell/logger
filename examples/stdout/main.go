package main

import (
	"github.com/jgrancell/logger"
	"github.com/jgrancell/logger/formats"
)

func main() {
	DefaultLogger()
	DefaultJsonLogger()
	PrettyJsonLogger()
}

// By default, Logger logs to standard out using the Logger Standard Log Format
//   This shows how Logger runs with no configurations passed.
func DefaultLogger() {
	log := logger.Logger{}
	if err := log.Init(); err != nil {
		_ = log.Fallback.Log(err.Error())
	}

	// We can catch any errors the logger has and use the FallbackHandler to ensure they arent lost.
	if err := log.Init(); err != nil {
		_ = log.Fallback.Log(err.Error())
	}

	log.Info("example standard out message")
	log.Batch("aaaa-bbbb-cccc-dddd").Info("example standard out message with batched log id")
}

func DefaultJsonLogger() {
	// You can also override the default formatter. We use the JSON formatter below.
	jog := logger.Logger{
		Format: &formats.LoggerFormat{
			AsJson: true,
		},
	}
	_ = jog.Init()
	jog.Info("example json formatted standard out message")
	jog.Batch("aaaa-bbbb-cccc-dddd").Info("example json formatted standard out message with batched log id")
}

func PrettyJsonLogger() {
	// JSON format can even be made pretty, if you need to be able to read it more easily
	jog := logger.Logger{
		Format: &formats.LoggerFormat{
			AsJson:      true,
			PrettyPrint: true,
		},
	}
	_ = jog.Init()
	jog.Info("example json formatted standard out message")
	jog.Batch("aaaa-bbbb-cccc-dddd").Info("example json formatted standard out message with batched log id")
}
