package wallabago

import "encoding/json"

// Information represents the object being returned from the API request /info
type Information struct {
	Appname             string `json:"appname"`
	Version             string `json:"version"`
	AllowedRegistration bool   `json:"allowed_registration"`
}

// Info returns the info of the configured wallabag instance
func Info(bodyByteGetterFunc BodyByteGetter) (Information, error) {
	var info Information
	infoJSONByte, err := bodyByteGetterFunc(LibConfig.WallabagURL+"/api/info", "GET", nil)
	if err != nil {
		return info, err
	}
	err = json.Unmarshal(infoJSONByte, &info)
	if err != nil {
		return info, err
	}
	return info, err
}
