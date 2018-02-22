package wallabago

import "testing"

const annotationsResult = `{"total":2,"rows":[{"annotator_schema_version":"v1.0","id":29,"text":"Anmerkung 1","created_at":"2017-06-19T21:54:37+0200","updated_at":"2017-06-19T21:54:37+0200","quote":"Bundesnetzagentur nicht zweifelsfrei","ranges":[{"start":"\/div[1]\/p[2]","startOffset":28,"end":"\/div[1]\/p[2]","endOffset":65}]},{"annotator_schema_version":"v1.0","id":30,"text":"Anmerkung 2","created_at":"2017-06-19T21:54:45+0200","updated_at":"2017-06-19T21:54:45+0200","quote":"werden m\u00fcssten","ranges":[{"start":"\/div[1]\/p[3]","startOffset":158,"end":"\/div[1]\/p[3]","endOffset":172}]}]}`

func TestGetAnnotations(t *testing.T) {
	articleID := 3977
	expectedTotal := 2
	expectedText := "Anmerkung 1"
	expectedQuote := "Bundesnetzagentur nicht zweifelsfrei"
	expectedRangeStartOffset := 28
	annos, err := GetAnnotations(mockGetAnnotations, articleID)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	if annos.Total != expectedTotal {
		t.Errorf("expected id=%v, but got %v", expectedTotal, annos.Total)
	}
	if annos.Rows[0].Text != expectedText {
		t.Errorf("expected id=%v, but got %v", expectedText, annos.Rows[0].Text)
	}
	if annos.Rows[0].Quote != expectedQuote {
		t.Errorf("expected id=%v, but got %v", expectedQuote, annos.Rows[0].Quote)
	}
	if annos.Rows[0].Ranges[0].StartOffset != expectedRangeStartOffset {
		t.Errorf("expected id=%v, but got %v", expectedRangeStartOffset, annos.Rows[0].Ranges[0].StartOffset)
	}
}

func mockGetAnnotations(url string, httpMethod string, postData []byte) []byte {
	return []byte(annotationsResult)
}
