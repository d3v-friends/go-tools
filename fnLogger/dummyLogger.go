package fnLogger

type DummyLogger struct {
}

func (x *DummyLogger) Trace(format string, args ...any) {
}

func (x *DummyLogger) Info(format string, args ...any) {
}

func (x *DummyLogger) Warn(format string, args ...any) {
}

func (x *DummyLogger) Error(format string, args ...any) {
}

func (x *DummyLogger) Fatal(format string, args ...any) {
}

func (x *DummyLogger) WithFields(fields ...Fields) IfLogger {
	return &DummyLogger{}
}

func (x *DummyLogger) SetLevel(lv Level) {
}
