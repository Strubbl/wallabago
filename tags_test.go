package wallabago

import "testing"

const tagsResult = `[{"id":53,"label":"2min","slug":"2min"},{"id":57,"label":"android","slug":"android"},{"id":58,"label":"linux","slug":"linux"}]`

func TestGetTagsOfEntry(t *testing.T) {
	articleID := 3977
	expectedID := 57
	expectedLabel := "android"
	expectedSlug := "android"
	tags := GetTagsOfEntry(mockGetTagsOfEntry, articleID)
	if tags[1].ID != expectedID {
		t.Errorf("expected id=%v, but got %v", expectedID, tags[1].ID)
	}
	if tags[1].Label != expectedLabel {
		t.Errorf("expected label=%v, but got %v", expectedLabel, tags[1].Label)
	}
	if tags[1].Slug != expectedSlug {
		t.Errorf("expected slug=%v, but got %v", expectedSlug, tags[1].Slug)
	}
}

func mockGetTagsOfEntry(url string, httpMethod string, postData []byte) []byte {
	return []byte(tagsResult)
}
