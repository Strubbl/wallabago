package wallabago

import "testing"

const oneEntry = `{"page":1,"limit":1,"pages":3494,"total":3494,"_links":{"self":{"href":"https:\/\/wallabag.org\/api\/entries?sort=created&order=desc&tags=&since=0&page=1&perPage=1"},"first":{"href":"https:\/\/wallabag.org\/api\/entries?sort=created&order=desc&tags=&since=0&page=1&perPage=1"},"last":{"href":"https:\/\/wallabag.org\/api\/entries?sort=created&order=desc&tags=&since=0&page=3494&perPage=1"},"next":{"href":"https:\/\/wallabag.org\/api\/entries?sort=created&order=desc&tags=&since=0&page=2&perPage=1"}},"_embedded":{"items":[{"is_archived":1,"is_starred":0,"user_name":"wallabagotest","user_email":"wallabagotest@wallabag.org","user_id":1,"tags":[{"id":61,"label":"1min","slug":"1min"}],"id":3784,"title":"VfB Suhl LOTTO Th\u00fcringen \u00bb Landkreis Schmalkalden-Meiningen unterst\u00fctzt Profi-Volleyball in S\u00fcdth\u00fcringen","url":"http:\/\/1.bundesliga.vfb-suhl.de\/landkreis-schmalkalden-meiningen-unterstuetzt-profi-volleyball-in-suedthueringen\/","content":"<p>Der S\u00fcdth\u00fcringer Bundesligist bekommt gro\u00dfe Unterst\u00fctzung aus der Region. Der Landkreis Schmalkalden-Meiningen wird dem VfB 91 Suhl in der Spielzeit 2017\/2018 als Werbepartner zur Verf\u00fcgung stehen. Darauf einigte sich Landrat Peter Heimrich mit den Kreistagsfraktionsvorsitzenden.<\/p>\n<p>Geplant sind eine werbewirksame Platzierung auf den Wettkampftrikots der VfB-W\u00f6lfe und die Nutzung des Rotationsbandensystems, das durch den VBL-Livestream bundesweit f\u00fcr Aufmerksamkeit sorgt. \u201eDer Landkreis verspricht sich auf diese Weise einen \u00fcberregionalen Werbeeffekt und k\u00f6nnte beispielsweise Gewerbefl\u00e4chen im Industriegebiet Th\u00fcringer Tor vermarkten\u201c, so Heimrich.<\/p>\n<p>Nicht nur der Suhler Bundesligist profitiere von dieser Partnerschaft, auch die Volleyballvereine und der Nachwuchs im Landkreis behielten ihre sportlichen Vorbilder.<\/p>\n<p>\u201eGestern hatten wir Land in Sicht, heute stehen wir auf eben jenem\u201c, so VfB-Gesch\u00e4ftsf\u00fchrer Heiko Koch. \u201eDas ist ein sehr guter Beitrag der Region f\u00fcr unseren Club und unterstreicht den Stellenwert des VfB. Vor allem f\u00fcr die kurze Zeit der Entscheidungsfindung m\u00f6chte ich mich bedanken und hervorheben, dass Suhl, unter anderem \u00fcber die Beteiligung an den Umbaukosten f\u00fcr die Sporthalle Wolfsgrube, \u00a0mit starken Partnern durchaus etwas bewegen kann!\u201c<\/p>\n<p>Der VfB Suhl LOTTO Th\u00fcringen setzt seine seit Wochen zur\u00fcckgestellten Pl\u00e4ne f\u00fcr die neue Spielzeit ab sofort um.<\/p>","created_at":"2017-05-31T18:59:42+0200","updated_at":"2017-05-31T19:01:01+0200","annotations":[],"mimetype":"text\/html","language":"de-DE","reading_time":0,"domain_name":"1.bundesliga.vfb-suhl.de","http_status":"200","_links":{"self":{"href":"\/api\/entries\/3784"}}}]}}`
const entriesExists = `{"http:\/\/0.0.0.0\/entry10":false,"http:\/\/0.0.0.0\/entry2":false, "http://0.0.0.0/entry3":true}`
const oneItem = `{"is_archived":0,"is_starred":0,"user_name":"wallabagotest","user_email":"wallabagotest@wallabag.org","user_id":1,"tags":[{"id":61,"label":"1min","slug":"1min"}],"id":3977,"title":"Datenschutz: \n        Freifunker m\u00fcssen erstmal keine Vorratsdaten speichern","url":"https:\/\/www.golem.de\/news\/datenschutz-freifunker-muessen-erstmal-keine-vorratsdaten-speichern-1706-128443.html","content":"\n<img src=\"https:\/\/wallabag.linux4tw.de\/assets\/images\/4\/d\/4d85bba1\/53e8966b.jpeg\" alt=\"Freifunk bleibt auch am Deutschen Eck in Koblenz frei von Vorratsdatenspeicherung - erst einmal.\" \/>Freifunk bleibt auch am Deutschen Eck in Koblenz frei von Vorratsdatenspeicherung - erst einmal. (Bild: <a href=\"https:\/\/commons.wikimedia.org\/w\/index.php?curid=15001660\" rel=\"nofollow\" target=\"_blank\">Holger Weinandt\/Bearbeitung Golem.de<\/a>\/<a href=\"http:\/\/creativecommons.org\/licenses\/by-sa\/3.0\/\" rel=\"nofollow\" target=\"_blank\">CC-BY-SA 3.0<\/a>)<p>\nFallen <a href=\"https:\/\/www.golem.de\/specials\/freifunk\/\" target=\"_blank\" class=\"golem-internal-url golem-url-specials\">Freifunk<\/a>-Initiativen unter die <a href=\"https:\/\/www.golem.de\/specials\/vorratsdatenspeicherung\/\" target=\"_blank\" class=\"golem-internal-url golem-url-specials\">Vorratsdatenspeicherung<\/a>? Der Verein Freifunk Rheinland gibt erst einmal Entwarnung: Nach Aussage der Bundesnetzagentur handele es sich nicht um einen Internetzugangsdienst.\n<\/p>\n<div class=\"formatted\">\n\n<p>Der Verein Freifunk im Rheinland muss auch nach Inkrafttreten der Speicherpflicht der umstrittenen Vorratsdatenspeicherung am 1. Juli erst einmal keine Daten der Nutzer speichern und auch keine entsprechende Infrastruktur aufbauen, <a href=\"https:\/\/www.freifunk-rheinland.net\/2017\/06\/19\/vds-ab-juli-nicht-fuer-freifunk-und-den-ffrl\/\" target=\"_blank\" class=\"golem-external-url\">wie er in einem Blogpost mitteilt<\/a>. Dies sei aus Diskussionen mit der Bundesnetzagentur hervorgegangen, teilt das B\u00fcndnis mit. Die Entscheidung d\u00fcrfte Signalwirkung f\u00fcr andere Freifunkprojekte haben, ist aber m\u00f6glicherweise nicht von Dauer.<\/p>\n&#13;\n&#13;\n&#13;\n\n<p>Derzeit ist nach Ansicht der Bundesnetzagentur nicht zweifelsfrei gekl\u00e4rt, ob der Freifunk Rheinland e.V. sowie vergleichbare Zugangsmodelle unter die Speicherpflicht fallen, weil es <em>\"zumindest erhebliche Zweifel\"<\/em> gebe, dass Freifunk ein Internetzugangsdienst in diesem Sinne sei.<\/p>\n<h3>Keine vorbeugende Vorratsdatenspeicherung<\/h3>\n<p>Selbst wenn die Bundesnetzagentur die Freifunker zu einem sp\u00e4teren Zeitpunkt als Internetzugangsdienste einstufen w\u00fcrde, sei nicht klar, ob Daten gespeichert werden m\u00fcssten. Denn der Verein vergebe keine Benutzerkennungen. Bis diese Fragen gekl\u00e4rt sind, gilt laut Bundesnetzagentur: <em>\"Bis zu einer Entscheidung \u00fcber die Einstufung des Modells bzw. \u00fcber die Verpflichtung zur Verkehrsdatenspeicherung m\u00fcssen keine Aktivit\u00e4ten - quasi vorbeugend - zur Umsetzung der Speicherpflicht nach \u00a7 113a Abs. 3 TKG unternommen werden.\"<\/em><\/p>\n<p id=\"gfpop\">Die Umsetzung der Vorratsdatenspeicherung ist f\u00fcr gr\u00f6\u00dfere Anbieter <a href=\"https:\/\/www.golem.de\/news\/infrastruktur-vorratsdaten-kosten-mobilfunkbetreiber-bis-zu-15-millionen-1706-128281.html\" target=\"_blank\" class=\"golem-internal-url golem-url-news\">mit zum Teil enormen Kosten verbunden<\/a>. Der Prozess l\u00e4sst sich f\u00fcr kleinere Anbieter aber auch <a href=\"https:\/\/www.golem.de\/news\/cloud-anbieter-vorratsdatenspeicherung-as-a-service-1705-127791.html\" target=\"_blank\" class=\"golem-internal-url golem-url-news\">zu deutlich geringeren Kosten auslagern<\/a>. Neben mehreren B\u00fcrgerrechtlern klagen auch <a href=\"https:\/\/www.golem.de\/news\/ip-adressen-deutsche-telekom-klagt-gegen-vorratsdatenspeicherung-1706-128232.html\" target=\"_blank\" class=\"golem-internal-url golem-url-news\">die Deutsche Telekom<\/a> und <a href=\"https:\/\/www.golem.de\/news\/bundesverfassungsgericht-vorratsdatenspeicherung-laesst-sich-vorerst-nicht-stoppen-1704-127303.html\" target=\"_blank\" class=\"golem-internal-url golem-url-news\">der Internetverband Eco<\/a> gegen die Vorratsdatenspeicherung.<\/p>\n\n<\/div>","created_at":"2017-06-19T14:14:06+0200","updated_at":"2017-06-19T14:14:07+0200","annotations":[{"annotator_schema_version":"v1.0","id":29,"text":"Anmerkung 1","created_at":"2017-06-19T21:54:37+0200","updated_at":"2017-06-19T21:54:37+0200","quote":"Bundesnetzagentur nicht zweifelsfrei","ranges":[{"start":"\/div[1]\/p[2]","startOffset":28,"end":"\/div[1]\/p[2]","endOffset":65}]},{"annotator_schema_version":"v1.0","id":30,"text":"Anmerkung 2","created_at":"2017-06-19T21:54:45+0200","updated_at":"2017-06-19T21:54:45+0200","quote":"werden m\u00fcssten","ranges":[{"start":"\/div[1]\/p[3]","startOffset":158,"end":"\/div[1]\/p[3]","endOffset":172}]}],"mimetype":"text\/html","language":"de-DE","reading_time":1,"domain_name":"www.golem.de","preview_picture":"https:\/\/wallabag.linux4tw.de\/assets\/images\/4\/d\/4d85bba1\/53e8966b.jpeg","http_status":"200","_links":{"self":{"href":"\/api\/entries\/3977"}}}`

func TestGetEntries(t *testing.T) {
	expectedLimit := 1
	expectedTotal := 3494
	expectedPage := 1
	expectedPages := 3494
	entries, _ := GetEntries(mockGetOneEntry, 0, 0, "", "", 0, expectedLimit, "")
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
	expectedID := 3977
	expectedIsArchived := 0
	expectedTitle := "Datenschutz: \n        Freifunker m\u00fcssen erstmal keine Vorratsdaten speichern"
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
