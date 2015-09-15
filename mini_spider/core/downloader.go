package core

import (
    "time"
)

type DownLoader struct {
    rootUrl       string
    crawlTimeout  time.Duration
}

func NewDownLoader() *DownLoader {
    return &DownLoader{

    }
}


func (d *DownLoader)Download() {
    rh := NewReqHttp(d.rootUrl, "GET", d.crawlTimeout)
    rh.AddHeader("User-agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
    rh.DoGetData()
}
