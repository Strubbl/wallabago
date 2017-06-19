package wallabago

import "testing"

const oneEntry = `{"page":1,"limit":1,"pages":3494,"total":3494,"_links":{"self":{"href":"https:\/\/wallabag.org\/api\/entries?sort=created&order=desc&tags=&since=0&page=1&perPage=1"},"first":{"href":"https:\/\/wallabag.org\/api\/entries?sort=created&order=desc&tags=&since=0&page=1&perPage=1"},"last":{"href":"https:\/\/wallabag.org\/api\/entries?sort=created&order=desc&tags=&since=0&page=3494&perPage=1"},"next":{"href":"https:\/\/wallabag.org\/api\/entries?sort=created&order=desc&tags=&since=0&page=2&perPage=1"}},"_embedded":{"items":[{"is_archived":1,"is_starred":0,"user_name":"wallabagotest","user_email":"wallabagotest@wallabag.org","user_id":1,"tags":[{"id":61,"label":"1min","slug":"1min"}],"id":3784,"title":"VfB Suhl LOTTO Th\u00fcringen \u00bb Landkreis Schmalkalden-Meiningen unterst\u00fctzt Profi-Volleyball in S\u00fcdth\u00fcringen","url":"http:\/\/1.bundesliga.vfb-suhl.de\/landkreis-schmalkalden-meiningen-unterstuetzt-profi-volleyball-in-suedthueringen\/","content":"<p>Der S\u00fcdth\u00fcringer Bundesligist bekommt gro\u00dfe Unterst\u00fctzung aus der Region. Der Landkreis Schmalkalden-Meiningen wird dem VfB 91 Suhl in der Spielzeit 2017\/2018 als Werbepartner zur Verf\u00fcgung stehen. Darauf einigte sich Landrat Peter Heimrich mit den Kreistagsfraktionsvorsitzenden.<\/p>\n<p>Geplant sind eine werbewirksame Platzierung auf den Wettkampftrikots der VfB-W\u00f6lfe und die Nutzung des Rotationsbandensystems, das durch den VBL-Livestream bundesweit f\u00fcr Aufmerksamkeit sorgt. \u201eDer Landkreis verspricht sich auf diese Weise einen \u00fcberregionalen Werbeeffekt und k\u00f6nnte beispielsweise Gewerbefl\u00e4chen im Industriegebiet Th\u00fcringer Tor vermarkten\u201c, so Heimrich.<\/p>\n<p>Nicht nur der Suhler Bundesligist profitiere von dieser Partnerschaft, auch die Volleyballvereine und der Nachwuchs im Landkreis behielten ihre sportlichen Vorbilder.<\/p>\n<p>\u201eGestern hatten wir Land in Sicht, heute stehen wir auf eben jenem\u201c, so VfB-Gesch\u00e4ftsf\u00fchrer Heiko Koch. \u201eDas ist ein sehr guter Beitrag der Region f\u00fcr unseren Club und unterstreicht den Stellenwert des VfB. Vor allem f\u00fcr die kurze Zeit der Entscheidungsfindung m\u00f6chte ich mich bedanken und hervorheben, dass Suhl, unter anderem \u00fcber die Beteiligung an den Umbaukosten f\u00fcr die Sporthalle Wolfsgrube, \u00a0mit starken Partnern durchaus etwas bewegen kann!\u201c<\/p>\n<p>Der VfB Suhl LOTTO Th\u00fcringen setzt seine seit Wochen zur\u00fcckgestellten Pl\u00e4ne f\u00fcr die neue Spielzeit ab sofort um.<\/p>","created_at":"2017-05-31T18:59:42+0200","updated_at":"2017-05-31T19:01:01+0200","annotations":[],"mimetype":"text\/html","language":"de-DE","reading_time":0,"domain_name":"1.bundesliga.vfb-suhl.de","http_status":"200","_links":{"self":{"href":"\/api\/entries\/3784"}}}]}}`
const entriesExists = `{"http:\/\/0.0.0.0\/entry10":false,"http:\/\/0.0.0.0\/entry2":false, "http://0.0.0.0/entry3":true}`
const oneItem = `{"is_archived":0,"is_starred":0,"user_name":"wallabagotest","user_email":"wallabagotest@wallabag.org","user_id":1,"tags":[{"id":61,"label":"1min","slug":"1min"}],"id":3977,"title":"Datenschutz: \n        Freifunker m\u00fcssen erstmal keine Vorratsdaten speichern","url":"https:\/\/www.golem.de\/news\/datenschutz-freifunker-muessen-erstmal-keine-vorratsdaten-speichern-1706-128443.html","content":"\n<img src=\"https:\/\/wallabag.linux4tw.de\/assets\/images\/4\/d\/4d85bba1\/53e8966b.jpeg\" alt=\"Freifunk bleibt auch am Deutschen Eck in Koblenz frei von Vorratsdatenspeicherung - erst einmal.\" \/>","created_at":"2017-06-19T14:14:06+0200","updated_at":"2017-06-19T14:14:07+0200","annotations":[],"mimetype":"text\/html","language":"de-DE","reading_time":1,"domain_name":"www.golem.de","preview_picture":"https:\/\/wallabag.linux4tw.de\/assets\/images\/4\/d\/4d85bba1\/53e8966b.jpeg","http_status":"200","_links":{"self":{"href":"\/api\/entries\/3977"}}}`

func TestGetEntries(t *testing.T) {
	expectedLimit := 1
	expectedTotal := 3494
	expectedPage := 1
	expectedPages := 3494
	entries := GetEntries(mockGetOneEntry, 0, 0, "", "", 0, expectedLimit, "")
	if entries.Total != expectedTotal {
		t.Errorf("expected %v entry, but got %v", expectedTotal, entries.Total)
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

func mockGetOneEntry(url string, httpMethod string, postData []byte) []byte {
	return []byte(oneEntry)
}

func TestGetEntriesExists(t *testing.T) {
	urls := []string{"http://0.0.0.0/entry10", "http://0.0.0.0/entry2", "http://0.0.0.0/entry3"}
	existsResult := []bool{false, false, true}
	e := GetEntriesExists(mockGetEntriesExists, urls)
	for index := 0; index < len(urls); index++ {
		if e[urls[index]] != existsResult[index] {
			t.Errorf("for url %v expected exists=%v, but got %v", urls[index], e[urls[index]], existsResult[index])
		}
	}
}

func mockGetEntriesExists(url string, httpMethod string, postData []byte) []byte {
	return []byte(entriesExists)
}

func TestGetEntry(t *testing.T) {
	expectedID := 3977
	expectedIsArchived := 0
	e := GetEntry(mockGetEntry, expectedID)
	if e.ID != expectedID {
		t.Errorf("expected id=%v, but got %v", expectedID, e.ID)
	}
	if e.IsArchived != expectedIsArchived {
		t.Errorf("expected is_archive=%v, but got %v", expectedIsArchived, e.IsArchived)
	}
}

func mockGetEntry(url string, httpMethod string, postData []byte) []byte {
	return []byte(oneItem)
}
