package logger

import "fmt"

type Logger interface {
	Log()
}

func Log(level int, message string) {
	fmt.Println(message)
}
