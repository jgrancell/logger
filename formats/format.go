package formats

import "time"

type Format interface {
	Init()
	GetDefaultParameters() map[string]string
	Format(time.Time, string, string, map[string]string) (string, error)
}

type Metadata struct {
	Level     int
	Message   string
	Timestamp string
}

type Message struct {
	Severity   string
	Raw        string
	Timestamp  time.Time
	Parameters map[string]string
}
