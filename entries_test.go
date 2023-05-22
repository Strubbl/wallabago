package wallabago

import "testing"

const oneEntry = `{"page":1,"limit":30,"pages":399,"total":11959,"_links":{"self":{"href":"https:\/\/wallabag.test\/api\/entries?sort=created&order=desc&tags=&since=0&detail=full&page=1&perPage=30"},"first":{"href":"https:\/\/wallabag.test\/api\/entries?sort=created&order=desc&tags=&since=0&detail=full&page=1&perPage=30"},"last":{"href":"https:\/\/wallabag.test\/api\/entries?sort=created&order=desc&tags=&since=0&detail=full&page=399&perPage=30"},"next":{"href":"https:\/\/wallabag.test\/api\/entries?sort=created&order=desc&tags=&since=0&detail=full&page=2&perPage=30"}},"_embedded":{"items":[{"is_archived":1,"is_starred":0,"user_name":"strubbl","user_email":"wallabag.test@wallabag.test","user_id":1,"tags":[{"id":4,"label":"3min","slug":"3min"}],"is_public":false,"id":12327,"uid":null,"title":"Formel 1, Racing Point zittert vor Finale: Viel zu verlieren","url":"https:\/\/www.motorsport-magazin.com\/formel1\/news-268468-formel-1-2020-abu-dhabi-qualifying-rennen-racing-point-zittert-vor-finale-viel-zu-verlieren-sergio-perez-gesamtwertung-lance-stroll-mclaren-mercedes\/","hashed_url":"537b9884cb6d90ec538eb79ef41beb5bcb506b27","origin_url":null,"given_url":"https:\/\/www.motorsport-magazin.com\/formel1\/news-268468-formel-1-2020-abu-dhabi-qualifying-rennen-racing-point-zittert-vor-finale-viel-zu-verlieren-sergio-perez-gesamtwertung-lance-stroll-mclaren-mercedes\/","hashed_given_url":"537b9884cb6d90ec538eb79ef41beb5bcb506b27","archived_at":"2020-12-12T21:33:51+0100","content":"Er baut auf eine weitere starke Performance in der Anfangsphase des Rennens.","created_at":"2020-12-12T21:31:04+0100","updated_at":"2020-12-12T21:33:51+0100","published_at":null,"published_by":null,"starred_at":null,"annotations":[],"mimetype":"text\/html; charset=utf-8","language":"de","reading_time":3,"domain_name":"www.motorsport-magazin.com","preview_picture":"https:\/\/wallabag.test\/assets\/images\/7\/0\/70cc31ae\/a3e0a994.jpeg","http_status":"200","headers":{"server":"nginx\/1.10.3 (Ubuntu)","date":"Sat, 12 Dec 2020 20:31:03 GMT","content-type":"text\/html; charset=utf-8","content-length":"55474","connection":"keep-alive","set-cookie":"PHPSESSID=ed369ef2e34da3dcb696bdf0b94b60e4; path=\/; secure; HttpOnly","expires":"Thu, 19 Nov 1981 08:52:00 GMT","cache-control":"no-store, no-cache, must-revalidate","pragma":"no-cache","vary":"Accept-Encoding","x-storage":"mem","age":"0","x-cache":"MISS","x-cache-hits":"0","x-cache-grace":"none","x-cache-debug":"Main Site","access-control-allow-origin":"https:\/\/ads.motorsport-magazin.com","accept-ranges":"bytes"},"_links":{"self":{"href":"\/api\/entries\/12327"}}}]}}`
const entriesExists = `{"http:\/\/0.0.0.0\/entry10":false,"http:\/\/0.0.0.0\/entry2":false, "http://0.0.0.0/entry3":true}`
const oneItem = `{"is_archived":1,"is_starred":0,"user_name":"strubbl","user_email":"wallabag.test@wallabag.test","user_id":1,"tags":[{"id":4,"label":"3min","slug":"3min"}],"is_public":false,"id":12327,"uid":null,"title":"Formel 1, Racing Point zittert vor Finale: Viel zu verlieren","url":"https:\/\/www.motorsport-magazin.com\/formel1\/news-268468-formel-1-2020-abu-dhabi-qualifying-rennen-racing-point-zittert-vor-finale-viel-zu-verlieren-sergio-perez-gesamtwertung-lance-stroll-mclaren-mercedes\/","hashed_url":"537b9884cb6d90ec538eb79ef41beb5bcb506b27","origin_url":null,"given_url":"https:\/\/www.motorsport-magazin.com\/formel1\/news-268468-formel-1-2020-abu-dhabi-qualifying-rennen-racing-point-zittert-vor-finale-viel-zu-verlieren-sergio-perez-gesamtwertung-lance-stroll-mclaren-mercedes\/","hashed_given_url":"537b9884cb6d90ec538eb79ef41beb5bcb506b27","archived_at":"2020-12-12T21:33:51+0100","content":"Er baut auf eine weitere starke Performance in der Anfangsphase des Rennens.","created_at":"2020-12-12T21:31:04+0100","updated_at":"2020-12-12T21:33:51+0100","published_at":null,"published_by":null,"starred_at":null,"annotations":[],"mimetype":"text\/html; charset=utf-8","language":"de","reading_time":3,"domain_name":"www.motorsport-magazin.com","preview_picture":"https:\/\/wallabag.test\/assets\/images\/7\/0\/70cc31ae\/a3e0a994.jpeg","http_status":"200","headers":{"server":"nginx\/1.10.3 (Ubuntu)","date":"Sat, 12 Dec 2020 20:31:03 GMT","content-type":"text\/html; charset=utf-8","content-length":"55474","connection":"keep-alive","set-cookie":"PHPSESSID=ed369ef2e34da3dcb696bdf0b94b60e4; path=\/; secure; HttpOnly","expires":"Thu, 19 Nov 1981 08:52:00 GMT","cache-control":"no-store, no-cache, must-revalidate","pragma":"no-cache","vary":"Accept-Encoding","x-storage":"mem","age":"0","x-cache":"MISS","x-cache-hits":"0","x-cache-grace":"none","x-cache-debug":"Main Site","access-control-allow-origin":"https:\/\/ads.motorsport-magazin.com","accept-ranges":"bytes"},"_links":{"self":{"href":"\/api\/entries\/12327"}}}`

func TestGetEntries(t *testing.T) {
	expectedLimit := 30
	expectedTotal := 11959
	expectedPage := 1
	expectedPages := 399
	entries, _ := GetEntries(mockGetOneEntry, 0, 0, "", "", 0, expectedLimit, "", 0, -1, "", "")
	if entries.Total != expectedTotal {
		t.Errorf("expected %v total entries, but got %v", expectedTotal, entries.Total)
	}
	if entries.Page != expectedPage {
		t.Errorf("expected %v page, but got %v", expectedPage, entries.Page)
	}
	if entries.Pages != expectedPages {
		t.Errorf("expected %v pages, but got %v", expectedPages, entries.Pages)
	}
	if entries.Limit != expectedLimit {
		t.Errorf("expected %v limit, but got %v", expectedLimit, entries.Limit)
	}
}

func mockGetOneEntry(url string, httpMethod string, postData []byte) ([]byte, error) {
	return []byte(oneEntry), nil
}

func TestGetEntriesExists(t *testing.T) {
	urls := []string{"http://0.0.0.0/entry10", "http://0.0.0.0/entry2", "http://0.0.0.0/entry3"}
	existsResult := []bool{false, false, true}
	e, _ := GetEntriesExists(mockGetEntriesExists, urls)
	for index := 0; index < len(urls); index++ {
		if e[urls[index]] != existsResult[index] {
			t.Errorf("for url %v expected exists=%v, but got %v", urls[index], e[urls[index]], existsResult[index])
		}
	}
}

func mockGetEntriesExists(url string, httpMethod string, postData []byte) ([]byte, error) {
	return []byte(entriesExists), nil
}

func TestGetEntry(t *testing.T) {
	expectedID := 12327
	expectedIsArchived := 1
	expectedTitle := "Formel 1, Racing Point zittert vor Finale: Viel zu verlieren"
	e, _ := GetEntry(mockGetEntry, expectedID)
	if e.ID != expectedID {
		t.Errorf("expected id=%v, but got %v", expectedID, e.ID)
	}
	if e.IsArchived != expectedIsArchived {
		t.Errorf("expected is_archive=%v, but got %v", expectedIsArchived, e.IsArchived)
	}
	if e.Title != expectedTitle {
		t.Errorf("expected is_archive=%v, but got %v", expectedTitle, e.Title)
	}
}

func mockGetEntry(url string, httpMethod string, postData []byte) ([]byte, error) {
	return []byte(oneItem), nil
}
