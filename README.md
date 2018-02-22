# wallabago

Go wrapper for the [Wallabag](https://github.com/wallabag/wallabag/) API


## Links

- http://wallabag.org/
- https://github.com/wallabag/wallabag/
- http://doc.wallabag.org/en/master/developer/api.html


## Project Status

### Currently supported wallabag version

* wallago version 1.0.0 and before: compatible with wallabag 2.2.1 - 2.2.3
* wallago version 2.0.0: tested only with wallabag 2.3.2


### Travis CI

[![Build Status](https://travis-ci.org/Strubbl/wallabago.svg?branch=master)](https://travis-ci.org/Strubbl/wallabago)


### Go Report Card

[![Go Report Card Badge](https://goreportcard.com/badge/github.com/Strubbl/wallabago)](https://goreportcard.com/report/github.com/Strubbl/wallabago)


### Status of the implementation of the API calls

#### GET
- [x] `GET /api/annotations/{entry}.{_format}`
- [x] `GET /api/entries.{_format}`
- [x] `GET /api/entries/exists.{_format}`
- [x] `GET /api/entries/{entry}.{_format}`
- [ ] `GET /api/entries/{entry}/export.{_format}`
- [x] `GET /api/entries/{entry}/tags.{_format}`
- [x] `GET /api/tags.{_format}`
- [x] `GET /api/version.{_format}`

#### POST
- [ ] `POST /api/annotations/{entry}.{_format}`
- [x] `POST /api/entries.{_format}`
- [ ] `POST /api/entries/{entry}/tags.{_format}`

#### PUT
- [ ] `PUT /api/annotations/{annotation}.{_format}`

#### DELETE
- [ ] `DELETE /api/annotations/{annotation}.{_format}`
- [ ] `DELETE /api/entries/{entry}.{_format}`
- [ ] `DELETE /api/entries/{entry}/tags/{tag}.{_format}`
- [ ] `DELETE /api/tag/label.{_format}`
- [ ] `DELETE /api/tags/label.{_format}`
- [ ] `DELETE /api/tags/{tag}.{_format}`

#### PATCH
- [ ] `PATCH /api/entries/{entry}.{_format}`
- [ ] `PATCH /api/entries/{entry}/reload.{_format}`

