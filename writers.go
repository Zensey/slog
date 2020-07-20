package slog

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"strconv"
	"time"
)

func FileLogger(prefix string) *slog {
	path := prefix + strconv.Itoa(int(time.Now().Unix())) + ".txt"
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", path, ":", err)
	}

	return &slog{lvl: LevelOff, b: NewBackend(f)}
}

type ConsoleWriter struct{}

func (c *ConsoleWriter) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}

func ConsoleLogger() *slog {
	w := &ConsoleWriter{}
	return &slog{lvl: LevelOff, b: NewBackend(w)}
}

func SyslogLogger(src string) *slog {
	sys, _ := syslog.New(syslog.LOG_INFO, src)
	return &slog{lvl: LevelOff, b: NewBackend(sys)}
}
