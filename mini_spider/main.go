package main

import (
    "os"
    "fmt"
    "flag"
    "encoding/json"
)

import (
    "gopkg.in/gcfg.v1"
)

import (
    "Go-id-3957/mini_spider/conf"
    "Go-id-3957/mini_spider/core"
)

const (
    VERSION string = "1.0.0"
    CONF_NAME string = "conf.ini"
)

func version() string {
	//fmt.Printf("mini_spider version %s \n", VERSION)
    return "mini_spider version " + VERSION
}

func showHelpInfo() {
    fmt.Println(`Usage of ./mini_spider:
      -c string
        	config conf path (default "../conf")
      -l string
        	config conf path (default "../log")
      -v	show version info
      -h	show help info
      `)
}

func parseSeedUrls(path string) ([]string, error) {
    var seedUrlList []string
    file, err := os.Open(path)
    if err != nil {
    	//log.Error(err.Error())
    	return nil, err
    }
    defer file.Close()

    dec := json.NewDecoder(file)
    err = dec.Decode(&seedUrlList)
    if err != nil {
    	return nil, err
    }
    return seedUrlList, nil
}

func main() {
    showHelp := flag.Bool("h", false, "show help info")
    //versionInfo := flag.Bool("v", true, "show version info")
    showVersion := flag.Bool("v", false, "show version info")
    confPath := flag.String("c", "../conf", "config conf path")
    logPath := flag.String("l", "../log", "config conf path")
    flag.Parse()
    fmt.Println(*logPath)
    if *showHelp {
        showHelpInfo()
        os.Exit(0)
    }

    if *showVersion {
        fmt.Println(version())
        os.Exit(0)
    }

    var cfg conf.Config
    err := gcfg.ReadFileInto(&cfg, *confPath + "/" + CONF_NAME)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(cfg)
    }

    seedUrlList, err := parseSeedUrls(cfg.Spider.UrlListFile)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(seedUrlList)
    }
    spider := core.NewSpider(cfg, seedUrlList)
    spider.Start()
}
