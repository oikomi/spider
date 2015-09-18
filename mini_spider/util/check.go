package util

import (
	//"net/url"
	//"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// func CheckBaseurl(rawurl string) (string, error) {
// 	u, err := url.Parse(rawurl)
//     if err != nil {
// 		return "", err
// 	}
// 	if u.Scheme == "" {
// 		u = "http://" + u
// 	}
// 	if flag := strings.HasSuffix(u, "/"); flag != true {
// 		u = u + "/"
// 	}
// 	return u
// }
//
// func CheckLink(link string) string {
// 	u, _ := url.Parse(link)
// 	if u.Scheme != "" {
// 		return ""
// 	}
// 	if u.Scheme == "http" || u.Scheme == "https" {
// 		return link
// 	}
// 	if flag := strings.HasPrefix(link, Config.StartUrl); flag != true {
// 		link = strings.Join([]string{Config.StartUrl, link}, "")
// 		return link
// 	}
// 	return ""
// }
