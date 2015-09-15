
package conf

import (
    "time"
)

type Config struct {
    Spider struct {
        UrlListFile     string
        OutputDirectory string
        MaxDepth        int
        CrawlInterval   int
        CrawlTimeout    time.Duration
        TargetUrl       string
        ThreadCount     int
    }
}
