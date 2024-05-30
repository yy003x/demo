package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

var _ logrus.Formatter = (*logFormatter)(nil)

type logFormatter struct{}

func (lf *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+6)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	data["x_level"] = strings.ToUpper(entry.Level.String())
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	encoder := json.NewEncoder(b)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON %+v", err)
	}

	return b.Bytes(), nil
}
