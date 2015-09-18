package util

import (
	"io"
	"net/url"
	"io/ioutil"
	//"strings"
)

import (
	"github.com/golang/glog"
	"golang.org/x/net/html/charset"
)

func ParseSchemeHost(rawurl string) (string, error) {
    u, err := url.Parse(rawurl)
    if err != nil {
    	return "", err
    }

    return u.Scheme + "://" + u.Host, nil
}

func ChangeCharsetEncodingAuto(sor io.ReadCloser, contentTypeStr string) string {
	var err error
	destReader, err := charset.NewReader(sor, contentTypeStr)

	if err != nil {
		glog.Error(err.Error())
		destReader = sor
	}

	var sorbody []byte
	if sorbody, err = ioutil.ReadAll(destReader); err != nil {
		glog.Error(err.Error())
		// For gb2312, an error will be returned.
		// Error like: simplifiedchinese: invalid GBK encoding
		// return ""
	}
	//e,name,certain := charset.DetermineEncoding(sorbody,contentTypeStr)
	bodystr := string(sorbody)

	return bodystr
}
