package wallabago

import (
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	expected := LoggedInUser{
		ID:        1,
		UserName:  "john",
		Email:     "no-reply@example.com",
		CreatedAt: &WallabagTime{time.Date(2011, 1, 1, 21, 31, 5, 0, time.FixedZone("+0000", 0))},
		UpdatedAt: &WallabagTime{time.Date(2024, 2, 1, 16, 16, 17, 0, time.FixedZone("+0000", 0))},
	}
	got, _ := User(mockGetBodyOfUserApiCall)
	if expected.ID != got.ID {
		t.Errorf("expected ID %v but got %v", expected.ID, got.ID)
	}
	if expected.UserName != got.UserName {
		t.Errorf("expected UserName %v but got %v", expected.UserName, got.UserName)
	}
	if expected.Email != got.Email {
		t.Errorf("expected Email %v but got %v", expected.Email, got.Email)
	}
	if !expected.CreatedAt.Equal(got.CreatedAt) {
		t.Errorf("expected CreatedAt:\n%v\nbut got:\n%v", expected.CreatedAt, got.CreatedAt)
	}
	if !expected.UpdatedAt.Equal(got.UpdatedAt) {
		t.Errorf("expected UpdatedAt:\n%v\nbut got:\n%v", expected.UpdatedAt, got.UpdatedAt)
	}
	if expected.CreatedAt.Equal(got.UpdatedAt) {
		t.Errorf("expected CreatedAt:\n%v\nto be different than UpdatedAt but it's equal:\n%v", expected.CreatedAt, got.UpdatedAt)
	}
}

func mockGetBodyOfUserApiCall(url string, httpMethod string, postData []byte) ([]byte, error) {
	return []byte(`{"id":1,"username":"john","email":"no-reply@example.com","created_at":"2011-01-01T21:31:05+0000","updated_at":"2024-02-01T16:16:17+0000"}`), nil
}
