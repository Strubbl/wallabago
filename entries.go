package wallabago

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Entries represents the object being returned from the API request /entries
type Entries struct {
	Page      int
	Limit     int
	Pages     int
	Total     int
	NaviLinks Links    `json:"_links"`
	Embedded  Embedded `json:"_embedded"`
}

// Embedded items in the API request
type Embedded struct {
	Items []Item `json:"items"`
}

// Item represents individual items in API responses
type Item struct {
	Links          Links         `json:"_links"`
	Annotations    []Annotation  `json:"annotations"`
	ArchivedAt     *WallabagTime `json:"archived_at"`
	CreatedAt      *WallabagTime `json:"created_at"`
	Content        string        `json:"content"`
	DomainName     string        `json:"domain_name"`
	GivenURL       string        `json:"given_url"`
	HashedGivenURL string        `json:"hashed_given_url"`
	HashedURL      string        `json:"hashed_url"`
	ID             int           `json:"id"`
	IsArchived     int           `json:"is_archived"`
	IsPublic       bool          `json:"is_public"`
	IsStarred      int           `json:"is_starred"`
	Language       string        `json:"language"`
	Mimetype       string        `json:"mimetype"`
	OriginURL      string        `json:"origin_url"`
	PreviewPicture string        `json:"preview_picture"`
	PublishedAt    *WallabagTime `json:"published_at"`
	PublishedBy    []string      `json:"published_by"`
	ReadingTime    int           `json:"reading_time"`
	StarredAt      *WallabagTime `json:"starred_at"`
	Tags           []Tag         `json:"tags"`
	Title          string        `json:"title"`
	UID            string        `json:"uid"`
	UpdatedAt      *WallabagTime `json:"updated_at"`
	URL            string        `json:"url"`
	UserEmail      string        `json:"user_email"`
	UserID         int           `json:"user_id"`
	UserName       string        `json:"user_name"`
}

// WallabagTimeLayout is a variation of RFC3339 but without colons in
// the timezone delimeter, breaking the RFC
const WallabagTimeLayout = "2006-01-02T15:04:05-0700"

// WallabagTime overrides builtin time to allow for custom time parsing
type WallabagTime struct {
	time.Time
}

// UnmarshalJSON parses the custom date format wallabag returns
func (t *WallabagTime) UnmarshalJSON(buf []byte) (err error) {
	s := strings.Trim(string(buf), `"`)
	if s == "null" {
		t.Time = time.Time{}
		return err
	}
	t.Time, err = time.Parse(WallabagTimeLayout, s)
	if err != nil {
		t.Time = time.Time{}
		return err
	}
	return err
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
func GetEntries(bodyByteGetterFunc BodyByteGetter, archive int, starred int, sort string, order string, page int, perPage int, tags string, since int, public int, detail string, domain_name string) (Entries, error) {
	var e Entries
	entriesURL := Config.WallabagURL + "/api/entries.json?"
	if archive == 0 || archive == 1 {
		entriesURL += "archive=" + strconv.Itoa(archive) + "&"
	}
	if starred == 0 || starred == 1 {
		entriesURL += "starred=" + strconv.Itoa(starred) + "&"
	}
	if sort == "created" || sort == "updated" || sort == "archived" {
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
	if since > 0 {
		entriesURL += "since=" + strconv.Itoa(since) + "&"
	}
	if public >= 0 && (public == 0 || public == 1) {
		entriesURL += "public=" + strconv.Itoa(public) + "&"
	}
	if detail == "metadata" || detail == "full" {
		entriesURL += "detail=" + detail + "&"
	}
	if domain_name != "" {
		domain_name_encoded := url.QueryEscape(domain_name)
		entriesURL += "domain_name=" + domain_name_encoded + "&"
	}

	//log.Printf("getEntries: entriesURL=%s", entriesURL)
	body, err := bodyByteGetterFunc(entriesURL, "GET", nil)
	if err != nil {
		return e, err
	}
	//log.Printf("getEntries: body=\n%v\n\n\n\n", string(body))
	err = json.Unmarshal(body, &e)
	return e, err
}

// GetAllEntries calls GetEntries with no parameters, thus using the default values of the API request /entries and returning all articles as []wallabago.Item
func GetAllEntries() ([]Item, error) {
	page := -1
	perPage := -1
	e, err := GetEntries(APICall, -1, -1, "", "", page, perPage, "", 0, -1, "", "")
	if err != nil {
		log.Println("GetAllEntries: first GetEntries call failed", err)
		return nil, err
	}
	allEntries := e.Embedded.Items
	if e.Total > len(allEntries) {
		secondPage := e.Page + 1
		perPage = e.Limit
		pages := e.Pages
		for i := secondPage; i <= pages; i++ {
			e, err := GetEntries(APICall, -1, -1, "", "", i, perPage, "", 0, -1, "", "")
			if err != nil {
				log.Printf("GetAllEntries: GetEntries for page %d failed: %v", i, err)
				return nil, err
			}
			tmpAllEntries := e.Embedded.Items
			allEntries = append(allEntries, tmpAllEntries...)
		}
	}
	return allEntries, err
}

// GetNumberOfTotalArticles returns the number of all articles saved in wallabag
func GetNumberOfTotalArticles() (int, error) {
	e, err := GetEntries(APICall, -1, -1, "", "", -1, -1, "", 0, -1, "", "")
	if err != nil {
		return -1, err
	}
	return e.Total, err
}

// GetNumberOfArchivedArticles returns the number of archived articles in wallabag
func GetNumberOfArchivedArticles() (int, error) {
	e, err := GetEntries(APICall, 1, -1, "", "", -1, -1, "", 0, -1, "", "")
	if err != nil {
		return -1, err
	}
	return e.Total, err
}

// GetNumberOfStarredArticles returns the number of starred articles in wallabag (including unread and archived starred ones)
func GetNumberOfStarredArticles() (int, error) {
	e, err := GetEntries(APICall, -1, 1, "", "", -1, -1, "", 0, -1, "", "")
	if err != nil {
		return -1, err
	}
	return e.Total, err
}

// PostEntry creates a new article in wallabag
func PostEntry(url, title, tags string, starred, archive int) error {
	postData := map[string]string{
		"url":     url,
		"title":   title,
		"tags":    tags,
		"starred": strconv.Itoa(starred),
		"archive": strconv.Itoa(archive),
	}
	postDataJSON, err := json.Marshal(postData)
	if err != nil {
		return err
	}
	entriesURL := Config.WallabagURL + "/api/entries.json"
	_, err = APICall(entriesURL, "POST", postDataJSON)
	//log.Println("PostEntry: response:", string(body))
	return err
}

// GetEntriesExists queries the API for articles according to the API request /entries/exists
// it checks if the urls in the given array exist
// returns a map with the URL as key and the result as value
func GetEntriesExists(bodyByteGetterFunc BodyByteGetter, urls []string) (map[string]bool, error) {
	var m map[string]bool
	entriesExistsURL := Config.WallabagURL + "/api/entries/exists.json?"
	if len(urls) > 0 {
		for i := 0; i < len(urls); i++ {
			entriesExistsURL += "urls[]=" + urls[i] + "&"
		}
	}
	body, err := bodyByteGetterFunc(entriesExistsURL, "GET", nil)
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(body, &m)
	return m, err
}

// GetEntry queries the API for a specific article according to the API request /entries/ID
func GetEntry(bodyByteGetterFunc BodyByteGetter, articleID int) (Item, error) {
	var item Item
	entryURL := Config.WallabagURL + "/api/entries/" + strconv.Itoa(articleID) + ".json"
	body, err := bodyByteGetterFunc(entryURL, "GET", nil)
	if err != nil {
		return item, err
	}
	err = json.Unmarshal(body, &item)
	return item, err
}

// ExportEntry queries the API to retrieve a single entry in a predefined format according to the API request /entries/ID/export.FORMAT
func ExportEntry(bodyByteGetterFunc BodyByteGetter, articleID int, format string) ([]byte, error) {
	entryURL := Config.WallabagURL + "/api/entries/" + strconv.Itoa(articleID) + "/export." + format
	body, err := bodyByteGetterFunc(entryURL, "GET", nil)
	if err != nil {
		return nil, err
	}
	return body, err
}
