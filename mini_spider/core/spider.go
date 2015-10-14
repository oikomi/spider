package core

import (
	"time"
)

import (
	"github.com/golang/glog"
)

import (
	"Go-id-3957/mini_spider/conf"
)

type Spider struct {
	cfg conf.Config

	rootUrlList   []string
	maxDepth      int
	crawlInterval time.Duration
	crawlTimeout  time.Duration
	targetUrl     string
	threadCount   int
}

func NewSpider(cfg conf.Config, seedUrlList []string) *Spider {
	return &Spider{
		cfg:           cfg,
		rootUrlList:   seedUrlList,
		maxDepth:      cfg.Spider.MaxDepth,
		crawlInterval: cfg.Spider.CrawlInterval,
		crawlTimeout:  cfg.Spider.CrawlTimeout,
		targetUrl:     cfg.Spider.TargetUrl,
		threadCount:   cfg.Spider.ThreadCount,
	}
}

func (s *Spider) Start() {
	var err error
	for _, rootUrl := range s.rootUrlList {
		c := NewCrawler(rootUrl, s.cfg)
		err = c.crawling()
		if err != nil {
			glog.Error(err.Error())
		}
	}
}

func (s *Spider) Stop() {
	glog.Info("Stop")
}
