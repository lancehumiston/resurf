package main

import (
	"log"
	"os"

	"github.com/lancehumiston/resurf/editor"
	"github.com/lancehumiston/resurf/garmin"
	"github.com/lancehumiston/resurf/resurf"
	"github.com/lancehumiston/resurf/surfline"
)

func main() {
	log.Println("Welcome to resurf")

	file, err := os.Open(garmin.TimesFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	waveTimes := garmin.GetWaveTimes(file)
	log.Println(waveTimes)

	camRewinds := surfline.GetCamRewinds("wc-windansea")

	rewindPtrs, err := resurf.FilterCamRewinds(waveTimes, camRewinds)
	if err != nil {
		log.Fatal(err)
	}

	rewindPtrs, err = surfline.DownloadRecordings(rewindPtrs)
	if err != nil {
		log.Fatal(err)
	}

	editor.TrimRecordings(waveTimes, rewindPtrs)

	log.Println("Happy resurfing!")
}
