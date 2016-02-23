package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Logger struct {
	Level uint8
	Name  string
}

var (
	Level uint8
	Name  string
	start time.Time
)

func init() {
	start = time.Now()
}

func New(name string) *Logger {
	return &Logger{
		Name:  "main",
		Level: 6,
	}
}

func Debug(logger *Logger, message string) {
	if Level >= 0 {
		fmt.Println(timestamp(), since(), "[DEBUG]", name(), message)
	}
}

func Info(logger *Logger, message string) {
	if Level >= 1 {
		fmt.Println(timestamp(), since(), "[INFO ]", name(), message)
	}
}

func Warn(logger *Logger, message string) {
	if Level >= 2 {
		fmt.Println(timestamp(), since(), "[WARN ]", name(), message)
	}
}

func Error(logger *Logger, message string) {
	if Level >= 3 {
		fmt.Println(timestamp(), since(), "[ERROR]", name(), message)
	}
}

func Panic(logger *Logger, message string) {
	if Level >= 5 {
		fmt.Println(timestamp(), since(), "[PANIC]", name(), message)
	}
}

func Fatal(logger *Logger, message string) {
	if Level >= 5 {
		fmt.Println(timestamp(), since(), "[FATAL]", name(), message)
		os.Exit(1)
	}
}

func name() string {
	return "[" + Name + "]"
}

func timestamp() string {
	t := time.Now()
	return "[" + t.Format("15:04:05.000") + "]"
}

func since() string {
	ms := int64(time.Since(start) / time.Millisecond)
	return "[" + strconv.FormatInt(ms, 10) + "]"
}
