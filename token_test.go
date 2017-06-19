package wallabago

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAuthTokenHeader(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()
	wbgcfg := NewWallabagConfig(server.URL, "asdf", "ghkj", "wallabago", "555nase")
	SetConfig(wbgcfg)
	fmt.Printf("Config: %v, %v, %v, %v, %v\n", Config.WallabagURL, Config.ClientID, Config.ClientSecret, Config.UserName, Config.UserPassword)
	mux.HandleFunc("/oauth/v2/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"access_token":"294hf92ufurjfgoiqjfioj4","refresh_token": "ZGE5MDg3ZTNjNmNkYTY0ZWZh","expires_in":3600,"scope": "null", "token_type": "bearer"}`)
	})

	token := GetAuthTokenHeader()
	expectedToken := "Bearer 294hf92ufurjfgoiqjfioj4"
	if token != expectedToken {
		t.Errorf("GetAuthTokenHeader(): expected %v, got %v", expectedToken, token)
	}
}
