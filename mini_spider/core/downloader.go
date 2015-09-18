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
    "Go-id-3957/mini_spider/conf"
)

type DownLoader struct {
    cfg           conf.Config
    host          string
    seed          string
    crawlTimeout  time.Duration

    linkQueue     *LinkQueue
}

func NewDownLoader(seed string, cfg conf.Config) *DownLoader {
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
        cfg  : cfg,
        host : host,
        seed : seed,
        crawlTimeout : cfg.Spider.CrawlTimeout,
        linkQueue : initLq,
    }
}

func (d *DownLoader)crawling() error {
    for {
        if !d.linkQueue.unVistedUrlsEnmpy() {
            url := d.linkQueue.getUnvisitedUrl()
            if d.linkQueue.isUrlInVisted(url) {
                continue
            }
            fmt.Println(url)
            d.getHyperLinks(url)

            d.linkQueue.addVistedUrl(url)
        } else {
            d.linkQueue.dispalyVisted()
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
                d.linkQueue.addUnVistedUrl(link)
            }
        }
    })

    return nil
}
