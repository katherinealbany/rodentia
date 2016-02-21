package logger

import (
	"fmt"
	"time"
)

var Level int

func Debug(message string) {
	if Level >= 0 {
		fmt.Println(timestamp(), "[DEBUG]", message)
	}
}

func Info(message string) {
	if Level >= 1 {
		fmt.Println(timestamp(), "[INFO ]", message)
	}
}

func Warn(message string) {
	if Level >= 2 {
		fmt.Println(timestamp(), "[WARN ]", message)
	}
}

func Error(message string) {
	if Level >= 3 {
		fmt.Println(timestamp(), "[ERROR]", message)
	}
}

func Fatal(message string) {
	if Level >= 5 {
		fmt.Println(timestamp(), "[FATAL]", message)
	}
}

func timestamp() (timestamp string) {
	t := time.Now()
	return "[" + string(t.String()) + "]"
}
