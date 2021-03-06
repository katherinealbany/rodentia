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

var start time.Time

func init() {
	start = time.Now()
}

func New(name string) *Logger {
	return &Logger{
		Name:  name,
		Level: 6,
	}
}

func (logger *Logger) Debug(args ...interface{}) {
	if logger.Level >= 0 {
		fmt.Println(timestamp(), since(), "[DEBUG]", name(logger.Name), args)
	}
}

func (logger *Logger) Info(args ...interface{}) {
	if logger.Level >= 1 {
		fmt.Println(timestamp(), since(), "[INFO ]", name(logger.Name), args)
	}
}

func (logger *Logger) Warn(args ...interface{}) {
	if logger.Level >= 2 {
		fmt.Println(timestamp(), since(), "[WARN ]", name(logger.Name), args)
	}
}

func (logger *Logger) Error(args ...interface{}) {
	if logger.Level >= 3 {
		fmt.Println(timestamp(), since(), "[ERROR]", name(logger.Name), args)
	}
}

func (logger *Logger) Panic(err error, args ...interface{}) {
	if logger.Level >= 5 {
		fmt.Println(timestamp(), since(), "[PANIC]", name(logger.Name), args)
		panic(err)
	}
}

func (logger *Logger) Fatal(args ...interface{}) {
	if logger.Level >= 5 {
		fmt.Println(timestamp(), since(), "[FATAL]", name(logger.Name), args)
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
