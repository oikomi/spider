package core

import (
    "io"
    "io/ioutil"
)

import (
    "github.com/golang/glog"
    "golang.org/x/net/html/charset"
)

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
