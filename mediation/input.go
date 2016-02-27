package mediation

import (
	// "fmt"
	"github.com/katherinealbany/rodentia/logger"
	// "io/ioutil"
	"os"
)

var log = logger.New("mediation")

func Read() {
	var filename string = "./test.csv1"
	log.Info("Reading", filename)

	/*	data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		fmt.Print(string(data))*/

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
