/*
*
* Author: Hui Ye - <bonjovis@163.com>
*
* Last modified: 2017-03-17 09:28
*
* Filename: workpool.go
*
* Copyright (c) 2016 JOVI
*
 */
package work

import "sync"

type Worker interface {
	Work()
}

type Pool struct {
	works    chan Worker
	active   chan int
	kill     chan string
	shutdown chan int
	wg       sync.WaitGroup
	routines int
	logFunc  func(message string)
}

func WorkPool(routines int, logFunc func(message string)) *Pool {
	p := Pool{
		routines: routines,
		works:    make(chan Worker),
		active:   make(chan int),
		kill:     make(chan string),
		shutdown: make(chan int),
		logFunc:  logFunc,
	}
	p.initThreads()
	for i := 0; i < routines; i++ {
		p.active <- i
	}
	return &p
}

func (p *Pool) SetWork(work Worker) {
	p.works <- work
}

func (p *Pool) work() {
done:
	for {
		select {
		case w := <-p.works:
			w.Work()
		case <-p.kill:
			break done
		}
	}
	p.wg.Done()
	p.log("Worker : Shutting Down")
}

func (p *Pool) initThreads() {
	p.wg.Add(1)
	go func() {
		p.log("Threads : Started")
		for {
			select {
			case <-p.shutdown:
				for i := 0; i < p.routines; i++ {
					p.kill <- "over"
				}
				p.wg.Done()
				p.log("Threads : Closed")
				return
			case <-p.active:
				p.wg.Add(1)
				go p.work()
			}
		}
	}()
}

func (p *Pool) Shutdown() {
	close(p.shutdown)
	p.wg.Wait()
}

func (p *Pool) log(message string) {
	if p.logFunc != nil {
		p.logFunc(message)
	}
}
