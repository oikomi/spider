package core

import (
    "net/http"
)

import (
    "github.com/PuerkitoBio/goquery"
)

import (
    //"Go-id-3957/mini_spider/util"
)

func parse(httpRes *http.Response) error {
    linklist := make([]string, 0)

    doc, err := goquery.NewDocumentFromResponse(httpRes)
    if err != nil {
		return err
	}
    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        link, exits := s.Attr("href")
        if exits {
            //link = util.CheckLink(link)
            if link != "" {
                linklist = append(linklist, link)
            }
        }
    })

    return nil
}
