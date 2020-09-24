package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"time"

	"github.com/lancehumiston/resurf/garmin"
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

	rewindPtrs, err := filterCamRewinds(waveTimes, camRewinds)
	if err != nil {
		log.Fatal(err)
	}

	rewindPtrs, err = surfline.DownloadRecordings(rewindPtrs)
	if err != nil {
		log.Fatal(err)
	}

	cutRecordings(waveTimes, rewindPtrs)
}

func cutRecordings(waveTimes []garmin.WaveTime, camRewinds []*surfline.CamRewind) {
	var waveIdx int

	for _, c := range camRewinds {
		for _, w := range waveTimes[waveIdx:] {
			if c.EndAtUtc.Before(w.StartAtUtc) {
				continue
			}

			fmt.Println("w.StartAtUtc", w.StartAtUtc)
			fmt.Println("c.StartAtUtc", c.StartAtUtc)

			// use buffer to account for delay in the watch recognizing the wave
			const bufferSeconds = 5
			diff := w.StartAtUtc.Sub(c.StartAtUtc)
			seconds := math.Max(diff.Seconds()-bufferSeconds, 0)
			diffTime := time.Date(0, 0, 0, 0, 0, int(seconds), 0, time.UTC)
			offset := fmt.Sprintf("%02d:%02d:%02d", diffTime.Hour(), diffTime.Minute(), diffTime.Second())

			duration := w.EndAtUtc.Sub(w.StartAtUtc)
			lengthTime := time.Date(0, 0, 0, 0, 0, int(duration.Seconds())+bufferSeconds, 0, time.UTC)
			length := fmt.Sprintf("%02d:%02d:%02d", lengthTime.Hour(), lengthTime.Minute(), lengthTime.Second())

			filePath := fmt.Sprintf(fmt.Sprintf("./out-%d.mp4", waveIdx))

			args := []string{"-i", c.LocalFilePath, "-ss", offset, "-t", length, filePath}
			cmd := exec.Command("ffmpeg", args...)
			log.Printf("Running command with args %v and waiting for it to finish...", args)
			err := cmd.Run()
			log.Printf("Command finished with error: %v", err)

			waveIdx++
		}
	}
}

func filterCamRewinds(waveTimes []garmin.WaveTime, camRewinds []surfline.CamRewind) ([]*surfline.CamRewind, error) {
	if camRewinds == nil || len(camRewinds) == 0 {
		return nil, fmt.Errorf("Empty list of camRewinds")
	}

	if waveTimes == nil || len(waveTimes) == 0 {
		return nil, fmt.Errorf("Empty list of waveTimes")
	}

	if waveTimes[len(waveTimes)-1].EndAtUtc.Before(camRewinds[0].StartAtUtc) {
		return nil, fmt.Errorf("Wave times are out of rewind range")
	}

	if waveTimes[0].StartAtUtc.After(camRewinds[len(camRewinds)-1].EndAtUtc) {
		return nil, fmt.Errorf("Surfline camRewinds data is stale")
	}

	var waveIdx int
	rewindsOfInterest := make([]*surfline.CamRewind, len(waveTimes))
	for i, c := range camRewinds {
		// roi have been found for all waves
		if waveIdx == len(waveTimes) {
			return rewindsOfInterest, nil
		}

		// avoid additional checks if the cam rewind is outside of the wave times range
		if c.EndAtUtc.Before(waveTimes[waveIdx].StartAtUtc) {
			continue
		}
		if c.StartAtUtc.After(waveTimes[len(waveTimes)-1].EndAtUtc) {
			return rewindsOfInterest, nil
		}

		// check for overlap in times
		for _, w := range waveTimes[waveIdx:] {
			if w.StartAtUtc.Before(c.EndAtUtc) {
				rewindsOfInterest[waveIdx] = &camRewinds[i]
				waveIdx++
			}
		}
	}

	return rewindsOfInterest, nil
}
