package fnError

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func New(message string) error {
	return errors.New(message)
}

func NewF(format string, args ...any) error {
	return errors.Errorf(format, args...)
}

func NewFields(message string, field map[string]any) error {
	var body, err = json.Marshal(field)
	if err != nil {
		return NewF("%s: marshal_error=%s", message, err.Error())
	}

	return NewF("%s: %s", message, string(body))
}
