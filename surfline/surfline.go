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
	"github.com/playwright-community/playwright-go"
)

// GetCamRewinds - returns the recent surf cam rewinds from specified camAlias in descending datetime order
func GetCamRewinds(camAlias string) []CamRewind {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// Bypass CF bot detection
	page.SetExtraHTTPHeaders(map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		"Accept-Language": "en-US,en;q=0.9",
	})
	if _, err = page.Goto("https://www.surfline.com/surfdata/video-rewind/video_rewind.cfm?camalias=" + camAlias); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	html, err := page.Content()
	if err != nil {
		log.Fatalf("could get html: %v", err)
	}
	rewinds, err := parse(strings.NewReader(html))
	if err != nil {
		log.Fatalf("could not parse html: %v", err)
	}

	page.Close()
	browser.Close()

	return rewinds
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

func parse(reader io.Reader) ([]CamRewind, error) {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the recordings from HTML
	re := regexp.MustCompile(`var recordings = (\[.*\])`)
	submatchGroups := re.FindSubmatch(body)
	if len(submatchGroups) <= 1 {
		log.Println(string(body))
		return nil, errors.New("Failed to parse input")
	}

	camRewinds := []CamRewind{}
	json.Unmarshal(submatchGroups[1], &camRewinds)

	return reverse(camRewinds), nil
}

// reverse - returns the slice with its items in the reverse order https://stackoverflow.com/a/28058324
func reverse(s []CamRewind) []CamRewind {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
