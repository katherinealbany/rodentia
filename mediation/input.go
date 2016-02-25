package mediation

import (
	"fmt"
	"github.com/katherinealbany/rodentia/logger"
	"io/ioutil"
)

var log = logger.New("mediation")

func Read() {
	log.Info("Reading ...")
	data, err := ioutil.ReadFile("./test.csv")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}
