package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	log.SetOutput(os.Stdout)
	log.Print("fetch")
	loginPage := getBodyOfURL("https://wallabag.linux4tw.de")
	fmt.Printf("%s\n", loginPage)
	log.Printf("csrfToken: %v\n", getCSRFToken(loginPage))
	log.Printf("time elapsed: %.2fs\n", time.Since(start).Seconds())
}

// makes a HTTP request and returns the HTML code of that URL
func getBodyOfURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	log.Print(resp.Status)
	//fmt.Printf("%s", b)
	return string(b)
}

// parses the HTML code from the wallabag v2 login page to retrieve the CSRF token and returns that
func getCSRFToken(htmlCode string) string {
	const csrfTokenStartRegex string = "<input type=\"hidden\" name=\"_csrf_token\" value=\""
	csrfTokenStart := len(csrfTokenStartRegex) + strings.Index(htmlCode, csrfTokenStartRegex)
	log.Printf("csrfTokenStart: %d\n", csrfTokenStart)
	if csrfTokenStart < 0 {
		log.Print("csrf token not found in login page")
		os.Exit(1)
	}
	csrfTokenEnd := csrfTokenStart + strings.Index(htmlCode[csrfTokenStart:], "\" />")
	log.Printf("csrfTokenEnd: %d\n", csrfTokenEnd)
	csrfToken := htmlCode[csrfTokenStart:csrfTokenEnd]
	log.Printf("csrfToken: %v\n", csrfToken)
	return csrfToken
}
