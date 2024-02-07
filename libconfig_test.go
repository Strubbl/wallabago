package wallabago

import "testing"

func TestReadJson(t *testing.T) {
	var tests = []struct {
		input                string
		expectedWallabagURL  string
		expectedClientID     string
		expectedClientSecret string
		expectedUserName     string
		expectedUserPassword string
		expectedIsErrNil     bool
	}{
		{"{\"WallabagURL\": \"http://localhost\", \"ClientId\": \"555_puf29hbu4bnu2\", \"ClientSecret\": \"f2o9uhf32j8fj23fji2huo\", \"UserName\": \"john\", \"UserPassword\": \"passworddd\"}", "http://localhost", "555_puf29hbu4bnu2", "f2o9uhf32j8fj23fji2huo", "john", "passworddd", true},
		{"", "", "", "", "", "", false},
	}
	for _, test := range tests {
		var raw = []byte(test.input)
		c, e := readJSON(raw)
		if c.WallabagURL != test.expectedWallabagURL {
			t.Errorf("readJson(%v): expectedWallabagURL %v, got %v", test.input, test.expectedWallabagURL, c.WallabagURL)
		}
		if c.ClientID != test.expectedClientID {
			t.Errorf("readJson(%v): expectedClientId %v, got %v", test.input, test.expectedClientID, c.ClientID)
		}
		if c.ClientSecret != test.expectedClientSecret {
			t.Errorf("readJson(%v): expectedClientSecret %v, got %v", test.input, test.expectedClientSecret, c.ClientSecret)
		}
		if c.UserName != test.expectedUserName {
			t.Errorf("readJson(%v): expectedUserName %v, got %v", test.input, test.expectedUserName, c.UserName)
		}
		if c.UserPassword != test.expectedUserPassword {
			t.Errorf("readJson(%v): expectedUserPassword %v, got %v", test.input, test.expectedUserPassword, c.UserPassword)
		}
		isErrNil := (e == nil)
		if isErrNil != test.expectedIsErrNil {
			t.Errorf("readJson(%v): expectedIsErrNil %v, got %v", test.input, test.expectedIsErrNil, isErrNil)
		}
	}
}
