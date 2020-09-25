package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Settings - application settings
type Settings struct {
	CamAliases []string `json:"camAliases"`
}

const (
	filepath = "./config.json"
)

// Appsettings - application settings
var Appsettings Settings

func init() {
	if fileExists(filepath) {
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Panic(err)
		}

		if err := json.Unmarshal(content, &Appsettings); err != nil {
			log.Panic(err)
		}
		return
	}

	Appsettings.CamAliases = make([]string, 0)
	appsettings, err := json.Marshal(Appsettings)
	if err != nil {
		log.Panic(err)
	}
	if err := ioutil.WriteFile(filepath, appsettings, 0644); err != nil {
		log.Panic(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
