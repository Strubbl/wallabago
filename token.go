package wallabago

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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

func getToken() Token {
	tokenURL := Config.WallabagURL + "/oauth/v2/token"
	resp, err := http.PostForm(tokenURL,
		url.Values{"grant_type": {"password"},
			"client_id":     {Config.ClientID},
			"client_secret": {Config.ClientSecret},
			"username":      {Config.UserName},
			"password":      {Config.UserPassword},
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "getToken: getting token failed %s: %v\n", tokenURL, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getToken: error while ioutil.ReadAll %v\n", err)
	}
	//log.Printf("GetToken: body=%v\n", string(body))
	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		fmt.Fprintf(os.Stderr, "getToken: getting token failed %s: %v\n", tokenURL, err)
	}
	return token
}

func checkForToken() {
	if token.TokenType == "" || token.AccessToken == "" {
		token = getToken()
	}
}

func getAuthTokenHeader() string {
	checkForToken()
	return strings.ToUpper(string(token.TokenType[0])) + token.TokenType[1:len(token.TokenType)] + " " + token.AccessToken
}
