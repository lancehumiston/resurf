package resurf

import (
	"fmt"

	"github.com/lancehumiston/resurf/garmin"
	"github.com/lancehumiston/resurf/surfline"
)

// FilterCamRewinds - filters the cam rewinds to only include those associated with a wave time
func FilterCamRewinds(waveTimes []garmin.WaveTime, camRewinds []surfline.CamRewind) ([]*surfline.CamRewind, error) {
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
