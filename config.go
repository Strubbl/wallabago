package wallabago

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// Config containing all data to access wallabag API
var Config WallabagConfig

// WallabagConfig contains all data needed to connect to wallabag API like URL, id and secret of the API client and user name and according password
type WallabagConfig struct {
	WallabagURL  string
	ClientID     string
	ClientSecret string
	UserName     string
	UserPassword string
}

// NewWallabagConfig initializes a new WallabagConfig
func NewWallabagConfig(wallabagURL, clientID, clientSecret, userName, userPassword string) WallabagConfig {
	return WallabagConfig{wallabagURL, clientID, clientSecret, userName, userPassword}
}

// SetConfig sets global wallabago Config
func SetConfig(wbgConfig WallabagConfig) {
	Config = wbgConfig
}

// ReadConfig will read the configuration from the given configJSON
// file and set the global Config setting with the results of the
// parsing
func ReadConfig(configJSON string) (err error) {
	Config, err = getConfig(configJSON)
	return
}

// getConfig reads a given configJSON file and parses the result, returning a parsed config object
func getConfig(configJSON string) (config WallabagConfig, err error) {
	raw, err := ioutil.ReadFile(configJSON)
	if err != nil {
		return
	}
	config, err = readJSON(raw)
	return
}

// readJSON parses a byte stream into a WallabagConfig object
func readJSON(raw []byte) (config WallabagConfig, err error) {
	// trim BOM bytes that make the JSON parser crash
	raw = bytes.TrimPrefix(raw, []byte("\xef\xbb\xbf"))
	err = json.Unmarshal(raw, &config)
	return
}
