package garmin

import (
	"time"
)

// WaveTime - timestamps indicating when the watch determined a wave was being ridden
type WaveTime struct {
	StartAtUtc time.Time
	EndAtUtc   time.Time
}
