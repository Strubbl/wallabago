package wallabago

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var config WallabagConfig
var token Token

// WallabagConfig contains all data needed to connect to wallabag API like URL, id and secret of the API client and user name and according password
type WallabagConfig struct {
	WallabagURL  string
	ClientID     string
	ClientSecret string
	UserName     string
	UserPassword string
}

// Token represents the object being returned from the oauth process at the API containing the access token, expire time, type of token, scope and a refresh token
type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string
	RefreshToken string `json:"refresh_token"`
}

func getToken() Token {
	tokenURL := config.WallabagURL + "/oauth/v2/token"
	resp, err := http.PostForm(tokenURL,
		url.Values{"grant_type": {"password"},
			"client_id":     {config.ClientID},
			"client_secret": {config.ClientSecret},
			"username":      {config.UserName},
			"password":      {config.UserPassword},
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "token: getting token failed %s: %v\n", tokenURL, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//log.Printf("GetToken: body=%v\n", string(body))
	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		fmt.Fprintf(os.Stderr, "token: getting token failed %s: %v\n", tokenURL, err)
	}
	return token
}
