package wallabago

import (
	"encoding/json"
	"fmt"
	"os"
)

type Entries struct {
	Page      int
	Limit     int
	Pages     int
	Total     int
	NaviLinks *Links `json:"_links"`
}

type Links struct {
	Self  *Link
	First *Link
	Last  *Link
	Next  *Link
}

type Link struct {
	Href string
}

func getEntries(archive int, starred int, sort string, order string, page int, perPage int, tags string) Entries {
	entriesURL := WallabagURL + "/api/entries.json?"
	if archive == 0 || archive == 1 {
		entriesURL += "archive=" + string(archive) + "&"
	}
	if starred == 0 || starred == 1 {
		entriesURL += "starred=" + string(starred) + "&"
	}
	if sort == "created" || sort == "updated" {
		entriesURL += "sort=" + sort + "&"
	}
	if order == "asc" || order == "desc" {
		entriesURL += "order=" + order + "&"
	}
	if page > 0 {
		entriesURL += "page=" + string(page) + "&"
	}
	if perPage > 0 {
		entriesURL += "perPage=" + string(perPage) + "&"
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

func getAllEntries() Entries {
	return getEntries(-1, -1, "", "", -1, -1, "")
}

func getNumberOfTotalArticles() int {
	return getAllEntries().Total
}

func getNumberOfArchivedArticles() int {
	return getEntries(1, -1, "", "", -1, -1, "").Total
}

func getNumberOfStarredArticles() int {
	return getEntries(-1, 1, "", "", -1, -1, "").Total
}
