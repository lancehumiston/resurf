package garmin

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

const times string = `2021-02-24T12:03:37-08:00|2021-02-24T12:11:34-08:00
2021-02-24T12:18:21-08:00|2021-02-24T12:18:28-08:00`

func TestGetWaveTimes_ValidTimes_ReturnsWaveTimes(t *testing.T) {
	// Arrange
	r := strings.NewReader(times)
	expected := []WaveTime{
		{
			StartAtUtc: mustParseTime(t, "2021-02-24T20:03:37Z"),
			EndAtUtc:   mustParseTime(t, "2021-02-24T20:11:34Z"),
		},
		{
			StartAtUtc: mustParseTime(t, "2021-02-24T20:18:21Z"),
			EndAtUtc:   mustParseTime(t, "2021-02-24T20:18:28Z"),
		},
	}

	// Act
	response := GetWaveTimes(r)

	// Assert
	if !reflect.DeepEqual(response, expected) {
		t.Fatalf("expected:%v actual: %v", expected, response)
	}
}

func mustParseTime(t *testing.T, value string) time.Time {
	time, err := time.Parse(time.RFC3339, value)
	if err != nil {
		t.Fatalf("%s not in valid `RFC3339` format", value)
	}

	return time
}
