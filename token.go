package wallabago

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var token *Token

func setToken(newToken *Token) {
	token = newToken
}

// Token represents the object being returned from the oauth process at the API
// containing the access token, expire time (after converting it from the
// number of seconds the token is valid to the point in time where it will
// expires), type of token, scope and a refresh token
type Token struct {
	AccessToken    string
	ExpirationTime time.Time
	TokenType      string
	Scope          string
	RefreshToken   string
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

// parseTokenResponse consumes a stream (typically a http.Response.Body) that
// is expected to contain JSON that unmarshals into a tokenResponse struct
func parseTokenResponse(reader io.ReadCloser) (*tokenResponse, error) {
	var tokenResponse tokenResponse
	err := json.NewDecoder(reader).Decode(&tokenResponse)
	if err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}

// responseToToken converts a tokenResponse into a Token, computing the point
// in time where the token will expire using the ExpiresIn field in
// tokenResponse and the moment in which the function is called
func responseToToken(tokenResponse *tokenResponse) *Token {
	expiresIn := time.Duration(tokenResponse.ExpiresIn) * time.Second
	expirationTime := time.Now().Add(expiresIn)

	return &Token{
		AccessToken:    tokenResponse.AccessToken,
		ExpirationTime: expirationTime,
		TokenType:      tokenResponse.TokenType,
		Scope:          tokenResponse.Scope,
		RefreshToken:   tokenResponse.RefreshToken,
	}
}

// getToken will use the credentials set in the configuration to
// request an access token from the wallabag API
func getToken() (*tokenResponse, error) {
	tokenURL := Config.WallabagURL + "/oauth/v2/token"
	resp, err := http.PostForm(tokenURL,
		url.Values{
			"grant_type":    {"password"},
			"client_id":     {Config.ClientID},
			"client_secret": {Config.ClientSecret},
			"username":      {Config.UserName},
			"password":      {Config.UserPassword},
		})
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"getToken: bad response from server: %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	return parseTokenResponse(resp.Body)
}

// refreshToken will use the credentials set in the configuration to
// refresh the token stored in a previous request. It errors if there is no
// token stored or if the refresh request fails
func refreshToken() (*tokenResponse, error) {
	if token == nil {
		return nil, fmt.Errorf("A nil token cannot be refreshed")
	}

	tokenURL := Config.WallabagURL + "/oauth/v2/token"
	resp, err := http.PostForm(tokenURL,
		url.Values{
			"grant_type":    {"refresh_token"},
			"client_id":     {Config.ClientID},
			"client_secret": {Config.ClientSecret},
			"refresh_token": {token.RefreshToken},
		})
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"refreshToken: bad response from server: %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	return parseTokenResponse(resp.Body)
}

func checkForToken() error {
	if token == nil {
		tokenResponse, err := getToken()
		if err != nil {
			return err
		}

		setToken(responseToToken(tokenResponse))
		return nil
	}

	if token.ExpirationTime.Before(time.Now()) {
		tokenResponse, err := refreshToken()
		if err != nil {
			return err
		}

		setToken(responseToToken(tokenResponse))
		return err
	}

	return nil
}

// GetAuthTokenHeader will make sure there's a working token and
// return a valid string to be used as an Authentication: header
func GetAuthTokenHeader() (string, error) {
	err := checkForToken()
	if err != nil {
		return "", err
	}

	return strings.Title(token.TokenType) + " " + token.AccessToken, nil
}
