package main

import (
	"github.com/katherinealbany/rodentia/logger"
	"github.com/katherinealbany/rodentia/mediation"
)

var log = logger.New("main")

func main() {
	logger.Info(log, "Begin")
	mediation.Read()
	logger.Info(log, "End")
}
