package formats

import "time"

type LoggerMessage struct {
	Timestamp string `json:"time"`
	Severity  string `json:"level"`
	Message   string `json:"msg"`
	Batch     string `json:"batch,omitempty"`
}

type KubernetesMessage struct {
	Timestamp  time.Time
	Severity   string
	Message    string
	Parameters map[string]string
}
