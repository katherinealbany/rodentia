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
		Name:  name,
		Level: 6,
	}
}

func Debug(logger *Logger, message string) {
	if Level >= 0 {
		fmt.Println(timestamp(), since(), "[DEBUG]", name(logger.Name), message)
	}
}

func Info(logger *Logger, message string) {
	if Level >= 1 {
		fmt.Println(timestamp(), since(), "[INFO ]", name(logger.Name), message)
	}
}

func Warn(logger *Logger, message string) {
	if Level >= 2 {
		fmt.Println(timestamp(), since(), "[WARN ]", name(logger.Name), message)
	}
}

func Error(logger *Logger, message string) {
	if Level >= 3 {
		fmt.Println(timestamp(), since(), "[ERROR]", name(logger.Name), message)
	}
}

func Panic(logger *Logger, message string) {
	if Level >= 5 {
		fmt.Println(timestamp(), since(), "[PANIC]", name(logger.Name), message)
	}
}

func Fatal(logger *Logger, message string) {
	if Level >= 5 {
		fmt.Println(timestamp(), since(), "[FATAL]", name(logger.Name), message)
		os.Exit(1)
	}
}

func name(name string) string {
	return "[" + name + "]"
}

func timestamp() string {
	t := time.Now()
	return "[" + t.Format("15:04:05.000") + "]"
}

func since() string {
	ms := int64(time.Since(start) / time.Millisecond)
	return "[" + strconv.FormatInt(ms, 10) + "]"
}
