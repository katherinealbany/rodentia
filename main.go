package main

import (
	"github.com/katherinealbany/rodentia/logger"
	"github.com/katherinealbany/rodentia/mediation"
)

var log = logger.New("main")

func main() {
	log.Info(log, "Begin")
	mediation.Read()
	log.Info(log, "End")
}
