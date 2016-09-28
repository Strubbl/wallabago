package wallabago

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string
	RefreshToken string `json:"refresh_token"`
}

func GetToken() Token {
	tokenURL := WallabagURL + "/oauth/v2/token"
	resp, err := http.PostForm(tokenURL,
		url.Values{"grant_type": {"password"},
			"client_id":     {ClientId},
			"client_secret": {ClientSecret},
			"username":      {UserName},
			"password":      {UserPassword},
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
