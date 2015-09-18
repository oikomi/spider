package util

import (
	"net/url"
	//"strings"
)

func ParseHost(rawurl string) (string, error) {
    u, err := url.Parse(rawurl)
    if err != nil {
    	return "", err
    }

    return u.Host, nil
}
