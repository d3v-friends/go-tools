package fnLogger

type GoError struct {
	message    string
	params     map[string]string
	stacktrace string
}

func NewGoError(message string, params ...map[string]string) {
	var e = new(GoError)
	e.message = message
	e.stacktrace = ""
}

func (x *GoError) Error() string {
	return x.message
}
