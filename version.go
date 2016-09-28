package wallabago

func Version() string {
	return getBodyOfURL(Config.WallabagURL + "/api/version")
}
