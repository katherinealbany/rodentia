package main

import "github.com/katherinealbany/rodentia/logger"

func main() {
	logger.Level = 1
	logger.Debug("level 1")
	logger.Info("level 2")
	logger.Warn("level 3")
	logger.Error("level 4")
	logger.Fatal("level 5")
}
