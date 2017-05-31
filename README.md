# wallabago

Go wrapper for the [Wallabag](https://github.com/wallabag/wallabag/) API


## Links

- http://wallabag.org/
- https://github.com/wallabag/wallabag/
- http://doc.wallabag.org/en/master/developer/api.html


## Project Status

### Currently supported wallabag version

2.2.1 - 2.2.3


### Travis CI

[![Build Status](https://travis-ci.org/Strubbl/wallabago.svg?branch=master)](https://travis-ci.org/Strubbl/wallabago)


### Go Report Card

[![Go Report Card Badge](https://goreportcard.com/badge/github.com/Strubbl/wallabago)](https://goreportcard.com/report/github.com/Strubbl/wallabago)


### Status of the implementation of the API calls

- [ ] `DELETE /annotations/{annotation}.{_format}`
- [ ] `PUT /annotations/{annotation}.{_format}`
- [ ] `GET /annotations/{entry}.{_format}`
- [ ] `POST /annotations/{entry}.{_format}`
- [x] `GET /api/entries.{_format}`
- [x] `POST /api/entries.{_format}`
- [ ] `DELETE /api/entries/{entry}.{_format}`
- [ ] `GET /api/entries/{entry}.{_format}`
- [ ] `PATCH /api/entries/{entry}.{_format}`
- [ ] `GET /api/entries/{entry}/tags.{_format}`
- [ ] `POST /api/entries/{entry}/tags.{_format}`
- [ ] `DELETE /api/entries/{entry}/tags/{tag}.{_format}`
- [ ] `GET /api/tags.{_format}`
- [ ] `DELETE /api/tags/{tag}.{_format}`
- [ ] `GET /api/version.{_format}`
