package pathutil

import (
	"net/url"
	"path"
)

func UrlJoin(baseUrl string, paths ...string) string {
	u, _ := url.Parse(baseUrl)
	pathElements := append([]string{u.Path}, paths...)
	u.Path = path.Join(pathElements...)
	return u.String()
}
