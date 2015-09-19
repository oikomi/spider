#encoding=utf-8
from BeautifulSoup import BeautifulSoup
import socket
import urllib2
import re

class MyCrawler:
    def __init__(self,seeds):
        #使用种子初始化url队列
        self.linkQuence=linkQuence()
        if isinstance(seeds,str):
            self.linkQuence.addUnvisitedUrl(seeds)
        if isinstance(seeds,list):
            for i in seeds:
                self.linkQuence.addUnvisitedUrl(i)
        print "Add the seeds url \"%s\" to the unvisited url list"%str(self.linkQuence.unVisited)
    #抓取过程主函数
    def crawling(self,seeds,crawl_count):
        #循环条件：待抓取的链接不空且专区的网页不多于crawl_count
        while self.linkQuence.unVisitedUrlsEnmpy() is False and self.linkQuence.getVisitedUrlCount()<=crawl_count:
            #队头url出队列
            visitUrl=self.linkQuence.unVisitedUrlDeQuence()
            print "Pop out one url \"%s\" from unvisited url list"%visitUrl
            if visitUrl is None or visitUrl=="":
                continue
            #获取超链接
            links=self.getHyperLinks(visitUrl)
            print "Get %d new links"%len(links)
            #将url放入已访问的url中
            self.linkQuence.addVisitedUrl(visitUrl)
            print "Visited url count: "+str(self.linkQuence.getVisitedUrlCount())
            #未访问的url入列
            for link in links:
                self.linkQuence.addUnvisitedUrl(link)
            print "%d unvisited links:"%len(self.linkQuence.getUnvisitedUrl())

    #获取源码中得超链接
    def getHyperLinks(self,url):
        links=[]
        data=self.getPageSource(url)
        if data[0]=="200":
            soup=BeautifulSoup(data[1])
            a=soup.findAll("a",{"href":re.compile(".*")})
            for i in a:
                if i["href"].find("http://")!=-1:
                    links.append(i["href"])
        return links

    #获取网页源码
    def getPageSource(self,url,timeout=100,coding=None):
        try:
            socket.setdefaulttimeout(timeout)
            req = urllib2.Request(url)
            req.add_header('User-agent', 'Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)')
            response = urllib2.urlopen(req)
            if coding is None:
                coding= response.headers.getparam("charset")
            if coding is None:
                page=response.read()
            else:
                page=response.read()
                page=page.decode(coding).encode('utf-8')
            return ["200",page]
        except Exception,e:
            print str(e)
            return [str(e),None]

class linkQuence:
    def __init__(self):
        #已访问的url集合
        self.visted=[]
        #待访问的url集合
        self.unVisited=[]
    #获取访问过的url队列
    def getVisitedUrl(self):
        return self.visted
    #获取未访问的url队列
    def getUnvisitedUrl(self):
        return self.unVisited
    #添加到访问过得url队列中
    def addVisitedUrl(self,url):
        self.visted.append(url)
    #移除访问过得url
    def removeVisitedUrl(self,url):
        self.visted.remove(url)
    #未访问过得url出队列
    def unVisitedUrlDeQuence(self):
        try:
            return self.unVisited.pop()
        except:
            return None
    #保证每个url只被访问一次
    def addUnvisitedUrl(self,url):
        if url!="" and url not in self.visted and url not in self.unVisited:
            self.unVisited.insert(0,url)
    #获得已访问的url数目
    def getVisitedUrlCount(self):
        return len(self.visted)
    #获得未访问的url数目
    def getUnvistedUrlCount(self):
        return len(self.unVisited)
    #判断未访问的url队列是否为空
    def unVisitedUrlsEnmpy(self):
        return len(self.unVisited)==0

def main(seeds,crawl_count):
    craw=MyCrawler(seeds)
    craw.crawling(seeds,crawl_count)
if __name__=="__main__":
    main(["http://www.baidu.com"],50)
