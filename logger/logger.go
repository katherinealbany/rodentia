package logger

import "fmt"

var Level int

func Debug(message string) {
	if Level >= 0 {
		fmt.Println("[DEBUG]", message)
	}
}

func Info(message string) {
	if Level >= 1 {
		fmt.Println("[INFO ]", message)
	}
}

func Warn(message string) {
	if Level >= 2 {
		fmt.Println("[WARN ]", message)
	}
}

func Error(message string) {
	if Level >= 3 {
		fmt.Println("[ERROR]", message)
	}
}

func Fatal(message string) {
	if Level >= 5 {
		fmt.Println("[FATAL]", message)
	}
}
