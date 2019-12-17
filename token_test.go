package wallabago

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetToken(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()
	wbgcfg := NewWallabagConfig(server.URL, "asdf", "ghkj", "wallabago", "555nase")
	SetConfig(wbgcfg)
	mux.HandleFunc("/oauth/v2/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"access_token":"294hf92ufurjfgoiqjfioj4","refresh_token": "ZGE5MDg3ZTNjNmNkYTY0ZWZh","expires_in":3600,"scope": "null", "token_type": "bearer"}`)
	})

	token, err := getToken()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	expectedToken := tokenResponse{"294hf92ufurjfgoiqjfioj4", 3600, "bearer", "null", "ZGE5MDg3ZTNjNmNkYTY0ZWZh"}
	if *token != expectedToken {
		t.Errorf("TestGetToken(): expected %v, got %v", expectedToken, token)
	}
}

func TestRefreshToken(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()
	wbgcfg := NewWallabagConfig(server.URL, "asdf", "ghkj", "wallabago", "555nase")
	SetConfig(wbgcfg)
	mux.HandleFunc("/oauth/v2/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"access_token":"294hf92ufurjfgoiqjfioj4","refresh_token": "ZGE5MDg3ZTNjNmNkYTY0ZWZh","expires_in":3600,"scope": "null", "token_type": "bearer"}`)
	})

	// Set the token first so we can refresh it later
	token, err := getToken()
	setToken(responseToToken(token))
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	refreshedToken, err := refreshToken()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	expectedToken := tokenResponse{"294hf92ufurjfgoiqjfioj4", 3600, "bearer", "null", "ZGE5MDg3ZTNjNmNkYTY0ZWZh"}
	if *refreshedToken != expectedToken {
		t.Errorf("TestGetToken(): expected %v, got %v", expectedToken, token)
	}
}

func TestGetAuthTokenHeader(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()
	wbgcfg := NewWallabagConfig(server.URL, "asdf", "ghkj", "wallabago", "555nase")
	SetConfig(wbgcfg)
	mux.HandleFunc("/oauth/v2/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"access_token":"294hf92ufurjfgoiqjfioj4","refresh_token": "ZGE5MDg3ZTNjNmNkYTY0ZWZh","expires_in":3600,"scope": "null", "token_type": "bearer"}`)
	})

	authTokenHeader, err := GetAuthTokenHeader()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	expectedAuthTokenHeader := "Bearer 294hf92ufurjfgoiqjfioj4"
	if authTokenHeader != expectedAuthTokenHeader {
		t.Errorf("TestGetAuthTokenHeader(): expected %v, got %v", expectedAuthTokenHeader, authTokenHeader)
	}
}
