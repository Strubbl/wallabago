package wallabago

import (
	"encoding/json"
	"strconv"
)

// Annotations represents an annotation result API call for an article
type Annotations struct {
	Rows  []Annotation `json:"rows"`
	Total int          `json:"total"`
}

// Annotation represents one annotation made to an article
type Annotation struct {
	AnnotatorSchemaVersion string       `json:"annotator_schema_version"`
	CreatedAt              WallabagTime `json:"created_at"`
	ID                     int          `json:"id"`
	Quote                  string       `json:"quote"`
	Ranges                 []Range      `json:"ranges"`
	Text                   string       `json:"text"`
	UpdatedAt              WallabagTime `json:"updated_at"`
}

// Range represents the text borders of an annotation
type Range struct {
	End         string      `json:"end"`
	EndOffset   interface{} `json:"endOffset"`
	Start       string      `json:"start"`
	StartOffset interface{} `json:"startOffset"`
}

// GetAnnotations queries the API for all annotations of an article according to /api/annotations/ID
func GetAnnotations(bodyByteGetterFunc BodyByteGetter, articleID int) (Annotations, error) {
	var annotations Annotations
	annoURL := Config.WallabagURL + "/api/annotations/" + strconv.Itoa(articleID) + ".json"
	body, err := bodyByteGetterFunc(annoURL, "GET", nil)
	if err != nil {
		return annotations, err
	}
	err = json.Unmarshal(body, &annotations)
	return annotations, err
}
