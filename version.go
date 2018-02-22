package wallabago

// Version returns the version of the configured wallabag instance
func Version(bodyStringGetterFunc BodyStringGetter) (string, error) {
	v, err := bodyStringGetterFunc(Config.WallabagURL+"/api/version", "GET", nil)
	if err != nil {
		return "", err
	}
	// strip of the quotation marks from the version string being return from the API
	return v[1 : len(v)-1], err
}
