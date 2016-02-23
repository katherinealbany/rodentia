package mediation

import (
	"github.com/katherinealbany/rodentia/logger"
)

var log2 = logger.New("mediation")

func Read() {
	logger.Info(log2, "Reading ...")
}
