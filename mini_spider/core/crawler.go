package core

import (
    "fmt"
    "time"
    //"sync"
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

// type WaitGroupWrapper struct {
//     sync.WaitGroup
// }

// func (w *WaitGroupWrapper) Wrap(cb func()) {
//     w.Add(1)
//     go func() {
//         cb()
//         w.Done()
//     }()
// }

type Crawler struct {
    cfg           conf.Config

    host          string
    seed          string
    crawlTimeout  time.Duration

    currentDeepth int

    linkQueue     *LinkQueue

    waitGroup     *util.WaitGroupWrapper
}

func NewCrawler(seed string, cfg conf.Config) *Crawler {
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

    waitGroup := &util.WaitGroupWrapper{}

    return &Crawler {
        cfg  : cfg,
        host : host,
        seed : seed,
        crawlTimeout : cfg.Spider.CrawlTimeout,
        linkQueue : initLq,
        waitGroup : waitGroup,
    }
}

func (c *Crawler) crawling() error {
    var err error
    for {
        if 0 == c.cfg.Spider.MaxDepth {
            fmt.Println(c.seed)
            break
        }

        fmt.Println("****")

        if c.currentDeepth >= c.cfg.Spider.MaxDepth {
            glog.Info("========== All links result ==========")
            c.linkQueue.dispalyVisted()
            break
        } else {
            for {
                if c.currentDeepth > c.cfg.Spider.MaxDepth {
                    break
                }
                //d.linkQueue.dispalyUnVisted()
                //glog.Info("-------")
                //d.linkQueue.dispalyVisted()
                if !c.linkQueue.unVistedUrlsEmpty() { //&& c.currentDeepth <= c.cfg.Spider.MaxDepth {

                    //
                    unvisitedNum := c.linkQueue.getUnvistedUrlCount()
                    // fixme : 
                    for i := 0; i < unvisitedNum && i < c.cfg.Spider.ThreadCount; i++ {
                        // for j := 0; j < c.cfg.Spider.ThreadCount; j++ {

                        // }
                            url := c.linkQueue.getUnvisitedUrl()
                            glog.Info(url)
                            if c.linkQueue.isUrlInVisted(url) {
                                glog.Info("撞墙")
                                continue
                            }
                            glog.Info(url)
                            c.waitGroup.Wrap(func() {
                                err = c.getHyperLinks(url)
                                if err != nil {
                                    glog.Error(err.Error())
                                }

                                time.Sleep(c.cfg.Spider.CrawlInterval * time.Second)
                                //d.linkQueue.addVistedUrl(url)
                            })
                    }     
                    c.waitGroup.Wait()
                    c.currentDeepth ++
                } else {
                    break
                }
            }
        }

        //c.currentDeepth ++
    }

    return nil
}

func (c *Crawler)getHyperLinks(url string) error {
    //defer c.waitGroup.Done()
    c.linkQueue.addVistedUrl(url)

    rh := NewReqHttp(url, "GET", c.crawlTimeout)
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
            link, err = util.CheckLink(link, c.host)
            //fmt.Println(link)
            if err != nil {
                glog.Error(err.Error())
        	}
            if link != "" {
                //glog.Info("add url to unvisited list")
                //glog.Info(link)
                c.linkQueue.addUnVistedUrl(link)
            }
        }
    })

    doc.Find("img").Each(func(i int, s *goquery.Selection) {
        link, exits := s.Attr("src")
        if exits {
            link, err = util.CheckSrcLink(link, url)
            //fmt.Println(link)
            if err != nil {
                glog.Error(err.Error())
            }
            if link != "" {
                //glog.Info("add url to unvisited list")
                //glog.Info(link)
                c.linkQueue.addUnVistedUrl(link)
            }
        }
    })

    return nil
}
