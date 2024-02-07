package wallabago

import "encoding/json"

// LoggedInUser represents the object being returned from the API request /user
type LoggedInUser struct {
	ID        int           `json:"id"`
	UserName  string        `json:"username"`
	Email     string        `json:"email"`
	CreatedAt *WallabagTime `json:"created_at"`
	UpdatedAt *WallabagTime `json:"updated_at"`
}

// User returns the user info of the logged in user of the configured wallabag instance
func User(bodyByteGetterFunc BodyByteGetter) (LoggedInUser, error) {
	var u LoggedInUser
	userJSONByte, err := bodyByteGetterFunc(LibConfig.WallabagURL+"/api/user", "GET", nil)
	if err != nil {
		return u, err
	}
	err = json.Unmarshal(userJSONByte, &u)
	if err != nil {
		return u, err
	}
	return u, err
}
