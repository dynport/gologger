package gologger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	TIME_FORMAT = "2006-01-02T15:04:05.000"
	DEBUG       = iota
	INFO
	WARN
	ERROR
)

var LogColors = map[int]int{
	DEBUG: 102,
	INFO:  28,
	WARN:  214,
	ERROR: 196,
}

var LogPrefixes = map[int]string{
	DEBUG: "DEBUG",
	INFO:  "INFO ",
	WARN:  "WARN ",
	ERROR: "ERROR",
}

func colorize(c int, s string) (r string) {
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", c, s)
}

type Logger struct {
	LogLevel int
	Started  time.Time
	Prefix   string
	Colored  bool
	Caller   bool
}

func New() *Logger {
	return &Logger{Colored: true, LogLevel: INFO}
}

// starts a timer, current runtime is then added to log output
func (self *Logger) Start() {
	self.Started = time.Now()
}

// stops the timer
func (self *Logger) Stop() {
	self.Started = time.Unix(0, 0)
}

func (self *Logger) Inspect(i interface{}) {
	self.Debugf("%+v", i)
}

func (l *Logger) Debugf(format string, n ...interface{}) {
	l.logf(DEBUG, format, n...)
}

func (l *Logger) Infof(format string, n ...interface{}) {
	l.logf(INFO, format, n...)
}

func (l *Logger) Warnf(format string, n ...interface{}) {
	l.logf(WARN, format, n...)
}

func (l *Logger) Errorf(format string, n ...interface{}) {
	l.logf(ERROR, format, n...)
}

func (l *Logger) Debug(n ...interface{}) {
	l.log(DEBUG, n...)
}

func (l *Logger) Info(n ...interface{}) {
	l.log(INFO, n...)
}

func (l *Logger) Warn(n ...interface{}) {
	l.log(WARN, n...)
}

func (l *Logger) Error(n ...interface{}) {
	l.log(ERROR, n...)
}

func (l *Logger) logf(level int, s string, n ...interface{}) {
	if level >= l.LogLevel {
		l.write(l.logPrefix(level), fmt.Sprintf(s, n...))
	}
}

func (l *Logger) log(level int, n ...interface{}) {
	if level >= l.LogLevel {
		all := append([]interface{}{l.logPrefix(level)}, n...)
		l.write(all...)
	}
}

func (l *Logger) logPrefix(i int) (s string) {
	s = time.Now().Format(TIME_FORMAT)
	if l.Started.Unix() > 0 {
		time := fmt.Sprintf("%.3f", time.Now().Sub(l.Started).Seconds())
		s += fmt.Sprintf(" [%8s]", time)
	}
	if l.Prefix != "" {
		s = s + " [" + l.Prefix + "]"
	}
	s = s + " " + l.LogLevelPrefix(i)
	if l.Caller {
		_, file, line, ok := runtime.Caller(3)
		if ok {
			s += fmt.Sprintf(" [%s:%d]", filepath.Base(file), line)
		}
	}
	return s
}

func (l *Logger) LogLevelPrefix(level int) (s string) {
	prefix := LogPrefixes[level]
	if l.Colored {
		color := LogColors[level]
		return colorize(color, prefix)
	}
	return prefix
}

func (self *Logger) write(n ...interface{}) {
	fmt.Fprintln(os.Stderr, n...)
}
