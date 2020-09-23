package surfline

import (
	"time"
)

// CamRewind - Surfline cam rewind data
type CamRewind struct {
	RecordingURL string    `json:"recordingUrl"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
}
