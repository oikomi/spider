package util

import (
	"net/url"
	"strings"
)

import (
    "github.com/golang/glog"
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

func CheckLink(link, host, rawurl, basePath string) (string, error) {
	basePathList := strings.Split(rawurl, "/")

	var newBasePath string

	// for _, p := range basePathList {
	// 	newBasePath += p 
	// 	newBasePath += "/"
	// }

	for i := 0; i < len(basePathList) -1; i++ {
		newBasePath += basePathList[i] 
		newBasePath += "/"
	}

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
		if basePath == "" {
			link = strings.Join([]string{host, link}, "/")
		} else {
			//glog.Info("basePath is not null ")
			//link = strings.Join([]string{host + newBasePath + "/" + basePath, link}, "/")
			link = newBasePath + "/" + basePath + "/" + link
		}
		//link = strings.Join([]string{host + "/" + basePath, link}, "/")
		return link, nil
	}
	return "", nil
}

func CheckSrcLink(link, currentPath string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		glog.Error(err.Error())
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
