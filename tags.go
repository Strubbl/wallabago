package wallabago

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Tag represents one tag with its properties
type Tag struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
	Slug  string `json:"slug"`
}

// GetTagsOfEntry queries the API for the tags of an article /entries/ID
func GetTagsOfEntry(bodyByteGetterFunc BodyByteGetter, articleID int) []Tag {
	entryTagsURL := Config.WallabagURL + "/api/entries/" + strconv.Itoa(articleID) + "/tags.json"
	body := bodyByteGetterFunc(entryTagsURL, "GET", nil)
	var tags []Tag
	if err := json.Unmarshal(body, &tags); err != nil {
		fmt.Fprintf(os.Stderr, "GetTagsOfEntry: json unmarshal failed: %v\n", err)
	}
	return tags
}
