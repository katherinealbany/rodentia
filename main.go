package main

import (
	"github.com/katherinealbany/rodentia/logger"
	"time"
)

func main() {
	logger.Level = 6

	time.Sleep(125 * time.Millisecond)
	logger.Debug("debug message")

	time.Sleep(25 * time.Millisecond)
	logger.Info("info  message")

	time.Sleep(25 * time.Millisecond)
	logger.Warn("warn  message")

	time.Sleep(25 * time.Millisecond)
	logger.Error("error message")

	time.Sleep(25 * time.Millisecond)
	logger.Panic("panic message")

	time.Sleep(25 * time.Millisecond)
	logger.Fatal("fatal message")
}
