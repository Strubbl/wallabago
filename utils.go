package wallabago

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

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

func getBodyOfAPIURL(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	authString := getAuthTokenHeader()
	// log.Print("getBodyOfAPIURL: authString=" + authString)
	req.Header.Add("Authorization", authString)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getBodyOfAPIURL: error while client.Do %v\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getBodyOfAPIURL: error while ioutil.ReadAll %v\n", err)
	}
	return body
}

func postToAPI(apiURL string, postData []byte) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(string(postData)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "postToAPI: posting form %v\n", err)
	}

	// auth
	authString := getAuthTokenHeader()
	req.Header.Add("Authorization", authString)
	// exec POST request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "postToAPI: error while getting response of our POST request %v\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "postToAPI: error while ioutil.ReadAll %v\n", err)
	}
	return body
}
