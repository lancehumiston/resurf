package surfline

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
)

// GetCamRewinds - returns the recent surf cam rewinds from specified camAlias in descending datetime order
func GetCamRewinds(camAlias string) []CamRewind {
	url := fmt.Sprintf("https://www.surfline.com/surfdata/video-rewind/video_rewind.cfm?camalias=%s", camAlias)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	return parse(res.Body)
}

// DownloadRecordings - downloads the rewind videos and sets the LocalFilePath on the camRewinds
func DownloadRecordings(camRewinds []*CamRewind) ([]*CamRewind, error) {
	var urls []string
	var wg sync.WaitGroup
	var err error

	for _, v := range camRewinds {
		if v.RecordingURL == "" {
			continue
		}

		v.LocalFilePath = fmt.Sprintf("./%v-%v.mp4", toFileFormat(v.StartAtUtc), toFileFormat(v.EndAtUtc))

		if contains(urls, v.RecordingURL) {
			continue
		}

		urls = append(urls, v.RecordingURL)

		if fileExists(v.LocalFilePath) {
			continue
		}

		wg.Add(1)
		go func(path string, url string) {
			defer wg.Done()

			if e := downloadFile(path, url); e != nil {
				errors.Wrap(err, e.Error())
			}
		}(v.LocalFilePath, v.RecordingURL)
	}

	wg.Wait()

	return camRewinds, err
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func toFileFormat(t time.Time) string {
	s := strings.Replace(t.Format(time.RFC3339), " ", "", -1)
	return strings.Replace(s, ":", "-", -1)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func parse(reader io.Reader) []CamRewind {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the recordings from HTML
	re := regexp.MustCompile(`var recordings = (\[.*?\])`)
	camRewinds := []CamRewind{}
	json.Unmarshal(re.FindSubmatch(body)[1], &camRewinds)

	return reverse(camRewinds)
}

// reverse - returns the slice with its items in the reverse order https://stackoverflow.com/a/28058324
func reverse(s []CamRewind) []CamRewind {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
