package main

import (
	"log"
	"os"

	"github.com/lancehumiston/resurf/garmin"
)

func main() {
	log.Println("Welcome to resurf")
	file, err := os.Open(garmin.TimesFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	surfTimes := garmin.GetSurfTimes(file)
	log.Println(surfTimes)
}
