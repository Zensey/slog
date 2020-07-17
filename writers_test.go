package slog

import (
	"testing"
)

func Test(t *testing.T) {
	l := FileLogger("log.")
	l.SetLevel(LevelTrace)
	l.Infof("Hello %d!", 123)

	l = ConsoleLogger()
	l.SetLevel(LevelTrace)
	l.Infof("Hello %d!", 123)

	l = SyslogLogger("myapp")
	l.SetLevel(LevelTrace)
	l.Infof("Hello %d!", 123)
}
