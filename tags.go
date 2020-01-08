package wallabago

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Tag represents one tag with its properties
type Tag struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
	Slug  string `json:"slug"`
}

// GetTagsOfEntry queries the API for the tags of an article /entries/ID
func GetTagsOfEntry(bodyByteGetterFunc BodyByteGetter, articleID int) ([]Tag, error) {
	var tags []Tag
	entryTagsURL := Config.WallabagURL + "/api/entries/" + strconv.Itoa(articleID) + "/tags.json"
	body, err := bodyByteGetterFunc(entryTagsURL, "GET", nil)
	if err != nil {
		return tags, err
	}
	err = json.Unmarshal(body, &tags)
	return tags, err
}

// GetTags queries the API for all tags in wallabag /tags
func GetTags(bodyByteGetterFunc BodyByteGetter) ([]Tag, error) {
	var tags []Tag
	tagsURL := Config.WallabagURL + "/api/tags"
	body, err := bodyByteGetterFunc(tagsURL, "GET", nil)
	if err != nil {
		return tags, err
	}
	err = json.Unmarshal(body, &tags)
	return tags, err
}

// AddEntryTags add tags to an entry
func AddEntryTags(entry int, tags ...string) error {
	url := Config.WallabagURL + "/api/entries/" + strconv.Itoa(entry) + "/tags.json"
	data := map[string]string{
		"tags": strings.Join(tags, ","),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = APICall(url, "POST", jsonData)
	return err
}

// DeleteEntryTag removes a tag from an entry
func DeleteEntryTag(entry int, tag int) error {
	url := Config.WallabagURL + "/api/entries/" + strconv.Itoa(entry) + "/tags/" + strconv.Itoa(tag) + ".json"
	_, err := APICall(url, "DELETE", nil)
	return err
}
