package logger

import (
	"fmt"
	"time"
)

var Level int

var start time.Time

func init() {
	start = time.Now()
}

func Debug(message string) {
	if Level >= 0 {
		fmt.Println(timestamp(), since(), "[DEBUG]", message)
	}
}

func Info(message string) {
	if Level >= 1 {
		fmt.Println(timestamp(), since(), "[INFO ]", message)
	}
}

func Warn(message string) {
	if Level >= 2 {
		fmt.Println(timestamp(), since(), "[WARN ]", message)
	}
}

func Error(message string) {
	if Level >= 3 {
		fmt.Println(timestamp(), since(), "[ERROR]", message)
	}
}

func Fatal(message string) {
	if Level >= 5 {
		fmt.Println(timestamp(), since(), "[FATAL]", message)
	}
}

func timestamp() (timestamp string) {
	t := time.Now()
	return "[" + t.Format("15:04:05.000") + "]"
}

func since() (timestamp string) {
	return "[" + time.Since(start).String() + "]"
}
