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
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getBodyOfURL: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	log.Print(resp.Status)
	//fmt.Printf("%s", b)
	return string(b)
}

func getBodyOfAPIURL(url string) []byte {
	if token.TokenType == "" || token.AccessToken == "" {
		token = getToken()
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	authString := strings.ToUpper(string(token.TokenType[0])) + token.TokenType[1:len(token.TokenType)] + " " + token.AccessToken
	// log.Print("getBodyOfAPIURL: authString=" + authString)
	req.Header.Add("Authorization", authString)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getBodyOfAPIURL: error while client.Do %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getBodyOfAPIURL: error while ioutil.ReadAll %v\n", err)
	}
	return body
}
