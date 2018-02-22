package wallabago

import (
	"encoding/json"
	"strconv"
)

// Tag represents one tag with its properties
type Tag struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
	Slug  string `json:"slug"`
}

// GetTagsOfEntry queries the API for the tags of an article /entries/ID
func GetTagsOfEntry(bodyByteGetterFunc BodyByteGetter, articleID int) ([]Tag, error) {
	entryTagsURL := Config.WallabagURL + "/api/entries/" + strconv.Itoa(articleID) + "/tags.json"
	body := bodyByteGetterFunc(entryTagsURL, "GET", nil)
	var tags []Tag
	err := json.Unmarshal(body, &tags)
	return tags, err
}

// GetTags queries the API for all tags in wallabag /tags
func GetTags(bodyByteGetterFunc BodyByteGetter) ([]Tag, error) {
	tagsURL := Config.WallabagURL + "/api/tags"
	body := bodyByteGetterFunc(tagsURL, "GET", nil)
	var tags []Tag
	err := json.Unmarshal(body, &tags)
	return tags, err
}
