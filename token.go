package wallabago

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var token Token

// Token represents the object being returned from the oauth process at the API containing the access token, expire time, type of token, scope and a refresh token
type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string
	RefreshToken string `json:"refresh_token"`
}

// getToken will use the credentials set in the configuration to
// request an access token from the wallabag API
func getToken() (Token, error) {
	var token Token
	tokenURL := Config.WallabagURL + "/oauth/v2/token"
	resp, err := http.PostForm(tokenURL,
		url.Values{"grant_type": {"password"},
			"client_id":     {Config.ClientID},
			"client_secret": {Config.ClientSecret},
			"username":      {Config.UserName},
			"password":      {Config.UserPassword},
		})
	if err != nil {
		return token, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return token, fmt.Errorf("getToken: bad response from server: %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}
	//log.Printf("GetToken: body=%v\n", string(body))
	err = json.Unmarshal(body, &token)
	return token, err
}

func checkForToken() error {
	var err error
	if token.TokenType == "" || token.AccessToken == "" {
		token, err = getToken()
	}
	return err
}

// GetAuthTokenHeader will make sure there's a working token and
// return a valid string to be used as an Authentication: header
func GetAuthTokenHeader() string {
	checkForToken()
	if token.TokenType == "" || token.AccessToken == "" {
		return ""
	}
	return strings.ToUpper(string(token.TokenType[0])) + token.TokenType[1:len(token.TokenType)] + " " + token.AccessToken
}
