package formats

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	KUBERNETES_DATETIME = "2006-01-02T15:04:05-0700"
	KUBERNETES_FORMAT   = "%s%s %s:%d \"%s\""
)

type KubernetesFormat struct {
	*Metadata
	AsJson         bool
	DateTimeFormat string
	PrettyPrint    bool
	LogParameters  map[string]string
}

func (f *KubernetesFormat) Init() {
	if f.DateTimeFormat == "" {
		f.DateTimeFormat = KUBERNETES_DATETIME
	}
}

func (f *KubernetesFormat) GetDefaultParameters() map[string]string {
	parameters := make(map[string]string)
	for key, value := range f.LogParameters {
		parameters[key] = value
	}
	return parameters
}

func (f *KubernetesFormat) Format(timestamp time.Time, severity string, message string, params map[string]string) (string, error) {

	k8s := KubernetesMessage{
		Severity:   severity,
		Message:    message,
		Timestamp:  timestamp,
		Parameters: params,
	}

	if f.AsJson {
		panic("Json output is not yet available for the Kubernetes Formatter")
		//return f.ToJson(k8s)
	}
	return f.ToString(k8s)
}

func (f *KubernetesFormat) ToString(k KubernetesMessage) (string, error) {
	file, line := f.GetCaller()

	t := k.Timestamp.Format(KUBERNETES_DATETIME)

	formatted := fmt.Sprintf(
		KUBERNETES_FORMAT,
		strings.ToUpper(string(k.Severity[0])),
		t,
		file,
		line,
		k.Message,
	)

	for key, value := range k.Parameters {
		formatted = fmt.Sprintf(
			"%s %s=\"%s\"",
			formatted,
			key,
			value,
		)
	}
	return formatted, nil
}

func (f *KubernetesFormat) GetCaller() (string, int) {
	_, file, line, ok := runtime.Caller(5)
	if !ok {
		file = "???"
		line = 1
	} else {
		if slash := strings.LastIndex(file, "/"); slash >= 0 {
			path := file
			file = path[slash+1:]
		}
	}
	return file, line
}
