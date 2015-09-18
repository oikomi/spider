package util

import (
	"net/url"
	//"strings"
)

func ParseSchemeHost(rawurl string) (string, error) {
    u, err := url.Parse(rawurl)
    if err != nil {
    	return "", err
    }

    return u.Scheme + "://" + u.Host, nil
}
