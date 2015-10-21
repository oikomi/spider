package core

import (
	//"fmt"
	"bytes"
	"net/http"
	"time"
	//"io/ioutil"
)

import (
	"github.com/golang/glog"
)

type ReqHttp struct {
	httpClient *http.Client
	method     string
	url        string
	header     http.Header
}

func NewReqHttp(url string, method string, timeout time.Duration) *ReqHttp {
	client := http.Client{
		Timeout: time.Duration(timeout * time.Second),
	}

	return &ReqHttp{
		method:     method,
		url:        url,
		httpClient: &client,
		header:     make(http.Header),
	}
}

func (r *ReqHttp) AddHeader(key, val string) {
	r.header.Add(key, val)
}

func (r *ReqHttp) SetHeader(key, val string) {
	r.header.Set(key, val)
}

func (r *ReqHttp) DoGetData() (*http.Response, error) {
	var err error
	request, err := http.NewRequest(r.method, r.url, nil)
	if err != nil {
		glog.Error(err.Error())
		return nil, err
	}

	//add header
	request.Header = r.header

	response, err := r.httpClient.Do(request)
	if err != nil {
		glog.Error(err.Error())
		return nil, err
	}

	if response.StatusCode == 200 {
		return response, nil
	} else {
		glog.Error(GET_DATA_FAILED)
		glog.Error("Failed url is : " + r.url)
		return nil, GET_DATA_FAILED
	}
}

func (r *ReqHttp) DoPostData(body []byte) error {
	var err error

	request, err := http.NewRequest(r.method, r.url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	//add header
	request.Header = r.header

	response, err := r.httpClient.Do(request)

	if err != nil {
		glog.Error(err.Error())
		return err
	}

	if response.StatusCode == 200 {

	} else {
		glog.Error(POST_DATA_FAILED)
		return POST_DATA_FAILED
	}

	return err
}
