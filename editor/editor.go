package editor

import (
	"fmt"
	"log"
	"math"
	"os/exec"
	"sync"
	"time"

	"github.com/lancehumiston/resurf/garmin"
	"github.com/lancehumiston/resurf/surfline"
)

// TrimRecordings - trims the mp4 files associated with the cam rewinds based on wave times
func TrimRecordings(waveTimes []garmin.WaveTime, camRewinds []*surfline.CamRewind) {
	var waveIdx int
	var wg sync.WaitGroup

	wg.Add(len(waveTimes))
	for _, c := range camRewinds {
		for _, w := range waveTimes[waveIdx:] {
			if c.EndAtUtc.Before(w.StartAtUtc) {
				continue
			}

			waveIdx++

			log.Printf("[%d] w.StartAtUtc %v", waveIdx, w.StartAtUtc)
			log.Printf("[%d] c.StartAtUtc %v", waveIdx, c.StartAtUtc)

			go trimRecording(waveIdx, w, *c, &wg)
		}
	}

	wg.Wait()
}

func trimRecording(id int, w garmin.WaveTime, c surfline.CamRewind, wg *sync.WaitGroup) {
	defer wg.Done()

	filePath := fmt.Sprintf("./out-%d.mp4", id)

	// use buffer to account for delay in the watch recognizing the wave
	const bufferSeconds = 7
	position := formatTimeArg(math.Max(w.StartAtUtc.Sub(c.StartAtUtc).Seconds()-bufferSeconds, 0))
	duration := formatTimeArg(w.EndAtUtc.Sub(w.StartAtUtc).Seconds() + bufferSeconds)

	args := []string{"-i", c.LocalFilePath, "-ss", position, "-t", duration, filePath}
	cmd := exec.Command("ffmpeg", args...)
	log.Printf("[%d] Running command with args %v...", id, args)

	if err := cmd.Run(); err != nil {
		log.Printf("[%d] Finished command with error: %v", id, err)
		return
	}

	log.Printf("[%d] Finished command successfully", id)
}

func formatTimeArg(seconds float64) string {
	t := time.Date(0, 0, 0, 0, 0, int(seconds), 0, time.UTC)
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}
