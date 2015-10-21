package core

import (
	"os"
	"strings"
	//"net/url"
	"io/ioutil"
	"encoding/base64"
)

import (
    "github.com/golang/glog"
)

import (
    "Go-id-3957/mini_spider/conf"
)

func StorageBinaryData(path string, cfg conf.Config) error {
    rh := NewReqHttp(path, "GET", cfg.Spider.CrawlTimeout)
    rh.AddHeader("User-agent", USER_AGENT)
    httpRes, err := rh.DoGetData()
    if err != nil {
        glog.Error(err.Error())
		return err
	}

	body, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		glog.Error(err.Error())
		return err
	}

	pathList := strings.Split(path, "/")

	fout, err := os.Create(cfg.Spider.OutputDirectory + "/" + pathList[len(pathList) - 1])
	defer fout.Close()
	if err != nil {
		glog.Error(err.Error())
		return err
	}

	fout.Write(body)

	return nil
}

func StoragePageData(path string, cfg conf.Config) error {
    rh := NewReqHttp(path, "GET", cfg.Spider.CrawlTimeout)
    rh.AddHeader("User-agent", USER_AGENT)
    httpRes, err := rh.DoGetData()
    if err != nil {
        glog.Error(err.Error())
		return err
	}

	body, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		glog.Error(err.Error())
		return err
	}

	encodeStr := base64.StdEncoding.EncodeToString([]byte(path))

	fout, err := os.Create(cfg.Spider.OutputDirectory + "/" + encodeStr)
	defer fout.Close()
	if err != nil {
		glog.Error(err.Error())
		return err
	}

	fout.Write(body)

	return nil
}