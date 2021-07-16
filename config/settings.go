package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Settings - application settings
type Settings struct {
	CamAliases []string      `json:"camAliases"`
	WaveTime   time.Duration `json:"waveTime"`
}

const (
	filepath = "./config.json"
)

// Appsettings - application settings
var Appsettings Settings

func init() {
	if fileExists(filepath) {
		readSettingsFromFilePath()
		return
	}

	writeSettingsToFilePath()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readSettingsFromFilePath() {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Panic(err)
	}

	if err := json.Unmarshal(content, &Appsettings); err != nil {
		log.Panic(err)
	}
}

func writeSettingsToFilePath() {
	Appsettings.CamAliases = make([]string, 0)
	Appsettings.WaveTime, _ = time.ParseDuration("30s")
	appsettings, err := json.Marshal(Appsettings)
	if err != nil {
		log.Panic(err)
	}
	if err := ioutil.WriteFile(filepath, appsettings, 0644); err != nil {
		log.Panic(err)
	}
}
