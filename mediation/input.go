package mediation

import (
	"github.com/katherinealbany/rodentia/logger"
	"os"
)

var log = logger.New("mediation")

func Read() {
	var filename string = "test.csv"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]byte, 10)
	log.Info("Reading", filename, "...")

	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Read", count, "bytes", data)

	count, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Read", count, "bytes", data)
}
