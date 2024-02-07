package wallabago

import "encoding/json"

// Configuration represents the object being returned from the API request /config
type Configuration struct {
	ID                int     `json:"id"`
	ItemsPerPage      int     `json:"items_per_page"`
	Language          string  `json:"language"`
	FeedToken         string  `json:"feed_token"`
	FeedLimit         int     `json:"feed_limit"`
	ReadingSpeed      float64 `json:"reading_speed"`
	ActionMarkAsRead  int     `json:"action_mark_as_read"`
	ListMode          int     `json:"list_mode"`
	DisplayThumbnails int     `json:"display_thumbnails"`
}

// Config returns the config of the configured wallabag instance
func Config(bodyByteGetterFunc BodyByteGetter) (Configuration, error) {
	var c Configuration
	configJSONByte, err := bodyByteGetterFunc(LibConfig.WallabagURL+"/api/config", "GET", nil)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(configJSONByte, &c)
	if err != nil {
		return c, err
	}
	return c, err
}
