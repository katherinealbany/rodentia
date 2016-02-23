package main

import (
	"github.com/katherinealbany/rodentia/logger"
	"github.com/katherinealbany/rodentia/mediation"
)

func init() {
	logger.Level = 6
	logger.Name = "main"
}

func main() {
	logger.Info("Begin")
	mediation.Read()
	logger.Info("End")
}
