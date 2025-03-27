package fnError

import "fmt"

func Concat(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}

	var msg = ""
	for _, err := range errs {
		msg = fmt.Sprintf("%s, %s", msg, err.Error())
	}

	return New(msg[2:])
}
