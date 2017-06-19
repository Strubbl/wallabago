package wallabago

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// BodyStringGetter represents a function returning the body of an HTTP response as string
type BodyStringGetter func(url string, httpMethod string, postData []byte) string

// BodyByteGetter represents a function returning the body of an HTTP response as byte array
type BodyByteGetter func(url string, httpMethod string, postData []byte) []byte

// makes a HTTP request and returns the HTML code of that URL
func getBodyOfURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getBodyOfURL: %v\n", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getBodyOfURL: reading %s: %v\n", url, err)
	}
	log.Print(resp.Status)
	//fmt.Printf("%s", b)
	return string(b)
}

// APICall authenticates to wallabag instane before issuing the HTTP request
func APICall(apiURL string, httpMethod string, postData []byte) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(httpMethod, apiURL, strings.NewReader(string(postData)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "APICall: creating request failed with error: %v\n", err)
	}
	// auth
	authString := GetAuthTokenHeader()
	req.Header.Add("Authorization", authString)
	// exec API request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "APICall: error while getting response of our API request %v\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "APICall: error while ioutil.ReadAll %v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error while API call, status code not ok, but instead: %d %s", resp.StatusCode, resp.Status)
	}
	return body
}
