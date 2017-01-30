package wallabago

import (
	"encoding/json"
	"io/ioutil"
)

// Config containg all data to access wallabag API
var Config WallabagConfig

// WallabagConfig contains all data needed to connect to wallabag API like URL, id and secret of the API client and user name and according password
type WallabagConfig struct {
	WallabagURL  string
	ClientID     string
	ClientSecret string
	UserName     string
	UserPassword string
}

func ReadConfig(configPath string) (err error) {
	Config, err = getConfig(configPath)
}

func getConfig(configPath string) (config wallabago.WallabagConfig, err error) {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	config, err = readJSON(raw)
}

func readJSON(raw []byte) (config wallabago.WallabagConfig, err error) {
	err = json.Unmarshal(raw, &config)
}
