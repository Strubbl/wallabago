package wallabago

import "testing"

func TestInfo(t *testing.T) {
	expected := Information{Appname: "wallabag", Version: "2.6.8", AllowedRegistration: false}
	got, _ := Info(mockGetBodyOfInfoApiCall)
	if expected != got {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func mockGetBodyOfInfoApiCall(url string, httpMethod string, postData []byte) ([]byte, error) {
	return []byte(`{"appname":"wallabag","version":"2.6.8","allowed_registration":false}`), nil
}
