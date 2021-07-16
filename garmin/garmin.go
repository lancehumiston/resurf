package garmin

import (
	"bufio"
	"io"
	"log"
	"strings"
	"time"
)

const (
	// TimesFilePath - file path to riding times log file
	TimesFilePath = "./SurfData.txt"
	// ManualCaptureKey - indicates a manually captured wave and that start time should be calculated
	ManualCaptureKey = "-"
)

// GetWaveTimes - returns a slice of wave times, parsed from the reader content, in ascending datetime order
func GetWaveTimes(reader io.Reader, defaultWaveTime time.Duration) []WaveTime {
	var waveTimes []WaveTime

	// use buffer to account for delay in the watch recognizing the wave
	bufferDuration, err := time.ParseDuration("7s")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		times := scanner.Text()
		parts := strings.Split(times, "|")
		var start time.Time
		var end time.Time
		if parts[0] == ManualCaptureKey {
			end, err = time.Parse(time.RFC3339, parts[1])
			if err != nil {
				log.Fatal(err)
			}

			start = end.Add(-defaultWaveTime)
			dur, _ := time.ParseDuration("7s")
			end = end.Add(dur)
		} else {
			start, err = time.Parse(time.RFC3339, parts[0])
			if err != nil {
				log.Fatal(err)
			}
			start.Add(-bufferDuration)

			end, err = time.Parse(time.RFC3339, parts[1])
			if err != nil {
				log.Fatal(err)
			}
		}

		waveTime := WaveTime{
			StartAtUtc: start.UTC(),
			EndAtUtc:   end.UTC(),
		}

		waveTimes = append(waveTimes, waveTime)
	}

	return waveTimes
}
