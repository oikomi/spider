package core

import (
	//"time"
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
}

func NewSpider(cfg conf.Config, seedUrlList []string) *Spider {
	return &Spider{
		cfg:           cfg,
		rootUrlList:   seedUrlList,
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

