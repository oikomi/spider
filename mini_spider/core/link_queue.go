
package core

type LinkQueue struct {
    visted      []string
    unVisited   []string
}


func (l *LinkQueue) getVisitedUrl() []string {
    return l.visted
}

func (l *LinkQueue) getUnvisitedUrl() []string {
    return l.unVisited
}

func (l *LinkQueue) addVisitedUrl(url string) {
    l.visted = append(l.visted, url)
}

func (l *LinkQueue) removeVisitedUrl(url string) {
    //append(l.visted, url)
}

func (l *LinkQueue) unVisitedUrlsEnmpy() bool {
    return len(l.unVisited) == 0
}

func (l *LinkQueue) getUnvistedUrlCount() int {
    return len(l.unVisited)
}
