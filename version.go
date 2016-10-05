package wallabago

// Version returns the version of the configured wallabag instance
func Version() string {
	v := getBodyOfURL(config.WallabagURL + "/api/version")
	return v[1 : len(v)-2]
}
