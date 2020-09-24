package surfline

import (
	"time"
)

// CamRewind - Surfline cam rewind data
type CamRewind struct {
	RecordingURL  string    `json:"recordingUrl"`
	StartAtUtc    time.Time `json:"startDate"`
	EndAtUtc      time.Time `json:"endDate"`
	LocalFilePath string
}
