package formats

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	LOGGER_DATETIME     = "2006-01-02T15:04:05-0700"
	LOGGER_FORMAT       = "time=\"%s\" level=\"%s\" msg=\"%s\""
	LOGGER_FORMAT_PARAM = "%s  %s=\"%s\""
)

type LoggerFormat struct {
	*Metadata
	AsJson         bool
	DateTimeFormat string
	PrettyPrint    bool
}

type LoggerMessage struct {
	Timestamp string `json:"time"`
	Severity  string `json:"level"`
	Message   string `json:"msg"`
	Batch     string `json:"batch,omitempty"`
}

func (f *LoggerFormat) GetDefaultParameters() map[string]string {
	return make(map[string]string)
}

func (f *LoggerFormat) Format(timestamp time.Time, severity string, message string, params map[string]string) (string, error) {
	if f.DateTimeFormat == "" {
		f.DateTimeFormat = LOGGER_DATETIME
	}

	msg := LoggerMessage{
		Severity:  severity,
		Message:   message,
		Timestamp: timestamp.Format(f.DateTimeFormat),
	}

	if batch, ok := params["batch"]; ok {
		msg.Batch = batch
	}

	if f.AsJson {
		return f.ToJson(msg)
	}
	return f.ToString(msg)
}

func (f *LoggerFormat) ToJson(msg LoggerMessage) (string, error) {
	var entry_json []byte
	var err error
	if f.PrettyPrint {
		entry_json, err = json.MarshalIndent(msg, "", "  ")
	} else {
		entry_json, err = json.Marshal(msg)
	}
	if err != nil {
		return "", err
	}

	formatted := string(entry_json)
	return formatted, nil
}

func (f *LoggerFormat) ToString(msg LoggerMessage) (string, error) {
	formatted := fmt.Sprintf(
		LOGGER_FORMAT,
		msg.Timestamp,
		msg.Severity,
		msg.Message,
	)
	return formatted, nil
}
