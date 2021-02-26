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
	TimesFilePath = "c:/users/lance.humiston/documents/projects/go/resurf/SurfData.txt"
)

// GetWaveTimes - returns a slice of wave times, parsed from the reader content, in ascending datetime order
func GetWaveTimes(reader io.Reader) []WaveTime {
	var waveTimes []WaveTime

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		times := scanner.Text()
		parts := strings.Split(times, "|")
		start, err := time.Parse(time.RFC3339, parts[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := time.Parse(time.RFC3339, parts[1])
		if err != nil {
			log.Fatal(err)
		}

		waveTime := WaveTime{
			StartAtUtc: start.UTC(),
			EndAtUtc:   end.UTC(),
		}

		waveTimes = append(waveTimes, waveTime)
	}

	return waveTimes
}
