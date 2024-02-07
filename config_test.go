package wallabago

import "testing"

func TestConfig(t *testing.T) {
	expected := Configuration{ID: 1, ItemsPerPage: 9, Language: "de", FeedToken: "asdfgh89zzzh123", FeedLimit: 100, ReadingSpeed: 208.0, ActionMarkAsRead: 1, ListMode: 0, DisplayThumbnails: 1}
	got, _ := Config(mockGetBodyOfConfigApiCall)
	if expected != got {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func mockGetBodyOfConfigApiCall(url string, httpMethod string, postData []byte) ([]byte, error) {
	return []byte(`{"id":1,"items_per_page":9,"language":"de","feed_token":"asdfgh89zzzh123","feed_limit":100,"reading_speed":208.0,"action_mark_as_read":1,"list_mode":0,"display_thumbnails":1}`), nil
}
