package surfline

import (
	"os"
	"strings"
	"testing"
)

func TestGetCamRewinds_ValidCamAlias_ReturnsCamRewinds(t *testing.T) {
	// Arrange
	const camAlias string = "wc-windansea"

	// Act
	response := GetCamRewinds(camAlias)

	// Assert
	if len(response) == 0 {
		t.Fatalf("unexpected response:%v", response)
	}
	actual := response[0].RecordingURL
	if !strings.Contains(actual, camAlias) {
		t.Fatalf("response recording url:%s does not contain expected cam alias:%s", actual, camAlias)
	}
}

func TestDownloadRecordings_ValidCamRewinds_ReturnsCamRewindsWithLocalFilePathToFile(t *testing.T) {
	// Arrange
	allCamRewinds := GetCamRewinds("wc-windansea")
	camRewinds := []*CamRewind{&allCamRewinds[0]}

	// Act
	response, err := DownloadRecordings(camRewinds)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if len(response) == 0 {
		t.Fatalf("unexpected response:%v", response)
	}
	actual := response[0].LocalFilePath
	if !fileExists(actual) {
		t.Fatalf("local file path not found %s", actual)
	}

	// Cleanup
	os.Remove(actual)
}
