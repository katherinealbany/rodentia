package main

import (
	"github.com/katherinealbany/rodentia/logger"
	"time"
)

func main() {
	logger.Level = 5

	time.Sleep(100 * time.Millisecond)
	logger.Debug("level 1")

	time.Sleep(100 * time.Millisecond)
	logger.Info("level 2")

	time.Sleep(100 * time.Millisecond)
	logger.Warn("level 3")

	time.Sleep(100 * time.Millisecond)
	logger.Error("level 4")

	time.Sleep(100 * time.Millisecond)
	logger.Fatal("level 5")
}
