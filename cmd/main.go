package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/lancehumiston/resurf/config"
	"github.com/lancehumiston/resurf/editor"
	"github.com/lancehumiston/resurf/garmin"
	"github.com/lancehumiston/resurf/resurf"
	"github.com/lancehumiston/resurf/surfline"
)

const (
	defaultCamAlias = "wc-hbpiersclose"
	camAliasUsage   = "the Surfline cam alias 'wc-*'"
)

func main() {
	log.Println("Welcome to resurf")

	var camAliasFlag string
	flag.StringVar(&camAliasFlag, "camAlias", defaultCamAlias, camAliasUsage)
	flag.StringVar(&camAliasFlag, "c", defaultCamAlias, camAliasUsage+" (shorthand)")
	flag.Parse()

	camAlias := getCamAlias(camAliasFlag)
	log.Println("Using cam alias:", camAlias)

	file, err := os.Open(garmin.TimesFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	waveTimes := garmin.GetWaveTimes(file)
	log.Println(waveTimes)

	camRewinds := surfline.GetCamRewinds(camAlias)

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

func getCamAlias(camAlias string) string {
	if camAlias != defaultCamAlias || len(config.Appsettings.CamAliases) == 0 {
		return camAlias
	}

	for i, v := range config.Appsettings.CamAliases {
		fmt.Printf("%d - %v\n", i, v)
	}
	fmt.Println("Which cam alias # would you like to use? (default '0')")
	var camIdx int
	if _, err := fmt.Scanf("%d", &camIdx); err != nil {
		camIdx = 0
	}

	return config.Appsettings.CamAliases[camIdx]
}
