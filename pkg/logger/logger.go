package logger

// Logger desc
type Logger interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

// FakeLogger desc
type FakeLogger struct {
}

// Trace desc
func (fake *FakeLogger) Trace(args ...interface{}) {
	return
}

// Debug desc
func (fake *FakeLogger) Debug(args ...interface{}) {
	return
}

// Info desc
func (fake *FakeLogger) Info(args ...interface{}) {
	return
}

// Warn desc
func (fake *FakeLogger) Warn(args ...interface{}) {
	return
}

// Error desc
func (fake *FakeLogger) Error(args ...interface{}) {
	return
}

// Fatal desc
func (fake *FakeLogger) Fatal(args ...interface{}) {
	return
}

// Panic desc
func (fake *FakeLogger) Panic(args ...interface{}) {
	return
}

// Tracef desc
func (fake *FakeLogger) Tracef(format string, args ...interface{}) {
	return
}

// Debugf desc
func (fake *FakeLogger) Debugf(format string, args ...interface{}) {
	return
}

// Infof desc
func (fake *FakeLogger) Infof(format string, args ...interface{}) {
	return
}

// Warnf desc
func (fake *FakeLogger) Warnf(format string, args ...interface{}) {
	return
}

// Errorf desc
func (fake *FakeLogger) Errorf(format string, args ...interface{}) {
	return
}

// Fatalf desc
func (fake *FakeLogger) Fatalf(format string, args ...interface{}) {
	return
}

// Panicf desc
func (fake *FakeLogger) Panicf(format string, args ...interface{}) {
	return
}
