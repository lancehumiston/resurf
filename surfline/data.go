package surfline

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// GetCamRewinds - returns the recent surf cam rewinds from specified camAlias in descending datetime order
func GetCamRewinds(camAlias string) []CamRewind {
	url := fmt.Sprintf("https://www.surfline.com/surfdata/video-rewind/video_rewind.cfm?camalias=%s", camAlias)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the recordings from HTML
	re := regexp.MustCompile(`var recordings = (\[.*?\])`)
	camRewinds := []CamRewind{}
	json.Unmarshal(re.FindSubmatch(body)[1], &camRewinds)

	return camRewinds
}
