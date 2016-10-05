package wallabago

// Version returns the version of the configured wallabag instance
func Version() string {
	v := getBodyOfURL(Config.WallabagURL + "/api/version")
	return v[1 : len(v)-2]
}
