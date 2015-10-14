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

func CheckLink(link, host string) (string, error) {
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

func CheckSrcLink(link, currentPath string) (string, error) {
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

	pathList := strings.Split(currentPath, "/")
	var tmpPath string

	for i := 0; i < len(pathList) - 1; i++ {
		tmpPath += pathList[i] + "/"
	}

	return tmpPath + link, nil
}
