package wallabago

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// BodyStringGetter represents a function returning the body of an HTTP response as string
type BodyStringGetter func(url string, httpMethod string, postData []byte) (string, error)

// BodyByteGetter represents a function returning the body of an HTTP response as byte array
type BodyByteGetter func(url string, httpMethod string, postData []byte) ([]byte, error)

// APICall authenticates to wallabag instane before issuing the HTTP request
func APICall(apiURL string, httpMethod string, postData []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(httpMethod, apiURL, strings.NewReader(string(postData)))
	if err != nil {
		log.Printf("APICall: creating request failed with error: %v\n", err)
		return nil, err
	}
	// auth
	authString, err := GetAuthTokenHeader()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", authString)
	req.Header.Add("Content-Type", "application/json")
	// exec API request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("APICall: error while getting response of our API request %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("APICall: error while io.ReadAll %v\n", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("error while API call, status code not ok, but instead: %d %s\n", resp.StatusCode, resp.Status)
		return nil, err
	}
	return body, err
}
