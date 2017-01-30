package wallabago

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Entries represents the object being returned from the API request /entries
type Entries struct {
	Page      int
	Limit     int
	Pages     int
	Total     int
	NaviLinks *Links `json:"_links"`
}

// Links contains four links (self, first, last, next), being part of the Entries object
type Links struct {
	Self  *Link
	First *Link
	Last  *Link
	Next  *Link
}

// Link object consists of its URL
type Link struct {
	Href string
}

// GetEntries queries the API for articles according to the API request /entries
func GetEntries(archive int, starred int, sort string, order string, page int, perPage int, tags string) Entries {
	entriesURL := Config.WallabagURL + "/api/entries.json?"
	if archive == 0 || archive == 1 {
		entriesURL += "archive=" + strconv.Itoa(archive) + "&"
	}
	if starred == 0 || starred == 1 {
		entriesURL += "starred=" + strconv.Itoa(starred) + "&"
	}
	if sort == "created" || sort == "updated" {
		entriesURL += "sort=" + sort + "&"
	}
	if order == "asc" || order == "desc" {
		entriesURL += "order=" + order + "&"
	}
	if page > 0 {
		entriesURL += "page=" + strconv.Itoa(page) + "&"
	}
	if perPage > 0 {
		entriesURL += "perPage=" + strconv.Itoa(perPage) + "&"
	}
	if tags != "" {
		entriesURL += "tags=" + tags + "&"
	}

	//log.Printf("getEntries: entriesURL=%s", entriesURL)
	body := getBodyOfAPIURL(entriesURL)
	//log.Printf("getEntries: body=\n%v\n", string(body))
	var e Entries
	if err := json.Unmarshal(body, &e); err != nil {
		fmt.Fprintf(os.Stderr, "getEntries: json unmarshal failed: %v\n", err)
	}
	return e
}

// GetAllEntries calls GetEntries with no parameters, thus using the default values of the API request /entries and returning all articles, of course not all at once, but limitted to page through
func GetAllEntries() Entries {
	return GetEntries(-1, -1, "", "", -1, -1, "")
}

// GetNumberOfTotalArticles returns the number of all articles saved in wallabag
func GetNumberOfTotalArticles() int {
	return GetAllEntries().Total
}

// GetNumberOfArchivedArticles returns the number of archived articles in wallabag
func GetNumberOfArchivedArticles() int {
	return GetEntries(1, -1, "", "", -1, -1, "").Total
}

// GetNumberOfStarredArticles returns the number of starred articles in wallabag (including unread and archived starred ones)
func GetNumberOfStarredArticles() int {
	return GetEntries(-1, 1, "", "", -1, -1, "").Total
}
