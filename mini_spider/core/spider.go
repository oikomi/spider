package core

import (
    "time"
)

import (
    "Go-id-3957/mini_spider/conf"
)

type Spider struct {
    cfg             conf.Config

    rootUrlList     []string
    maxDepth        int
    crawlInterval   time.Duration
    crawlTimeout    time.Duration
    targetUrl       string
    threadCount     int

}

func NewSpider(cfg conf.Config, seedUrlList []string) *Spider {

    return &Spider {
        cfg           : cfg,
        rootUrlList   : seedUrlList,
        maxDepth      : cfg.Spider.MaxDepth,
        crawlInterval : cfg.Spider.CrawlInterval,
        crawlTimeout  : cfg.Spider.CrawlTimeout,
        targetUrl     : cfg.Spider.TargetUrl,
        threadCount   : cfg.Spider.ThreadCount,
    }
}


func (s *Spider)Start() {
    for _, rootUrl := range s.rootUrlList {
        d := NewDownLoader(rootUrl, s.cfg)
        d.crawling()
    }

}
