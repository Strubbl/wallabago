package wallabago

// Version returns the version of the configured wallabag instance
func Version(bodyStringGetterFunc BodyStringGetter) string {
	v := bodyStringGetterFunc(Config.WallabagURL+"/api/version", "GET", nil)
	// strip of the quotation marks from the version string being return from the API
	return v[1 : len(v)-1]
}
