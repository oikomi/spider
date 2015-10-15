package core

import (
    "fmt"
    "time"
    "strings"
    "sync/atomic"
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

    currentDeepth uint64

    baseHref      string

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
                    for i := 0; i < unvisitedNum ; i++ { //&& i < c.cfg.Spider.ThreadCount; i++ {
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
                                time.Sleep(c.cfg.Spider.CrawlInterval * time.Second)
                                err = c.getHyperLinks(url)
                                if err != nil {
                                    glog.Error(err.Error())
                                }

                                //d.linkQueue.addVistedUrl(url)
                            })
                            //time.Sleep(c.cfg.Spider.CrawlInterval * time.Second)
                    }   
                    fmt.Println("---1----")  
                    c.waitGroup.Wait()
                    fmt.Println("---2----")  
                    //c.currentDeepth ++
                    atomic.AddUint64(&c.currentDeepth, 1)
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
    fmt.Println("getHyperLinks: " + url)
    c.baseHref = ""

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

    doc.Find("base").Each(func(i int, s *goquery.Selection) {
        tmp, exits := s.Attr("href")
        if exits {
            c.baseHref = tmp
        }
    })

    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        link, exits := s.Attr("href")
        if exits {
            if ! strings.Contains(strings.ToLower(link),strings.ToLower("javascript")) {
                link, err = util.CheckLink(link, c.host, url, c.baseHref)
                //fmt.Println(link)
                if err != nil {
                    glog.Error(err.Error())
            	}
                if link != "" {
                    //glog.Info("add url to unvisited list")
                    //glog.Info(link)
                    c.linkQueue.addUnVistedUrl(link)
                }
            } else {
                jslink := strings.SplitN(link, "=", 2)[1]
                
                jslink, err = util.CheckLink(strings.Replace(jslink, "\"", "", -1), c.host, "", c.baseHref)
                //fmt.Println(jslink)
                if err != nil {
                    glog.Error(err.Error())
                }
                if jslink != "" {
                    //glog.Info("add url to unvisited list")
                    //glog.Info(jslink)
                    c.linkQueue.addUnVistedUrl(jslink)
                }
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
