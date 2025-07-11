# wallabago

Go wrapper library for the [Wallabag](https://github.com/wallabag/wallabag/) API


## Links

- http://wallabag.org/
- https://github.com/wallabag/wallabag/
- https://doc.wallabag.org/en/developer/api/readme.html
- https://app.wallabag.it/api/doc
- [![Go Reference](https://pkg.go.dev/badge/github.com/Strubbl/wallabago.svg)](https://pkg.go.dev/github.com/Strubbl/wallabago)

## Project Status

### Currently supported wallabag version

* wallago version 1.0.0 and before: compatible with wallabag 2.2.1 - 2.2.3
* wallago version 2.0.0: tested only with wallabag 2.3.2
* wallago version 4.0.0 until 5.1.0: tested only with wallabag 2.3.8
* wallago version 6.0.0 until 6.0.6: tested only with wallabag 2.4.0 - 2.4.3
* wallago version 7.0.2: tested only with wallabag 2.4.3 - 2.5.1
* wallago version 7.0.3 until 7.0.4 version: tested only with wallabag 2.5.2 - 2.5.4
* wallago version 8.0.0 until 8.1.0: tested only with wallabag 2.5.4 - 2.6.4
* wallago version 8.2.0 until 8.2.1: tested only with wallabag 2.6.5 - 2.6.8
* wallago version 9.0.0 until latest version: tested only with wallabag 2.6.8


### Go Report Card

[![Go Report Card Badge](https://goreportcard.com/badge/github.com/Strubbl/wallabago)](https://goreportcard.com/report/github.com/Strubbl/wallabago)


### Status of the implementation of the API calls

#### GET
- [x] `GET /api/annotations/{entry}.{_format}`
- [x] `GET /api/config.{_format}`
- [x] `GET /api/entries.{_format}`
- [x] `GET /api/entries/exists.{_format}`
- [x] `GET /api/entries/{entry}.{_format}`
- [x] `GET /api/entries/{entry}/export.{_format}`
- [x] `GET /api/entries/{entry}/tags.{_format}`
- [x] `GET /api/info.{_format}`
- [ ] `GET /api/search.{_format}`
- [ ] `GET /api/taggingrule/export.{_format}`
- [x] `GET /api/tags.{_format}`
- [x] `GET /api/user.{_format}`
- [x] `GET /api/version.{_format}` DEPRECATED since wallabag version 2.4

#### POST
- [ ] `POST /api/annotations/{entry}.{_format}`
- [x] `POST /api/entries.{_format}`
- [ ] `POST /api/entries/lists.{_format}`
- [ ] `POST /api/entries/tags/lists.{_format}`
- [x] `POST /api/entries/{entry}/tags.{_format}`

#### PUT
- [ ] `PUT /api/annotations/{annotation}.{_format}`
- [ ] `PUT /api/user.{_format}`

#### DELETE
- [ ] `DELETE /api/annotations/{annotation}.{_format}`
- [ ] `DELETE /api/entries/list.{_format}`
- [ ] `DELETE /api/entries/tags/list.{_format}`
- [ ] `DELETE /api/entries/{entry}.{_format}`
- [x] `DELETE /api/entries/{entry}/tags/{tag}.{_format}`
- [ ] `DELETE /api/tag/label.{_format}`
- [ ] `DELETE /api/tags/label.{_format}`
- [ ] `DELETE /api/tags/{tag}.{_format}`

#### PATCH
- [ ] `PATCH /api/entries/{entry}.{_format}`
- [ ] `PATCH /api/entries/{entry}/reload.{_format}`

## Projects using wallabago

* [wallabako](https://gitlab.com/anarcat/wallabako) - wallabag client for Kobo readers
* [wallabag-stats](https://codeberg.org/strubbl/wallabag-stats) - draws a chart for unread and total articles in your wallabag instance
* [wallabag-add-article](https://codeberg.org/strubbl/wallabag-add-article) - commandline utility to add an article to wallabag
* [wallabag_import_pocket_tags](https://github.com/pbarry/wallabag_import_pocket_tags) - commandline utility to copy tags from an export of Pocket articles to your Wallabag articles
* [kindlebag](https://github.com/sum12/kindlebag) - download all articles from wallabag directly to a jailbroken kindle
* [walgot](https://git.sr.ht/~bacardi55/walgot) - a TUI wallabag client
* [wallabag-logseq](https://codeberg.org/strubbl/wallabag-logseq) - sync wallabag article annotations to Logseq
* [cleanABag](https://git.sr.ht/~bacardi55/cleanABag) - CLI tool for removing articles older than a given date from wallabag
* [article-summarizer](https://github.com/kahnwong/article-summarizer/)
* [wallabag-tagger](https://github.com/kahnwong/wallabag-tagger)
* [wallabag-to-epub](https://github.com/kahnwong/wallabag-to-epub)
