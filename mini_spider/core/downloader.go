package core

import (
    "time"
    "sync"
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

type WaitGroupWrapper struct {
    sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
    w.Add(1)
    go func() {
        cb()
        w.Done()
    }()
}

type DownLoader struct {
    cfg           conf.Config

    host          string
    seed          string
    crawlTimeout  time.Duration

    currentDeepth int

    linkQueue     *LinkQueue

    waitGroup     *WaitGroupWrapper
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

    waitGroup := &WaitGroupWrapper{}

    return &DownLoader {
        cfg  : cfg,
        host : host,
        seed : seed,
        crawlTimeout : cfg.Spider.CrawlTimeout,
        linkQueue : initLq,
        waitGroup : waitGroup,
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

                    //
                    for i := 0; i < d.linkQueue.getUnvistedUrlCount() && i < d.cfg.Spider.ThreadCount; i++ {
                        url := d.linkQueue.getUnvisitedUrl()
                        glog.Info(url)
                        if d.linkQueue.isUrlInVisted(url) {
                            glog.Info("撞墙")
                            continue
                        }
                        glog.Info(url)

                        //d.getHyperLinks(url)

                        d.waitGroup.Wrap(func() {
                            d.getHyperLinks(url)
                            time.Sleep(d.cfg.Spider.CrawlInterval * time.Second)
                            //d.linkQueue.addVistedUrl(url)
                        })
                    }

                    // url := d.linkQueue.getUnvisitedUrl()
                    // glog.Info(url)
                    // if d.linkQueue.isUrlInVisted(url) {
                    //     glog.Info("撞墙")
                    //     continue
                    // }
                    // glog.Info(url)
                    // time.Sleep(d.cfg.Spider.CrawlInterval * time.Second)
                    // //d.getHyperLinks(url)
                    //
                    // d.waitGroup.Wrap(func() {
                    //     d.getHyperLinks(url)
                    //     //d.linkQueue.addVistedUrl(url)
                    // })

                    d.waitGroup.Wait()

                    //d.linkQueue.addVistedUrl(url)
                } else {
                    glog.Info("!!!!!!!!!")
                    d.linkQueue.dispalyVisted()
                    break
                }
            }
        }

        d.currentDeepth ++
    }

    d.waitGroup.Wait()

    return nil
}

func (d *DownLoader)getHyperLinks(url string) error {
    //defer d.waitGroup.Done()
    d.linkQueue.addVistedUrl(url)

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
