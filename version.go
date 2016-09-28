package wallabago

func Version() string {
	return getBodyOfURL(WallabagURL + "/api/version")
}
