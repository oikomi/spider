package core

import (
    //"fmt"
    "time"
    //"net/http"
)

import (
    "github.com/PuerkitoBio/goquery"
)

type DownLoader struct {
    seeds         []string
    crawlTimeout  time.Duration

    lq            *LinkQueue
}

func NewDownLoader(seeds []string, timeout time.Duration) *DownLoader {
    initLq := NewLinkQueue()
    for _, s := range seeds {
        initLq.addUnVistedUrl(s)
    }

    return &DownLoader {
        seeds : seeds,
        crawlTimeout : timeout,
        lq : initLq,
    }
}


func (d *DownLoader)crawling() error {
    for {
        if !d.lq.unVistedUrlsEnmpy() {
            url := d.lq.getUnvisitedUrl()
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
    rh := NewReqHttp(url, "GET", d.crawlTimeout)
    rh.AddHeader("User-agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
    httpRes, err := rh.DoGetData()
    if err != nil {
		return err
	}

    doc, err := goquery.NewDocumentFromResponse(httpRes)
    if err != nil {
		return err
	}
    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        link, exits := s.Attr("href")
        if exits {
            //link = util.CheckLink(link)
            if link != "" {
                //d.lq.unVisited = append(d.lq.unVisited, link)
                d.lq.addUnVistedUrl(link)
            }
        }
    })

    return nil
}
