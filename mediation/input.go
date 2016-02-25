package mediation

import (
	"fmt"
	"github.com/katherinealbany/rodentia/logger"
	"io/ioutil"
	"os"
)

var log = logger.New("mediation")

func Read() {
	log.Info("Reading ...")

	data, err := ioutil.ReadFile("./test.csv")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))

	file, err := os.Open("./test.csv")
	if err != nil {
		fmt.Println("error:", err)
		// log.Fatal("error:" + err)
	}
	defer file.Close()
}
