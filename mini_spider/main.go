package main

import (
    "os"
    "fmt"
    //"sync"
    "flag"
    // "syscall"
    // "os/signal"
    "encoding/json"
)

import (
    "gopkg.in/gcfg.v1"
    "github.com/golang/glog"
)

import (
    "Go-id-4777/mini_spider/conf"
    "Go-id-4777/mini_spider/core"
)

const (
    VERSION string = "1.0.0"
    CONF_NAME string = "conf.ini"
)

func init() {
    flag.Set("alsologtostderr", "false")
    //flag.Set("log_dir", "false")
}

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
    	glog.Error(err.Error())
    	return nil, err
    }
    defer file.Close()

    dec := json.NewDecoder(file)
    err = dec.Decode(&seedUrlList)
    if err != nil {
        glog.Error(err.Error())
    	return nil, err
    }
    
    return seedUrlList, nil
}

func main() {
    // signalChan := make(chan os.Signal, 1)
    // exitChan := make(chan int)
    // go func() {
    //     <-signalChan
    //     exitChan <- 1
    // }()

    // signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

    showHelp := flag.Bool("h", false, "show help info")
    showVersion := flag.Bool("vv", false, "show version info")
    confPath := flag.String("c", "../conf", "config conf path")
    logPath := flag.String("l", "../log", "config conf path")
    flag.Parse()

    flag.Set("log_dir", *logPath)

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
        glog.Error(err.Error())
        os.Exit(1)
    }
    seedUrlList, err := parseSeedUrls(cfg.Spider.UrlListFile)
    if err != nil {
        glog.Error(err.Error())
        os.Exit(1)
    }
    spider := core.NewSpider(cfg, seedUrlList)

    spider.Start()
}
