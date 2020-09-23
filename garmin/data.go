package garmin

import (
	"bufio"
	"io"
	"log"
	"strings"
	"time"
)

const (
	// AppFilePath - file path to garmin application
	AppFilePath = "This PC/vívoactive 3 Music/Primary/GARMIN/APPS/SurfData.prg"
	// TimesFilePath = "This PC/vívoactive 3 Music/Primary/GARMIN/APPS/LOGS/SurfData.txt"
	// TimesFilePath - file path to riding times log file
	TimesFilePath = "c:/users/lance.humiston/documents/projects/go/resurf/SurfData.txt"
)

// GetSurfTimes - returns a map of wave times, parsed from the reader's content, in the format {<startTime>:<endTime>}
func GetSurfTimes(reader io.Reader) map[time.Time]time.Time {
	times := make(map[time.Time]time.Time)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		waveTimes := scanner.Text()
		parts := strings.Split(waveTimes, "|")
		start, err := time.Parse(time.RFC3339, parts[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := time.Parse(time.RFC3339, parts[1])
		if err != nil {
			log.Fatal(err)
		}
		times[start] = end
	}

	return times
}
