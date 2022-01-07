package logger

import "time"

func (e *EntryBatch) WithBatchId(id string) *EntryBatch {
	e.BatchId = id
	return e
}

func (e *EntryBatch) WithMessage(severity int, message string) *EntryBatch {
	m := &Message{
		Severity: severityStrings[severity],
		Raw:      message,
		Time:     time.Now(),
	}

	m.Parameters = e.Format.GetDefaultParameters()
	if e.BatchId != "" {
		m.Parameters["batch"] = e.BatchId
	}

	e.Messages = append(e.Messages, m)
	return e
}
