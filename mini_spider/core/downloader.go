package core

import (
    "fmt"
    "time"
    //"net/http"
)

import (
    "github.com/golang/glog"
    "github.com/PuerkitoBio/goquery"
)

import (
    "Go-id-3957/mini_spider/util"
)

type DownLoader struct {
    host          string
    seed          string
    crawlTimeout  time.Duration

    lq            *LinkQueue
}

func NewDownLoader(seed string, timeout time.Duration) *DownLoader {
    initLq := NewLinkQueue()
    initLq.addUnVistedUrl(seed)

    newSeed, err := util.CheckBaseurl(seed)
    if err != nil {
        glog.Error(err.Error())
    }

    host, err := util.ParseSchemeHost(newSeed)
    if err != nil {
        glog.Error(err.Error())
    }

    return &DownLoader {
        host : host,
        seed : seed,
        crawlTimeout : timeout,
        lq : initLq,
    }
}

func (d *DownLoader)crawling() error {
    for {
        if !d.lq.unVistedUrlsEnmpy() {
            url := d.lq.getUnvisitedUrl()
            if d.lq.isUrlInVisted(url) {
                continue
            }
            fmt.Println(url)
            d.getHyperLinks(url)

            d.lq.addVistedUrl(url)
        } else {
            d.lq.dispalyVisted()
            break
        }
    }

    return nil
}

func (d *DownLoader)getHyperLinks(url string) error {
    fmt.Println("*******")
    fmt.Println(url)
    fmt.Println("*******")
    rh := NewReqHttp(url, "GET", d.crawlTimeout)
    rh.AddHeader("User-agent", USER_AGENT)
    httpRes, err := rh.DoGetData()
    if err != nil {
        glog.Error(err.Error())
		return err
	}

    doc, err := goquery.NewDocumentFromResponse(httpRes)
    if err != nil {
        glog.Error(err.Error())
		return err
	}
    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        link, exits := s.Attr("href")
        if exits {
            link, err = util.CheckLink(link, d.host)
            fmt.Println(link)
            if err != nil {
                glog.Error(err.Error())
        	}
            if link != "" {
                fmt.Println("----")
                fmt.Println(link)
                fmt.Println("----")
                d.lq.addUnVistedUrl(link)
            }
        }
    })

    return nil
}
