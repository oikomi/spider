package core

import (
    "time"
)

import (
    "Go-id-3957/mini_spider/conf"
)

type Spider struct {
    rootUrlList     []string
    maxDepth        int
    crawlInterval   int
    crawlTimeout    time.Duration
    targetUrl       string
    threadCount     int
}

func NewSpider(cfg conf.Config, seedUrlList []string) *Spider {
    return &Spider {
        rootUrlList   : seedUrlList,
        maxDepth      : cfg.Spider.MaxDepth,
        crawlInterval : cfg.Spider.CrawlInterval,
        crawlTimeout  : cfg.Spider.CrawlTimeout,
        targetUrl     : cfg.Spider.TargetUrl,
        threadCount   : cfg.Spider.ThreadCount,
    }
}

// func (s *Spider)run(rootUrl string) {
//     d := NewDownLoader(rootUrl, s.crawlTimeout)
//     d.crawling()
// }

func (s *Spider)Start() {
    // for _, rootUrl := range s.rootUrlList {
    //     s.run(rootUrl)
    // }
    d := NewDownLoader(s.rootUrlList, s.crawlTimeout)
    d.crawling()
}
