package core

import (
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

    currentDeepth int

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
        if d.currentDeepth >= d.cfg.Spider.MaxDepth {
            glog.Info("!!!!!!!!!")
            d.linkQueue.dispalyVisted()
            break
        } else {
            for {
                //d.linkQueue.dispalyUnVisted()
                //glog.Info("-------")
                //d.linkQueue.dispalyVisted()
                if !d.linkQueue.unVistedUrlsEnmpy() {
                    url := d.linkQueue.getUnvisitedUrl()
                    glog.Info(url)
                    if d.linkQueue.isUrlInVisted(url) {
                        glog.Info("撞墙")
                        continue
                    }
                    glog.Info(url)
                    d.getHyperLinks(url)

                    d.linkQueue.addVistedUrl(url)
                } else {
                    glog.Info("!!!!!!!!!")
                    d.linkQueue.dispalyVisted()
                    break
                }
            }
        }

        d.currentDeepth ++
    }

    return nil
}

func (d *DownLoader)getHyperLinks(url string) error {
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
            //fmt.Println(link)
            if err != nil {
                glog.Error(err.Error())
        	}
            if link != "" {
                //glog.Info("add url to unvisited list")
                //glog.Info(link)
                d.linkQueue.addUnVistedUrl(link)
            }
        }
    })

    return nil
}
