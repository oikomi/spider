package util

import (
	"net/url"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckBaseurl(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
    if err != nil {
		return "", err
	}

	if u.Scheme == "" {
		rawUrl = "http://" + rawUrl
	}
	if flag := strings.HasSuffix(rawUrl, "/"); flag != true {
		rawUrl = rawUrl + "/"
	}
	return rawUrl, nil
}

func CheckLink(link string, host string) (string, error) {
	u, err := url.Parse(link)
    if err != nil {
		return "", err
	}
	if u.Scheme != "" {
		return "", nil
	}
	if u.Scheme == "http" || u.Scheme == "https" {
		return link, nil
	}
	if flag := strings.HasPrefix(link, host); flag != true {
		link = strings.Join([]string{host, link}, "/")
		return link, nil
	}
	return "", nil
}
