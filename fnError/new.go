package fnError

import (
	"github.com/d3v-friends/go-tools/fnString"
	"github.com/pkg/errors"
)

func New(message string) error {
	return errors.New(message)
}

func NewF(format string, args ...any) error {
	return errors.Errorf(format, args...)
}

func NewFields(message string, field map[string]any) error {
	if len(field) != 0 {
		message += ": " + fnString.Stringify(field)
	}
	return NewF(message)
}
