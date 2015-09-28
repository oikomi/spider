package core

import (
	"container/list"
	"fmt"
	"reflect"
)

type LinkQueue struct {
	visted   *Queue
	unVisted *Queue
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{
		visted:   NewQueue(),
		unVisted: NewQueue(),
	}
}

// func (l *LinkQueue) getVisitedUrl() []string {
//     return l.visted
// }
//
func (l *LinkQueue) getUnvisitedUrl() string {
	return l.unVisted.Dequeue().Value.(string)
}

func (l *LinkQueue) addVistedUrl(url string) {
	l.visted.Enqueue(url)
}

func (l *LinkQueue) addUnVistedUrl(url string) {
	l.unVisted.Enqueue(url)
}

func (l *LinkQueue) unVistedUrlsEmpty() bool {
	return l.unVisted.Size() == 0
}

func (l *LinkQueue) getUnvistedUrlCount() int {
	return l.unVisted.Size()
}

func (l *LinkQueue) isUrlInVisted(url string) bool {
	return l.visted.Contain(url)
}

func (l *LinkQueue) dispalyVisted() {
	l.visted.display()
}

func (l *LinkQueue) dispalyUnVisted() {
	l.unVisted.display()
}

//
// func (l *LinkQueue) getUnvistedUrlCount() int {
//     return len(l.unVisited)
// }

type Queue struct {
	sem  chan int
	list *list.List
}

var tFunc func(val interface{}) bool

func NewQueue() *Queue {
	sem := make(chan int, 1)
	list := list.New()
	return &Queue{
		sem:  sem,
		list: list,
	}
}

func (q *Queue) Size() int {
	return q.list.Len()
}

func (q *Queue) Enqueue(val interface{}) *list.Element {
	q.sem <- 1
	e := q.list.PushFront(val)
	<-q.sem
	return e
}

func (q *Queue) Dequeue() *list.Element {
	q.sem <- 1
	e := q.list.Back()
	q.list.Remove(e)
	<-q.sem
	return e
}

func (q *Queue) Query(queryFunc interface{}) *list.Element {
	q.sem <- 1
	e := q.list.Front()
	for e != nil {
		if reflect.TypeOf(queryFunc) == reflect.TypeOf(tFunc) {
			if queryFunc.(func(val interface{}) bool)(e.Value) {
				<-q.sem
				return e
			}
		} else {
			<-q.sem
			return nil
		}
		e = e.Next()
	}
	<-q.sem
	return nil
}

func (q *Queue) Contain(val interface{}) bool {
	q.sem <- 1
	e := q.list.Front()
	for e != nil {
		if e.Value == val {
			<-q.sem
			return true
		} else {
			e = e.Next()
		}
	}
	<-q.sem
	return false
}

func (q *Queue) display() {
	for e := q.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// type LinkQueue struct {
//     visted      []string
//     unVisited   []string
// }
//
// func NewLinkQueue() *LinkQueue{
//     return &LinkQueue {
//         visted : make([]string, 0),
//         unVisited : make([]string, 0),
//     }
// }
//
//
// func (l *LinkQueue) getVisitedUrl() []string {
//     return l.visted
// }
//
// func (l *LinkQueue) getUnvisitedUrl() []string {
//     return l.unVisited
// }
//
// func (l *LinkQueue) addVisitedUrl(url string) {
//     l.visted = append(l.visted, url)
// }
//
// func (l *LinkQueue) removeVisitedUrl(url string) {
//     //append(l.visted, url)
// }
//
// func (l *LinkQueue) unVisitedUrlsEnmpy() bool {
//     return len(l.unVisited) == 0
// }
//
// func (l *LinkQueue) getUnvistedUrlCount() int {
//     return len(l.unVisited)
// }
